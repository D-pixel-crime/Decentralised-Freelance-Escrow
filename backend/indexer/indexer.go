package indexer

import (
	"context"
	"os"
	"time"

	"github.com/D-pixel-crime/Freelance_Escrow/backend/contracts/escrow"
	"github.com/D-pixel-crime/Freelance_Escrow/backend/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartIndexer(jobsColl *mongo.Collection) {
	log.Infof("Indexer started. Waiting for blocks & events...")

	if utils.Web3Client == nil {
		log.Errorf("Web3Client is nil, indexer cannot proceed.")
		return
	}

	contractAddressHex := os.Getenv("ESCROW_CONTRACT_ADDRESS")
	if contractAddressHex == "" {
		log.Errorf("ESCROW_CONTRACT_ADDRESS not set, indexer exiting.")
		return
	}

	contractAddress := common.HexToAddress(contractAddressHex)
	escrowInstance, err := escrow.NewFreelanceEscrow(contractAddress, utils.Web3Client)
	if err != nil {
		log.Errorf("Failed to instantiate FreelanceEscrow contract: %v", err)
		return
	}

	go watchClientStakeCompleted(jobsColl, escrowInstance, contractAddressHex)
	go watchFreelancerStakeCompleted(jobsColl, escrowInstance, contractAddressHex)
	go watchBothPartyStakeCompleted(jobsColl, escrowInstance, contractAddressHex)
	go watchFreelancerCompletedAndClientConfirmationPending(jobsColl, escrowInstance, contractAddressHex)
	go watchJobCompletedAndFreelancerPaid(jobsColl, escrowInstance, contractAddressHex)
	go watchJobCompletionRejected(jobsColl, escrowInstance, contractAddressHex)
	go watchDealCancelRequested(jobsColl, escrowInstance, contractAddressHex)
	go watchRevertedDealBreak(jobsColl, escrowInstance, contractAddressHex)
	go watchDealBroken(jobsColl, escrowInstance, contractAddressHex)
	go watchRandomDisputeRaised(jobsColl, escrowInstance, contractAddressHex)
	go watchPaymentDisputeRaised(jobsColl, escrowInstance, contractAddressHex)
	go watchDisputeResolved(jobsColl, escrowInstance, contractAddressHex)

	// Block main indexer goroutine to keep listeners alive
	select {}
}

func watchClientStakeCompleted(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowClientStakeCompleted)
	sub, err := escrowInstance.WatchFreelanceEscrowClientStakeCompleted(&bind.WatchOpts{Context: context.Background()}, sink, nil, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to ClientStakeCompleted event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to ClientStakeCompleted events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("ClientStakeCompleted subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected ClientStakeCompleted! Amount: %s, Timestamp: %s", event.Amount.String(), event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status":             "CLIENT_STAKED",
					"clientStakedAmount": event.Amount.String(),
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to CLIENT_STAKED", contractAddressHex)
			}
		}
	}
}

func watchFreelancerStakeCompleted(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
	sub, err := escrowInstance.WatchFreelanceEscrowFreelancerStakeCompleted(&bind.WatchOpts{Context: context.Background()}, sink, nil, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to FreelancerStakeCompleted event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to FreelancerStakeCompleted events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("FreelancerStakeCompleted subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected FreelancerStakeCompleted! Amount: %s, Timestamp: %s", event.Amount.String(), event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "FREELANCER_STAKED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to FREELANCER_STAKED", contractAddressHex)
			}
		}
	}
}

func watchBothPartyStakeCompleted(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
	sub, err := escrowInstance.WatchFreelanceEscrowBothPartyStakeCompleted(&bind.WatchOpts{Context: context.Background()}, sink, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to BothPartyStakeCompleted event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to BothPartyStakeCompleted events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("BothPartyStakeCompleted subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected BothPartyStakeCompleted! Timestamp: %s", event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "ALL_STAKED_AND_PENDING",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to ALL_STAKED_AND_PENDING", contractAddressHex)
			}
		}
	}
}

func watchFreelancerCompletedAndClientConfirmationPending(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
	sub, err := escrowInstance.WatchFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(&bind.WatchOpts{Context: context.Background()}, sink)
	if err != nil {
		log.Errorf("Failed to subscribe to FreelancerCompletedAndClientConfirmationPending event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to FreelancerCompletedAndClientConfirmationPending events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("FreelancerCompletedAndClientConfirmationPending subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected FreelancerCompletedAndClientConfirmationPending! Timestamp: %s", event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "PENDING_CLIENT_CONFIRMATION",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to PENDING_CLIENT_CONFIRMATION", contractAddressHex)
			}
		}
	}
}

func watchJobCompletedAndFreelancerPaid(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
	sub, err := escrowInstance.WatchFreelanceEscrowJobCompletedAndFreelancerPaid(&bind.WatchOpts{Context: context.Background()}, sink, nil, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to JobCompletedAndFreelancerPaid event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to JobCompletedAndFreelancerPaid events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("JobCompletedAndFreelancerPaid subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected JobCompletedAndFreelancerPaid! Amount: %s, Timestamp: %s", event.Amount.String(), event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "JOB_COMPLETED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to JOB_COMPLETED", contractAddressHex)
			}
		}
	}
}

func watchJobCompletionRejected(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowJobCompletionRejected)
	sub, err := escrowInstance.WatchFreelanceEscrowJobCompletionRejected(&bind.WatchOpts{Context: context.Background()}, sink)
	if err != nil {
		log.Errorf("Failed to subscribe to JobCompletionRejected event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to JobCompletionRejected events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("JobCompletionRejected subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected JobCompletionRejected! Timestamp: %s", event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "PAYMENT_DISPUTED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to PAYMENT_DISPUTED", contractAddressHex)
			}
		}
	}
}

