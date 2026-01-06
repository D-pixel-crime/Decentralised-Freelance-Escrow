// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ReentrancyGuard} from "reentrancy-guard/ReentrancyGuard.sol";

/**
 * @title Freelance-Escrow
 * @author D-Pixel-Crime
 * @notice Decentralised solution for trustless freelancing
 */
contract FreelanceEscrow is ReentrancyGuard {
    enum EscrowState {
        AGREED,
        CLIENT_STAKED,
        FREELANCER_STAKED,
        ALL_STAKED_AND_PENDING,
        PENDING_CLIENT_CONFIRMATION,
        JOB_COMPLETED,
        CANCEL_REQUESTED,
        DEAL_BROKEN,
        RANDOM_DISPUTED,
        PAYMENT_DISPUTED
    }
    address private immutable i_owner;
    address private immutable i_client;
    address private immutable i_freelancer;
    address private immutable i_arbitrator;
    uint256 private immutable i_jobId;
    address private s_dealBreakInitiator = address(0);
    uint256 private s_clientStake = 0;
    uint256 private s_freelancerStake = 0;
    uint256 private s_unilateralCompletionTimestamp;
    uint256 private s_confirmationPeriod;
    EscrowState private currState;
    EscrowState private prevState;

    error FreelanceEscrow__NotParticipant();
    error FreelanceEscrow__InvalidArbitrator();
    error FreelanceEscrow__InvalidClient();
    error FreelanceEscrow__InvalidFreelancer();
    error FreelanceEscrow__ClientAlreadyStaked(EscrowState);
    error FreelanceEscrow__FreelancerAlreadyStaked(EscrowState);
    error FreenlanceEscrow__DealAlreadyBroken(EscrowState);
    error FreenlanceEscrow__Busy(EscrowState);
    error FreelanceEscrow__RefundError(address, uint256);
    error FreelanceEscrow__ProcessNotAllowed(EscrowState);
    error FreelanceEscrow__PaymentError(uint256, uint256);
    error FreelanceEscrow__NoCancelRequestedYet(EscrowState);
    error FreelanceEscrow__ActiveConfirmationTimePeriod(uint256, uint256);
    error FreelanceEscrow__InvalidFundsDistribution();

    event FreelanceEscrow__ClientStakeCompleted(
        uint256 indexed timestamp,
        uint256 indexed amount
    );
    event FreelanceEscrow__FreelancerStakeCompleted(
        uint256 indexed timestamp,
        uint256 indexed amount
    );
    event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp);
    event FreelanceEscrow__JobCompletedAndFreelancerPaid(
        uint256 indexed amount,
        uint256 indexed timestamp
    );
    event FreelanceEscrow__JobCompletionRejected(uint256 timestamp);
    event FreelanceEscrow__DealCancelRequested(address indexed initiator);
    event FreelanceEscrow__DealBroken(uint256 timestamp);
    event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
        uint256 timestamp
    );
    event FreelanceEscrow__AggreementCreated(
        address indexed client,
        address indexed freelancer
    );
    event FreelanceEscrow__RevertedDealBreak(address indexed reverter);
    event FreelanceEscrow__RandomDisputeRaised(
        address indexed raiser,
        uint256 timestamp
    );
    event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp);
    event FreelanceEscrow__RandomDisputeResolved(uint256 timestamp);

    modifier onlyClient() {
        _onlyClient();
        _;
    }

    modifier onlyFreelancer() {
        _onlyFreelancer();
        _;
    }
    modifier onlyArbitrator() {
        _onlyArbitrator();
        _;
    }

    modifier noDispute() {
        _noDispute();
        _;
    }

    modifier noCancel() {
        _noCancel();
        _;
    }

    constructor(
        uint256 jobId,
        address client,
        address freelancer,
        address arbitrator,
        uint256 confirmationPeriod
    ) {
        i_owner = msg.sender;
        i_jobId = jobId;
        i_client = client;
        i_freelancer = freelancer;
        i_arbitrator = arbitrator;
        currState = EscrowState.AGREED;
        prevState = EscrowState.AGREED;
        s_confirmationPeriod = confirmationPeriod;
        emit FreelanceEscrow__AggreementCreated(client, freelancer);
    }

    function raiseRandomDispute() public nonReentrant noDispute noCancel {
        if (msg.sender != i_client && msg.sender != i_freelancer) {
            revert FreelanceEscrow__NotParticipant();
        }

        prevState = currState;
        currState = EscrowState.RANDOM_DISPUTED;
        emit FreelanceEscrow__RandomDisputeRaised(msg.sender, block.timestamp);
    }

    function resolveDispute(
        uint256 clientPayment,
        uint256 freelancerPayment
    ) public nonReentrant onlyArbitrator noCancel {
        if (
            clientPayment != s_clientStake ||
            freelancerPayment != s_freelancerStake
        ) {
            revert FreelanceEscrow__InvalidFundsDistribution();
        }
        if (clientPayment == 0 && freelancerPayment == 0) {
            if (prevState == EscrowState.PENDING_CLIENT_CONFIRMATION) {
                currState = EscrowState.ALL_STAKED_AND_PENDING;
            } else {
                currState = prevState;
            }
            emit FreelanceEscrow__RandomDisputeResolved(block.timestamp);
            return;
        }

        uint256 payClient = s_clientStake;
        uint256 payFreelancer = s_freelancerStake;

        bool success;
        if (clientPayment != 0) {
            (success, ) = payable(i_client).call{value: payClient}("");
            if (!success) {
                revert FreelanceEscrow__RefundError(i_client, s_clientStake);
            }

            (success, ) = payable(i_freelancer).call{value: payFreelancer}("");
            if (!success) {
                revert FreelanceEscrow__PaymentError(
                    payFreelancer,
                    s_clientStake
                );
            }

            currState = EscrowState.DEAL_BROKEN;
            emit FreelanceEscrow__DealBroken(block.timestamp);
        } else {
            payFreelancer += s_clientStake;

            (success, ) = payable(i_freelancer).call{value: payFreelancer}("");
            if (!success) {
                revert FreelanceEscrow__PaymentError(
                    payFreelancer,
                    s_clientStake
                );
            }

            currState = EscrowState.JOB_COMPLETED;
            emit FreelanceEscrow__JobCompletedAndFreelancerPaid(
                payFreelancer,
                block.timestamp
            );
        }
    }

    function addClientStake()
        public
        payable
        nonReentrant
        onlyClient
        noDispute
        noCancel
    {
        if (currState == EscrowState.AGREED) {
            currState = EscrowState.CLIENT_STAKED;
        } else if (currState == EscrowState.FREELANCER_STAKED) {
            currState = EscrowState.ALL_STAKED_AND_PENDING;
            emit FreelanceEscrow__BothPartyStakeCompleted(block.timestamp);
        } else {
            revert FreelanceEscrow__ClientAlreadyStaked(currState);
        }
        s_clientStake = msg.value;
        emit FreelanceEscrow__ClientStakeCompleted(block.timestamp, msg.value);
    }

    function addfreelancerStake()
        public
        payable
        onlyFreelancer
        nonReentrant
        noDispute
        noCancel
    {
        if (currState == EscrowState.AGREED) {
            currState = EscrowState.FREELANCER_STAKED;
        } else if (currState == EscrowState.CLIENT_STAKED) {
            currState = EscrowState.ALL_STAKED_AND_PENDING;
            emit FreelanceEscrow__BothPartyStakeCompleted(block.timestamp);
        } else {
            revert FreelanceEscrow__FreelancerAlreadyStaked(currState);
        }
        s_freelancerStake = msg.value;
        emit FreelanceEscrow__FreelancerStakeCompleted(
            block.timestamp,
            msg.value
        );
    }

    function requestPayment()
        public
        onlyFreelancer
        nonReentrant
        noDispute
        noCancel
    {
        if (currState != EscrowState.ALL_STAKED_AND_PENDING) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        /** @dev Start a timer of 5 days for Client Confirmations */

        currState = EscrowState.PENDING_CLIENT_CONFIRMATION;
        emit FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
            block.timestamp
        );
    }

    function finaliseUnilateralJob() public nonReentrant noDispute noCancel {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
        if (
            block.timestamp <
            s_unilateralCompletionTimestamp + s_confirmationPeriod
        ) {
            revert FreelanceEscrow__ActiveConfirmationTimePeriod(
                s_unilateralCompletionTimestamp,
                block.timestamp
            );
        }

        uint256 toPay = address(this).balance;

        (bool success, ) = payable(i_freelancer).call{value: toPay}("");
        if (!success) {
            revert FreelanceEscrow__PaymentError(toPay, block.timestamp);
        }
        emit FreelanceEscrow__JobCompletedAndFreelancerPaid(
            toPay,
            block.timestamp
        );
        currState = EscrowState.JOB_COMPLETED;
    }

    function rejectJobCompletion()
        public
        nonReentrant
        onlyClient
        noDispute
        noCancel
    {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        currState = EscrowState.PAYMENT_DISPUTED;
        emit FreelanceEscrow__JobCompletionRejected(block.timestamp);
        emit FreelanceEscrow__PaymentDisputeRaised(block.timestamp);
    }

    function acceptJobCompletion()
        public
        nonReentrant
        onlyClient
        noDispute
        noCancel
    {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        uint256 toPay = address(this).balance;
        s_clientStake = 0;
        s_freelancerStake = 0;

        (bool success, ) = payable(i_freelancer).call{value: toPay}("");
        if (!success) {
            revert FreelanceEscrow__PaymentError(toPay, block.timestamp);
        }
        emit FreelanceEscrow__JobCompletedAndFreelancerPaid(
            toPay,
            block.timestamp
        );
        currState = EscrowState.JOB_COMPLETED;
    }

    function jobCompletionAgreedBilaterally()
        public
        onlyClient
        nonReentrant
        noDispute
        noCancel
    {
        if (currState != EscrowState.ALL_STAKED_AND_PENDING) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        uint256 toPay = address(this).balance;
        s_clientStake = 0;
        s_freelancerStake = 0;

        (bool success, ) = payable(i_freelancer).call{value: toPay}("");
        if (!success) {
            revert FreelanceEscrow__PaymentError(toPay, block.timestamp);
        }
        emit FreelanceEscrow__JobCompletedAndFreelancerPaid(
            toPay,
            block.timestamp
        );
        currState = EscrowState.JOB_COMPLETED;
    }

    function breakDeal() public nonReentrant noDispute noCancel {
        if (
            msg.sender == s_dealBreakInitiator ||
            currState == EscrowState.JOB_COMPLETED
        ) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
        if (currState == EscrowState.DEAL_BROKEN) {
            revert FreenlanceEscrow__DealAlreadyBroken(currState);
        }
        if (msg.sender != i_client && msg.sender != i_freelancer) {
            revert FreelanceEscrow__NotParticipant();
        }

        if (s_dealBreakInitiator == address(0)) {
            prevState = currState;
            s_dealBreakInitiator = msg.sender;
            currState = EscrowState.CANCEL_REQUESTED;
            emit FreelanceEscrow__DealCancelRequested(msg.sender);
            return;
        }

        uint256 toRefund = s_clientStake;
        s_clientStake = 0;
        (bool success, ) = payable(i_client).call{value: toRefund}("");
        if (!success) {
            revert FreelanceEscrow__RefundError(i_client, block.timestamp);
        }

        toRefund = s_freelancerStake;
        s_freelancerStake = 0;
        (success, ) = payable(i_freelancer).call{value: toRefund}("");
        if (!success) {
            revert FreelanceEscrow__RefundError(i_freelancer, block.timestamp);
        }

        currState = EscrowState.DEAL_BROKEN;
        emit FreelanceEscrow__DealBroken(block.timestamp);
    }

    function cancelDealBreak() public nonReentrant noDispute noCancel {
        if (msg.sender != i_client && msg.sender != i_freelancer) {
            revert FreelanceEscrow__NotParticipant();
        }
        if (currState != EscrowState.CANCEL_REQUESTED) {
            revert FreelanceEscrow__NoCancelRequestedYet(currState);
        }

        if (prevState == EscrowState.PENDING_CLIENT_CONFIRMATION) {
            currState = EscrowState.ALL_STAKED_AND_PENDING;
        } else {
            currState = prevState;
        }

        s_dealBreakInitiator = address(0);
        emit FreelanceEscrow__RevertedDealBreak(msg.sender);
    }

    /** @dev Getter Functions */
    function getEscrowState() public view returns (EscrowState) {
        return currState;
    }

    function getClientStake() public view returns (uint256) {
        return s_clientStake;
    }

    function getFreelancerStake() public view returns (uint256) {
        return s_freelancerStake;
    }

    function getConfirmationPeriod() public view returns (uint256) {
        return s_confirmationPeriod;
    }

    /** @dev Internal Functions */
    function _onlyClient() public view {
        if (msg.sender != i_client) {
            revert FreelanceEscrow__InvalidClient();
        }
    }

    function _onlyFreelancer() public view {
        if (msg.sender != i_freelancer) {
            revert FreelanceEscrow__InvalidFreelancer();
        }
    }

    function _onlyArbitrator() public view {
        if (msg.sender != i_arbitrator) {
            revert FreelanceEscrow__InvalidArbitrator();
        }
    }

    function _noDispute() public view {
        if (
            currState == EscrowState.RANDOM_DISPUTED ||
            currState == EscrowState.PAYMENT_DISPUTED
        ) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
    }

    function _noCancel() public view {
        if (
            currState == EscrowState.CANCEL_REQUESTED ||
            currState == EscrowState.DEAL_BROKEN
        ) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
    }
}