func mapEscrowStateToString(state uint8) string {
	switch state {
	case 0:
		return "AGREED"
	case 1:
		return "CLIENT_STAKED"
	case 2:
		return "FREELANCER_STAKED"
	case 3:
		return "ALL_STAKED_AND_PENDING"
	case 4:
		return "PENDING_CLIENT_CONFIRMATION"
	case 5:
		return "JOB_COMPLETED"
	case 6:
		return "CANCEL_REQUESTED"
	case 7:
		return "DEAL_BROKEN"
	case 8:
		return "RANDOM_DISPUTED"
	case 9:
		return "PAYMENT_DISPUTED"
	default:
		return "UNKNOWN_STATE"
	}
}

func watchDealCancelRequested(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowDealCancelRequested)
	sub, err := escrowInstance.WatchFreelanceEscrowDealCancelRequested(&bind.WatchOpts{Context: context.Background()}, sink, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to DealCancelRequested event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to DealCancelRequested events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("DealCancelRequested subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected DealCancelRequested! Initiator: %s", event.Initiator.Hex())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "CANCEL_REQUESTED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to CANCEL_REQUESTED", contractAddressHex)
			}
		}
	}
}

func watchRevertedDealBreak(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowRevertedDealBreak)
	sub, err := escrowInstance.WatchFreelanceEscrowRevertedDealBreak(&bind.WatchOpts{Context: context.Background()}, sink, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to RevertedDealBreak event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to RevertedDealBreak events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("RevertedDealBreak subscription error: %v", err)
			return
		case event := <-sink:
			callCtx, callCancel := context.WithTimeout(context.Background(), 5*time.Second)
			stateUint, err := escrowInstance.GetEscrowState(&bind.CallOpts{Context: callCtx})
			callCancel()
			if err != nil {
				log.Errorf("Failed to get EscrowState from contract: %v", err)
				continue
			}

			mappedState := mapEscrowStateToString(stateUint)
			log.Infof("Detected RevertedDealBreak! Reverter: %s. Restored state: %s", event.Reverter.Hex(), mappedState)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": mappedState,
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to %s", contractAddressHex, mappedState)
			}
		}
	}
}

func watchDealBroken(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowDealBroken)
	sub, err := escrowInstance.WatchFreelanceEscrowDealBroken(&bind.WatchOpts{Context: context.Background()}, sink)
	if err != nil {
		log.Errorf("Failed to subscribe to DealBroken event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to DealBroken events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("DealBroken subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected DealBroken! Timestamp: %s", event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "DEAL_BROKEN",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to DEAL_BROKEN", contractAddressHex)
			}
		}
	}
}

func watchRandomDisputeRaised(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
	sub, err := escrowInstance.WatchFreelanceEscrowRandomDisputeRaised(&bind.WatchOpts{Context: context.Background()}, sink, nil)
	if err != nil {
		log.Errorf("Failed to subscribe to RandomDisputeRaised event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to RandomDisputeRaised events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("RandomDisputeRaised subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected RandomDisputeRaised! Raiser: %s", event.Raiser.Hex())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "RANDOM_DISPUTED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to RANDOM_DISPUTED", contractAddressHex)
			}
		}
	}
}

func watchPaymentDisputeRaised(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
	sub, err := escrowInstance.WatchFreelanceEscrowPaymentDisputeRaised(&bind.WatchOpts{Context: context.Background()}, sink)
	if err != nil {
		log.Errorf("Failed to subscribe to PaymentDisputeRaised event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to PaymentDisputeRaised events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("PaymentDisputeRaised subscription error: %v", err)
			return
		case event := <-sink:
			log.Infof("Detected PaymentDisputeRaised! Timestamp: %s", event.Timestamp.String())

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": "PAYMENT_DISPUTED",
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to PAYMENT_DISPUTED", contractAddressHex)
			}
		}
	}
}

func watchDisputeResolved(jobsColl *mongo.Collection, escrowInstance *escrow.FreelanceEscrow, contractAddressHex string) {
	sink := make(chan *escrow.FreelanceEscrowFreelanceEscrowDisputeResolved)
	sub, err := escrowInstance.WatchFreelanceEscrowDisputeResolved(&bind.WatchOpts{Context: context.Background()}, sink)
	if err != nil {
		log.Errorf("Failed to subscribe to DisputeResolved event: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Infof("Successfully subscribed to DisputeResolved events at %s", contractAddressHex)

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("DisputeResolved subscription error: %v", err)
			return
		case event := <-sink:
			callCtx, callCancel := context.WithTimeout(context.Background(), 5*time.Second)
			stateUint, err := escrowInstance.GetEscrowState(&bind.CallOpts{Context: callCtx})
			callCancel()
			if err != nil {
				log.Errorf("Failed to get EscrowState from contract: %v", err)
				continue
			}

			mappedState := mapEscrowStateToString(stateUint)
			log.Infof("Detected DisputeResolved! Timestamp: %s. Restored state: %s", event.Timestamp.String(), mappedState)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

			filter := bson.M{"contractAddress": contractAddressHex}
			update := bson.M{
				"$set": bson.M{
					"status": mappedState,
				},
			}

			res, err := jobsColl.UpdateOne(ctx, filter, update)
			cancel()

			if err != nil {
				log.Errorf("Error updating MongoDB for Job at %s: %v", contractAddressHex, err)
			} else if res.MatchedCount == 0 {
				log.Warnf("No MongoDB Job found matching contract address %s", contractAddressHex)
			} else {
				log.Infof("Successfully updated MongoDB Job %s to %s", contractAddressHex, mappedState)
			}
		}
	}
}
