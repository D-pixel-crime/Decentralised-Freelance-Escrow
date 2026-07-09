// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FreelanceEscrowMetaData contains all meta data concerning the FreelanceEscrow contract.
var FreelanceEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"confirmationPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_noCancel\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_noDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyArbitrator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyClient\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyFreelancer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addClientStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addfreelancerStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"breakDeal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDealBreak\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finaliseUnilateralJob\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmationPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFreelancerStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestPayment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"clientPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"freelancerPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FreelanceEscrow__AggreementCreated\",\"inputs\":[{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__BothPartyStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__ClientStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealBroken\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealCancelRequested\",\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DisputeResolved\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletedAndFreelancerPaid\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletionRejected\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__PaymentDisputeRaised\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RandomDisputeRaised\",\"inputs\":[{\"name\":\"raiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RevertedDealBreak\",\"inputs\":[{\"name\":\"reverter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ActiveConfirmationTimePeriod\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ClientAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__FreelancerAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidClient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFreelancer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFundsDistribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NoCancelRequestedYet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__PaymentError\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ProcessNotAllowed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__RefundError\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__Busy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__DealAlreadyBroken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6101206040525f5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f6001555f600255348015610057575f5ffd5b506040516127ec3803806127ec833981810160405281019061007991906102f7565b600161009761008c61023460201b60201c565b61025d60201b60201c565b5f01819055503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508461010081815250508373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff16815250505f60055f6101000a81548160ff0219169083600981111561019a5761019961036e565b5b02179055505f600560016101000a81548160ff021916908360098111156101c4576101c361036e565b5b0217905550806004819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f60405160405180910390a3505050505061039b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f5ffd5b5f819050919050565b61027c8161026a565b8114610286575f5ffd5b50565b5f8151905061029781610273565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102c68261029d565b9050919050565b6102d6816102bc565b81146102e0575f5ffd5b50565b5f815190506102f1816102cd565b92915050565b5f5f5f5f5f60a086880312156103105761030f610266565b5b5f61031d88828901610289565b955050602061032e888289016102e3565b945050604061033f888289016102e3565b9350506060610350888289016102e3565b925050608061036188828901610289565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60805160a05160c05160e051610100516123aa6104425f395f50505f61132e01525f8181610529015281816106a901528181610ab601528181610d7f01528181610e0d01528181610f49015281816110f20152818161195c01528181611aa10152611d9701525f81816102f901528181610a5f01528181610c8801528181610d1401528181610ef20152818161109b0152818161185e01526118ec01525f50506123aa5ff3fe608060405260043610610113575f3560e01c806387d40b5a1161009f578063bdc84ac311610063578063bdc84ac314610263578063d31cdf891461028b578063dd33869f146102a1578063df44be2c146102b7578063ee4b7ab8146102e157610113565b806387d40b5a146101c5578063939ae273146101ef57806395744869146101f957806397c9522e1461020f578063ac755a631461023957610113565b806363bdb94b116100e657806363bdb94b146101575780636977e1761461016d5780636daa2d441461018357806375f84b40146101995780637f25e289146101af57610113565b806308e0771e1461011757806324dcdf641461012d5780633e216754146101375780634e07c9d11461014d575b5f5ffd5b348015610122575f5ffd5b5061012b6102f7565b005b61013561037e565b005b348015610142575f5ffd5b5061014b610527565b005b6101556105ae565b005b348015610162575f5ffd5b5061016b6107d5565b005b348015610178575f5ffd5b506101816108ec565b005b34801561018e575f5ffd5b50610197610ed8565b005b3480156101a4575f5ffd5b506101ad611089565b005b3480156101ba575f5ffd5b506101c361132c565b005b3480156101d0575f5ffd5b506101d96113b3565b6040516101e69190612102565b60405180910390f35b6101f76113bc565b005b348015610204575f5ffd5b5061020d611565565b005b34801561021a575f5ffd5b50610223611629565b6040516102309190612102565b60405180910390f35b348015610244575f5ffd5b5061024d611632565b60405161025a9190612102565b60405180910390f35b34801561026e575f5ffd5b5061028960048036038101906102849190612149565b61163b565b005b348015610296575f5ffd5b5061029f611bd6565b005b3480156102ac575f5ffd5b506102b5611c9b565b005b3480156102c2575f5ffd5b506102cb611ec3565b6040516102d891906121fa565b60405180910390f35b3480156102ec575f5ffd5b506102f5611ed8565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461037c576040517fe81d887300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610386610527565b61038e61201f565b610396611565565b61039e611bd6565b5f60098111156103b1576103b0612187565b5b60055f9054906101000a900460ff1660098111156103d2576103d1612187565b5b0361040657600260055f6101000a81548160ff021916908360098111156103fc576103fb612187565b5b02179055506104e8565b6001600981111561041a57610419612187565b5b60055f9054906101000a900460ff16600981111561043b5761043a612187565b5b0361049c57600360055f6101000a81548160ff0219169083600981111561046557610464612187565b5b0217905550427f84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da60405160405180910390a26104e7565b60055f9054906101000a900460ff166040517f16a523f80000000000000000000000000000000000000000000000000000000081526004016104de91906121fa565b60405180910390fd5b5b3460028190555034427f4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e60405160405180910390a3610525612041565b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ac576040517f2240971400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6105b661201f565b6105be6102f7565b6105c6611565565b6105ce611bd6565b600460098111156105e2576105e1612187565b5b60055f9054906101000a900460ff16600981111561060357610602612187565b5b141580156106445750600360098111156106205761061f612187565b5b60055f9054906101000a900460ff16600981111561064157610640612187565b5b14155b156106945760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161068b91906121fa565b60405180910390fd5b5f4790505f6001819055505f6002819055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16826040516106eb90612240565b5f6040518083038185875af1925050503d805f8114610725576040519150601f19603f3d011682016040523d82523d5f602084013e61072a565b606091505b50509050806107725781426040517f3f678334000000000000000000000000000000000000000000000000000000008152600401610769929190612254565b60405180910390fd5b42827f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a36005805f6101000a81548160ff021916908360098111156107c4576107c3612187565b5b021790555050506107d3612041565b565b6107dd610527565b6107e561201f565b6107ed611565565b6107f5611bd6565b6003600981111561080957610808612187565b5b60055f9054906101000a900460ff16600981111561082a57610829612187565b5b1461087a5760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161087191906121fa565b60405180910390fd5b600460055f6101000a81548160ff0219169083600981111561089f5761089e612187565b5b0217905550426003819055507ffbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6426040516108da9190612102565b60405180910390a16108ea612041565b565b6108f461201f565b6108fc611565565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061098857506005600981111561096557610964612187565b5b60055f9054906101000a900460ff16600981111561098657610985612187565b5b145b156109d85760055f9054906101000a900460ff166040517f361057570000000000000000000000000000000000000000000000000000000081526004016109cf91906121fa565b60405180910390fd5b600760098111156109ec576109eb612187565b5b60055f9054906101000a900460ff166009811115610a0d57610a0c612187565b5b03610a5d5760055f9054906101000a900460ff166040517f9153e0e2000000000000000000000000000000000000000000000000000000008152600401610a5491906121fa565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610b0557507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610b3c576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c785760055f9054906101000a900460ff16600560016101000a81548160ff02191690836009811115610bc257610bc1612187565b5b0217905550335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660055f6101000a81548160ff02191690836009811115610c2b57610c2a612187565b5b02179055503373ffffffffffffffffffffffffffffffffffffffff167ff0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce60405160405180910390a2610ece565b5f60015490505f6001819055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051610cca90612240565b5f6040518083038185875af1925050503d805f8114610d04576040519150601f19603f3d011682016040523d82523d5f602084013e610d09565b606091505b5050905080610d71577f0000000000000000000000000000000000000000000000000000000000000000426040517fc3940718000000000000000000000000000000000000000000000000000000008152600401610d689291906122ba565b60405180910390fd5b60025491505f6002819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051610dc190612240565b5f6040518083038185875af1925050503d805f8114610dfb576040519150601f19603f3d011682016040523d82523d5f602084013e610e00565b606091505b50508091505080610e6a577f0000000000000000000000000000000000000000000000000000000000000000426040517fc3940718000000000000000000000000000000000000000000000000000000008152600401610e619291906122ba565b60405180910390fd5b600760055f6101000a81548160ff02191690836009811115610e8f57610e8e612187565b5b02179055507fa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf2442604051610ec39190612102565b60405180910390a150505b610ed6612041565b565b610ee061201f565b610ee8611565565b610ef0611bd6565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610f9857507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610fcf576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055f9054906101000a900460ff16600560016101000a81548160ff0219169083600981111561100257611001612187565b5b0217905550600860055f6101000a81548160ff0219169083600981111561102c5761102b612187565b5b02179055503373ffffffffffffffffffffffffffffffffffffffff167f022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991426040516110779190612102565b60405180910390a2611087612041565b565b61109161201f565b611099611565565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561114157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15611178576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006600981111561118c5761118b612187565b5b60055f9054906101000a900460ff1660098111156111ad576111ac612187565b5b146111fd5760055f9054906101000a900460ff166040517feffe824c0000000000000000000000000000000000000000000000000000000081526004016111f491906121fa565b60405180910390fd5b6004600981111561121157611210612187565b5b600560019054906101000a900460ff16600981111561123357611232612187565b5b0361126757600360055f6101000a81548160ff0219169083600981111561125d5761125c612187565b5b02179055506112a0565b600560019054906101000a900460ff1660055f6101000a81548160ff0219169083600981111561129a57611299612187565b5b02179055505b5f5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a860405160405180910390a261132a612041565b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113b1576040517f5bffd58900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f600154905090565b6113c461201f565b6113cc6102f7565b6113d4611565565b6113dc611bd6565b5f60098111156113ef576113ee612187565b5b60055f9054906101000a900460ff1660098111156114105761140f612187565b5b0361144457600160055f6101000a81548160ff0219169083600981111561143a57611439612187565b5b0217905550611526565b6002600981111561145857611457612187565b5b60055f9054906101000a900460ff16600981111561147957611478612187565b5b036114da57600360055f6101000a81548160ff021916908360098111156114a3576114a2612187565b5b0217905550427f84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da60405160405180910390a2611525565b60055f9054906101000a900460ff166040517f92b538e000000000000000000000000000000000000000000000000000000000815260040161151c91906121fa565b60405180910390fd5b5b3460018190555034427f218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc60405160405180910390a3611563612041565b565b6008600981111561157957611578612187565b5b60055f9054906101000a900460ff16600981111561159a57611599612187565b5b14806115d757506009808111156115b4576115b3612187565b5b60055f9054906101000a900460ff1660098111156115d5576115d4612187565b5b145b156116275760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161161e91906121fa565b60405180910390fd5b565b5f600454905090565b5f600254905090565b61164361201f565b61164b61132c565b611653611bd6565b6008600981111561166757611666612187565b5b60055f9054906101000a900460ff16600981111561168857611687612187565b5b141580156116c857506009808111156116a4576116a3612187565b5b60055f9054906101000a900460ff1660098111156116c5576116c4612187565b5b14155b156117185760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161170f91906121fa565b60405180910390fd5b5f4790508183611728919061230e565b811015611761576040517f4aa5857a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8314801561176f57505f82145b15611854576004600981111561178857611787612187565b5b600560019054906101000a900460ff1660098111156117aa576117a9612187565b5b036117de57600360055f6101000a81548160ff021916908360098111156117d4576117d3612187565b5b0217905550611817565b600560019054906101000a900460ff1660055f6101000a81548160ff0219169083600981111561181157611810612187565b5b02179055505b7f0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e426040516118469190612102565b60405180910390a150611bca565b5f5f8414611a90577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16846040516118a090612240565b5f6040518083038185875af1925050503d805f81146118da576040519150601f19603f3d011682016040523d82523d5f602084013e6118df565b606091505b5050809150508061194b577f00000000000000000000000000000000000000000000000000000000000000006001546040517fc39407180000000000000000000000000000000000000000000000000000000081526004016119429291906122ba565b60405180910390fd5b5f84836119589190612341565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168160405161199e90612240565b5f6040518083038185875af1925050503d805f81146119d8576040519150601f19603f3d011682016040523d82523d5f602084013e6119dd565b606091505b50508092505081611a2957806001546040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611a20929190612254565b60405180910390fd5b600760055f6101000a81548160ff02191690836009811115611a4e57611a4d612187565b5b02179055507fa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf2442604051611a829190612102565b60405180910390a150611bc7565b5f8483611a9d9190612341565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681604051611ae390612240565b5f6040518083038185875af1925050503d805f8114611b1d576040519150601f19603f3d011682016040523d82523d5f602084013e611b22565b606091505b50508092505081611b6e57806001546040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611b65929190612254565b60405180910390fd5b6005805f6101000a81548160ff02191690836009811115611b9257611b91612187565b5b021790555042817f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a3505b50505b611bd2612041565b5050565b60066009811115611bea57611be9612187565b5b60055f9054906101000a900460ff166009811115611c0b57611c0a612187565b5b1480611c49575060076009811115611c2657611c25612187565b5b60055f9054906101000a900460ff166009811115611c4757611c46612187565b5b145b15611c995760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611c9091906121fa565b60405180910390fd5b565b611ca361201f565b611cab611565565b611cb3611bd6565b60046009811115611cc757611cc6612187565b5b60055f9054906101000a900460ff166009811115611ce857611ce7612187565b5b14611d385760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611d2f91906121fa565b60405180910390fd5b600454600354611d48919061230e565b421015611d9057600354426040517fc3abe3e2000000000000000000000000000000000000000000000000000000008152600401611d87929190612254565b60405180910390fd5b5f4790505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051611dd990612240565b5f6040518083038185875af1925050503d805f8114611e13576040519150601f19603f3d011682016040523d82523d5f602084013e611e18565b606091505b5050905080611e605781426040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611e57929190612254565b60405180910390fd5b42827f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a36005805f6101000a81548160ff02191690836009811115611eb257611eb1612187565b5b02179055505050611ec1612041565b565b5f60055f9054906101000a900460ff16905090565b611ee061201f565b611ee86102f7565b611ef0611565565b611ef8611bd6565b60046009811115611f0c57611f0b612187565b5b60055f9054906101000a900460ff166009811115611f2d57611f2c612187565b5b14611f7d5760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611f7491906121fa565b60405180910390fd5b600960055f6101000a81548160ff02191690836009811115611fa257611fa1612187565b5b02179055507f52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd42604051611fd69190612102565b60405180910390a17f7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e4260405161200d9190612102565b60405180910390a161201d612041565b565b61202761205b565b600261203961203461209c565b6120c5565b5f0181905550565b600161205361204e61209c565b6120c5565b5f0181905550565b6120636120ce565b1561209a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f60026120e16120dc61209c565b6120c5565b5f015414905090565b5f819050919050565b6120fc816120ea565b82525050565b5f6020820190506121155f8301846120f3565b92915050565b5f5ffd5b612128816120ea565b8114612132575f5ffd5b50565b5f813590506121438161211f565b92915050565b5f5f6040838503121561215f5761215e61211b565b5b5f61216c85828601612135565b925050602061217d85828601612135565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600a81106121c5576121c4612187565b5b50565b5f8190506121d5826121b4565b919050565b5f6121e4826121c8565b9050919050565b6121f4816121da565b82525050565b5f60208201905061220d5f8301846121eb565b92915050565b5f81905092915050565b50565b5f61222b5f83612213565b91506122368261221d565b5f82019050919050565b5f61224a82612220565b9150819050919050565b5f6040820190506122675f8301856120f3565b61227460208301846120f3565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6122a48261227b565b9050919050565b6122b48161229a565b82525050565b5f6040820190506122cd5f8301856122ab565b6122da60208301846120f3565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612318826120ea565b9150612323836120ea565b925082820190508082111561233b5761233a6122e1565b5b92915050565b5f61234b826120ea565b9150612356836120ea565b925082820390508181111561236e5761236d6122e1565b5b9291505056fea2646970667358221220022807ba0fe749e689d765520aa3b73dbd38dbd5d0c725cfc716eb8d9f5d18c664736f6c63430008210033",
}

// FreelanceEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use FreelanceEscrowMetaData.ABI instead.
var FreelanceEscrowABI = FreelanceEscrowMetaData.ABI

// FreelanceEscrowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FreelanceEscrowMetaData.Bin instead.
var FreelanceEscrowBin = FreelanceEscrowMetaData.Bin

// DeployFreelanceEscrow deploys a new Ethereum contract, binding an instance of FreelanceEscrow to it.
func DeployFreelanceEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, jobId *big.Int, client common.Address, freelancer common.Address, arbitrator common.Address, confirmationPeriod *big.Int) (common.Address, *types.Transaction, *FreelanceEscrow, error) {
	parsed, err := FreelanceEscrowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FreelanceEscrowBin), backend, jobId, client, freelancer, arbitrator, confirmationPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FreelanceEscrow{FreelanceEscrowCaller: FreelanceEscrowCaller{contract: contract}, FreelanceEscrowTransactor: FreelanceEscrowTransactor{contract: contract}, FreelanceEscrowFilterer: FreelanceEscrowFilterer{contract: contract}}, nil
}

// FreelanceEscrow is an auto generated Go binding around an Ethereum contract.
type FreelanceEscrow struct {
	FreelanceEscrowCaller     // Read-only binding to the contract
	FreelanceEscrowTransactor // Write-only binding to the contract
	FreelanceEscrowFilterer   // Log filterer for contract events
}

// FreelanceEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreelanceEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreelanceEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreelanceEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreelanceEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreelanceEscrowSession struct {
	Contract     *FreelanceEscrow  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreelanceEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreelanceEscrowCallerSession struct {
	Contract *FreelanceEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FreelanceEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreelanceEscrowTransactorSession struct {
	Contract     *FreelanceEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FreelanceEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreelanceEscrowRaw struct {
	Contract *FreelanceEscrow // Generic contract binding to access the raw methods on
}

// FreelanceEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreelanceEscrowCallerRaw struct {
	Contract *FreelanceEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// FreelanceEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreelanceEscrowTransactorRaw struct {
	Contract *FreelanceEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreelanceEscrow creates a new instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrow(address common.Address, backend bind.ContractBackend) (*FreelanceEscrow, error) {
	contract, err := bindFreelanceEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrow{FreelanceEscrowCaller: FreelanceEscrowCaller{contract: contract}, FreelanceEscrowTransactor: FreelanceEscrowTransactor{contract: contract}, FreelanceEscrowFilterer: FreelanceEscrowFilterer{contract: contract}}, nil
}

// NewFreelanceEscrowCaller creates a new read-only instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowCaller(address common.Address, caller bind.ContractCaller) (*FreelanceEscrowCaller, error) {
	contract, err := bindFreelanceEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowCaller{contract: contract}, nil
}

// NewFreelanceEscrowTransactor creates a new write-only instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*FreelanceEscrowTransactor, error) {
	contract, err := bindFreelanceEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowTransactor{contract: contract}, nil
}

// NewFreelanceEscrowFilterer creates a new log filterer instance of FreelanceEscrow, bound to a specific deployed contract.
func NewFreelanceEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*FreelanceEscrowFilterer, error) {
	contract, err := bindFreelanceEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFilterer{contract: contract}, nil
}

// bindFreelanceEscrow binds a generic wrapper to an already deployed contract.
func bindFreelanceEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FreelanceEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreelanceEscrow *FreelanceEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreelanceEscrow.Contract.FreelanceEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreelanceEscrow *FreelanceEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FreelanceEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreelanceEscrow *FreelanceEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FreelanceEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreelanceEscrow *FreelanceEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreelanceEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreelanceEscrow *FreelanceEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreelanceEscrow *FreelanceEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.contract.Transact(opts, method, params...)
}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) NoCancel(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_noCancel")

	if err != nil {
		return err
	}

	return err

}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) NoCancel() error {
	return _FreelanceEscrow.Contract.NoCancel(&_FreelanceEscrow.CallOpts)
}

// NoCancel is a free data retrieval call binding the contract method 0xd31cdf89.
//
// Solidity: function _noCancel() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) NoCancel() error {
	return _FreelanceEscrow.Contract.NoCancel(&_FreelanceEscrow.CallOpts)
}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) NoDispute(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_noDispute")

	if err != nil {
		return err
	}

	return err

}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) NoDispute() error {
	return _FreelanceEscrow.Contract.NoDispute(&_FreelanceEscrow.CallOpts)
}

// NoDispute is a free data retrieval call binding the contract method 0x95744869.
//
// Solidity: function _noDispute() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) NoDispute() error {
	return _FreelanceEscrow.Contract.NoDispute(&_FreelanceEscrow.CallOpts)
}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyArbitrator(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyArbitrator")

	if err != nil {
		return err
	}

	return err

}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyArbitrator() error {
	return _FreelanceEscrow.Contract.OnlyArbitrator(&_FreelanceEscrow.CallOpts)
}

// OnlyArbitrator is a free data retrieval call binding the contract method 0x7f25e289.
//
// Solidity: function _onlyArbitrator() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyArbitrator() error {
	return _FreelanceEscrow.Contract.OnlyArbitrator(&_FreelanceEscrow.CallOpts)
}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyClient(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyClient")

	if err != nil {
		return err
	}

	return err

}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyClient() error {
	return _FreelanceEscrow.Contract.OnlyClient(&_FreelanceEscrow.CallOpts)
}

// OnlyClient is a free data retrieval call binding the contract method 0x08e0771e.
//
// Solidity: function _onlyClient() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyClient() error {
	return _FreelanceEscrow.Contract.OnlyClient(&_FreelanceEscrow.CallOpts)
}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowCaller) OnlyFreelancer(opts *bind.CallOpts) error {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "_onlyFreelancer")

	if err != nil {
		return err
	}

	return err

}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowSession) OnlyFreelancer() error {
	return _FreelanceEscrow.Contract.OnlyFreelancer(&_FreelanceEscrow.CallOpts)
}

// OnlyFreelancer is a free data retrieval call binding the contract method 0x3e216754.
//
// Solidity: function _onlyFreelancer() view returns()
func (_FreelanceEscrow *FreelanceEscrowCallerSession) OnlyFreelancer() error {
	return _FreelanceEscrow.Contract.OnlyFreelancer(&_FreelanceEscrow.CallOpts)
}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetClientStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getClientStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetClientStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetClientStake(&_FreelanceEscrow.CallOpts)
}

// GetClientStake is a free data retrieval call binding the contract method 0x87d40b5a.
//
// Solidity: function getClientStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetClientStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetClientStake(&_FreelanceEscrow.CallOpts)
}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetConfirmationPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getConfirmationPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetConfirmationPeriod() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetConfirmationPeriod(&_FreelanceEscrow.CallOpts)
}

// GetConfirmationPeriod is a free data retrieval call binding the contract method 0x97c9522e.
//
// Solidity: function getConfirmationPeriod() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetConfirmationPeriod() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetConfirmationPeriod(&_FreelanceEscrow.CallOpts)
}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetEscrowState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getEscrowState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowSession) GetEscrowState() (uint8, error) {
	return _FreelanceEscrow.Contract.GetEscrowState(&_FreelanceEscrow.CallOpts)
}

// GetEscrowState is a free data retrieval call binding the contract method 0xdf44be2c.
//
// Solidity: function getEscrowState() view returns(uint8)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetEscrowState() (uint8, error) {
	return _FreelanceEscrow.Contract.GetEscrowState(&_FreelanceEscrow.CallOpts)
}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCaller) GetFreelancerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreelanceEscrow.contract.Call(opts, &out, "getFreelancerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowSession) GetFreelancerStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetFreelancerStake(&_FreelanceEscrow.CallOpts)
}

// GetFreelancerStake is a free data retrieval call binding the contract method 0xac755a63.
//
// Solidity: function getFreelancerStake() view returns(uint256)
func (_FreelanceEscrow *FreelanceEscrowCallerSession) GetFreelancerStake() (*big.Int, error) {
	return _FreelanceEscrow.Contract.GetFreelancerStake(&_FreelanceEscrow.CallOpts)
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AcceptJobCompletion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "acceptJobCompletion")
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() payable returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AcceptJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AcceptJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AcceptJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AcceptJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AddClientStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "addClientStake")
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AddClientStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddClientStake(&_FreelanceEscrow.TransactOpts)
}

// AddClientStake is a paid mutator transaction binding the contract method 0x939ae273.
//
// Solidity: function addClientStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AddClientStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddClientStake(&_FreelanceEscrow.TransactOpts)
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AddfreelancerStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "addfreelancerStake")
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AddfreelancerStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddfreelancerStake(&_FreelanceEscrow.TransactOpts)
}

// AddfreelancerStake is a paid mutator transaction binding the contract method 0x24dcdf64.
//
// Solidity: function addfreelancerStake() payable returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) AddfreelancerStake() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AddfreelancerStake(&_FreelanceEscrow.TransactOpts)
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) BreakDeal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "breakDeal")
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) BreakDeal() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.BreakDeal(&_FreelanceEscrow.TransactOpts)
}

// BreakDeal is a paid mutator transaction binding the contract method 0x6977e176.
//
// Solidity: function breakDeal() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) BreakDeal() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.BreakDeal(&_FreelanceEscrow.TransactOpts)
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) CancelDealBreak(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "cancelDealBreak")
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) CancelDealBreak() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.CancelDealBreak(&_FreelanceEscrow.TransactOpts)
}

// CancelDealBreak is a paid mutator transaction binding the contract method 0x75f84b40.
//
// Solidity: function cancelDealBreak() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) CancelDealBreak() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.CancelDealBreak(&_FreelanceEscrow.TransactOpts)
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) FinaliseUnilateralJob(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "finaliseUnilateralJob")
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) FinaliseUnilateralJob() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FinaliseUnilateralJob(&_FreelanceEscrow.TransactOpts)
}

// FinaliseUnilateralJob is a paid mutator transaction binding the contract method 0xdd33869f.
//
// Solidity: function finaliseUnilateralJob() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) FinaliseUnilateralJob() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.FinaliseUnilateralJob(&_FreelanceEscrow.TransactOpts)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RaiseDispute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "raiseDispute")
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RaiseDispute() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RaiseDispute(&_FreelanceEscrow.TransactOpts)
}

// RaiseDispute is a paid mutator transaction binding the contract method 0x6daa2d44.
//
// Solidity: function raiseDispute() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RaiseDispute() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RaiseDispute(&_FreelanceEscrow.TransactOpts)
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RejectJobCompletion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "rejectJobCompletion")
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RejectJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RejectJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// RejectJobCompletion is a paid mutator transaction binding the contract method 0xee4b7ab8.
//
// Solidity: function rejectJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RejectJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RejectJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) RequestPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "requestPayment")
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) RequestPayment() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RequestPayment(&_FreelanceEscrow.TransactOpts)
}

// RequestPayment is a paid mutator transaction binding the contract method 0x63bdb94b.
//
// Solidity: function requestPayment() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) RequestPayment() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.RequestPayment(&_FreelanceEscrow.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) ResolveDispute(opts *bind.TransactOpts, clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "resolveDispute", clientPayment, freelancerPayment)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowSession) ResolveDispute(clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.ResolveDispute(&_FreelanceEscrow.TransactOpts, clientPayment, freelancerPayment)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0xbdc84ac3.
//
// Solidity: function resolveDispute(uint256 clientPayment, uint256 freelancerPayment) returns()
func (_FreelanceEscrow *FreelanceEscrowTransactorSession) ResolveDispute(clientPayment *big.Int, freelancerPayment *big.Int) (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.ResolveDispute(&_FreelanceEscrow.TransactOpts, clientPayment, freelancerPayment)
}

// FreelanceEscrowFreelanceEscrowAggreementCreatedIterator is returned from FilterFreelanceEscrowAggreementCreated and is used to iterate over the raw logs and unpacked data for FreelanceEscrowAggreementCreated events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowAggreementCreatedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowAggreementCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowAggreementCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowAggreementCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowAggreementCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowAggreementCreated represents a FreelanceEscrowAggreementCreated event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowAggreementCreated struct {
	Client     common.Address
	Freelancer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowAggreementCreated is a free log retrieval operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowAggreementCreated(opts *bind.FilterOpts, client []common.Address, freelancer []common.Address) (*FreelanceEscrowFreelanceEscrowAggreementCreatedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var freelancerRule []interface{}
	for _, freelancerItem := range freelancer {
		freelancerRule = append(freelancerRule, freelancerItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__AggreementCreated", clientRule, freelancerRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowAggreementCreatedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__AggreementCreated", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowAggreementCreated is a free log subscription operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowAggreementCreated(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowAggreementCreated, client []common.Address, freelancer []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var freelancerRule []interface{}
	for _, freelancerItem := range freelancer {
		freelancerRule = append(freelancerRule, freelancerItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__AggreementCreated", clientRule, freelancerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowAggreementCreated)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__AggreementCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowAggreementCreated is a log parse operation binding the contract event 0xdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f.
//
// Solidity: event FreelanceEscrow__AggreementCreated(address indexed client, address indexed freelancer)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowAggreementCreated(log types.Log) (*FreelanceEscrowFreelanceEscrowAggreementCreated, error) {
	event := new(FreelanceEscrowFreelanceEscrowAggreementCreated)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__AggreementCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator is returned from FilterFreelanceEscrowBothPartyStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowBothPartyStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted represents a FreelanceEscrowBothPartyStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowBothPartyStakeCompleted is a free log retrieval operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowBothPartyStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int) (*FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__BothPartyStakeCompleted", timestampRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowBothPartyStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__BothPartyStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowBothPartyStakeCompleted is a free log subscription operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowBothPartyStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__BothPartyStakeCompleted", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__BothPartyStakeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowBothPartyStakeCompleted is a log parse operation binding the contract event 0x84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da.
//
// Solidity: event FreelanceEscrow__BothPartyStakeCompleted(uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowBothPartyStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowBothPartyStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__BothPartyStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator is returned from FilterFreelanceEscrowClientStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowClientStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowClientStakeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowClientStakeCompleted represents a FreelanceEscrowClientStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowClientStakeCompleted struct {
	Timestamp *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowClientStakeCompleted is a free log retrieval operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowClientStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int, amount []*big.Int) (*FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__ClientStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowClientStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__ClientStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowClientStakeCompleted is a free log subscription operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowClientStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowClientStakeCompleted, timestamp []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__ClientStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__ClientStakeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowClientStakeCompleted is a log parse operation binding the contract event 0x218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc.
//
// Solidity: event FreelanceEscrow__ClientStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowClientStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowClientStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowClientStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__ClientStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDealBrokenIterator is returned from FilterFreelanceEscrowDealBroken and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDealBroken events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealBrokenIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDealBroken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDealBroken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowDealBroken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDealBrokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDealBroken represents a FreelanceEscrowDealBroken event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealBroken struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDealBroken is a free log retrieval operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDealBroken(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowDealBrokenIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DealBroken")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDealBrokenIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DealBroken", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDealBroken is a free log subscription operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDealBroken(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDealBroken) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DealBroken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDealBroken)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealBroken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowDealBroken is a log parse operation binding the contract event 0xa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf24.
//
// Solidity: event FreelanceEscrow__DealBroken(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDealBroken(log types.Log) (*FreelanceEscrowFreelanceEscrowDealBroken, error) {
	event := new(FreelanceEscrowFreelanceEscrowDealBroken)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealBroken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator is returned from FilterFreelanceEscrowDealCancelRequested and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDealCancelRequested events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDealCancelRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDealCancelRequested represents a FreelanceEscrowDealCancelRequested event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDealCancelRequested struct {
	Initiator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDealCancelRequested is a free log retrieval operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDealCancelRequested(opts *bind.FilterOpts, initiator []common.Address) (*FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DealCancelRequested", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDealCancelRequestedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DealCancelRequested", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDealCancelRequested is a free log subscription operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDealCancelRequested(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDealCancelRequested, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DealCancelRequested", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealCancelRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowDealCancelRequested is a log parse operation binding the contract event 0xf0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce.
//
// Solidity: event FreelanceEscrow__DealCancelRequested(address indexed initiator)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDealCancelRequested(log types.Log) (*FreelanceEscrowFreelanceEscrowDealCancelRequested, error) {
	event := new(FreelanceEscrowFreelanceEscrowDealCancelRequested)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DealCancelRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowDisputeResolvedIterator is returned from FilterFreelanceEscrowDisputeResolved and is used to iterate over the raw logs and unpacked data for FreelanceEscrowDisputeResolved events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDisputeResolvedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowDisputeResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowDisputeResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowDisputeResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowDisputeResolved represents a FreelanceEscrowDisputeResolved event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowDisputeResolved struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowDisputeResolved is a free log retrieval operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowDisputeResolved(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowDisputeResolvedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowDisputeResolvedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowDisputeResolved is a free log subscription operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowDisputeResolved(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowDisputeResolved)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DisputeResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowDisputeResolved is a log parse operation binding the contract event 0x0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e.
//
// Solidity: event FreelanceEscrow__DisputeResolved(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowDisputeResolved(log types.Log) (*FreelanceEscrowFreelanceEscrowDisputeResolved, error) {
	event := new(FreelanceEscrowFreelanceEscrowDisputeResolved)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator is returned from FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending and is used to iterate over the raw logs and unpacked data for FreelanceEscrowFreelancerCompletedAndClientConfirmationPending events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator struct {
	Event *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending represents a FreelanceEscrowFreelancerCompletedAndClientConfirmationPending event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a free log retrieval operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPendingIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a free log subscription operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowFreelancerCompletedAndClientConfirmationPending is a log parse operation binding the contract event 0xfbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6.
//
// Solidity: event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowFreelancerCompletedAndClientConfirmationPending(log types.Log) (*FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending, error) {
	event := new(FreelanceEscrowFreelanceEscrowFreelancerCompletedAndClientConfirmationPending)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator is returned from FilterFreelanceEscrowFreelancerStakeCompleted and is used to iterate over the raw logs and unpacked data for FreelanceEscrowFreelancerStakeCompleted events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted represents a FreelanceEscrowFreelancerStakeCompleted event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted struct {
	Timestamp *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowFreelancerStakeCompleted is a free log retrieval operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowFreelancerStakeCompleted(opts *bind.FilterOpts, timestamp []*big.Int, amount []*big.Int) (*FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__FreelancerStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowFreelancerStakeCompletedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__FreelancerStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowFreelancerStakeCompleted is a free log subscription operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowFreelancerStakeCompleted(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted, timestamp []*big.Int, amount []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__FreelancerStakeCompleted", timestampRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerStakeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowFreelancerStakeCompleted is a log parse operation binding the contract event 0x4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e.
//
// Solidity: event FreelanceEscrow__FreelancerStakeCompleted(uint256 indexed timestamp, uint256 indexed amount)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowFreelancerStakeCompleted(log types.Log) (*FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted, error) {
	event := new(FreelanceEscrowFreelanceEscrowFreelancerStakeCompleted)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__FreelancerStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator is returned from FilterFreelanceEscrowJobCompletedAndFreelancerPaid and is used to iterate over the raw logs and unpacked data for FreelanceEscrowJobCompletedAndFreelancerPaid events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator struct {
	Event *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid represents a FreelanceEscrowJobCompletedAndFreelancerPaid event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid struct {
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowJobCompletedAndFreelancerPaid is a free log retrieval operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowJobCompletedAndFreelancerPaid(opts *bind.FilterOpts, amount []*big.Int, timestamp []*big.Int) (*FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__JobCompletedAndFreelancerPaid", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaidIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__JobCompletedAndFreelancerPaid", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowJobCompletedAndFreelancerPaid is a free log subscription operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowJobCompletedAndFreelancerPaid(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid, amount []*big.Int, timestamp []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__JobCompletedAndFreelancerPaid", amountRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletedAndFreelancerPaid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowJobCompletedAndFreelancerPaid is a log parse operation binding the contract event 0x4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e981.
//
// Solidity: event FreelanceEscrow__JobCompletedAndFreelancerPaid(uint256 indexed amount, uint256 indexed timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowJobCompletedAndFreelancerPaid(log types.Log) (*FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid, error) {
	event := new(FreelanceEscrowFreelanceEscrowJobCompletedAndFreelancerPaid)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletedAndFreelancerPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator is returned from FilterFreelanceEscrowJobCompletionRejected and is used to iterate over the raw logs and unpacked data for FreelanceEscrowJobCompletionRejected events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowJobCompletionRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowJobCompletionRejected represents a FreelanceEscrowJobCompletionRejected event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowJobCompletionRejected struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowJobCompletionRejected is a free log retrieval operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowJobCompletionRejected(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__JobCompletionRejected")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowJobCompletionRejectedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__JobCompletionRejected", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowJobCompletionRejected is a free log subscription operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowJobCompletionRejected(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowJobCompletionRejected) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__JobCompletionRejected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletionRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowJobCompletionRejected is a log parse operation binding the contract event 0x52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd.
//
// Solidity: event FreelanceEscrow__JobCompletionRejected(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowJobCompletionRejected(log types.Log) (*FreelanceEscrowFreelanceEscrowJobCompletionRejected, error) {
	event := new(FreelanceEscrowFreelanceEscrowJobCompletionRejected)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__JobCompletionRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator is returned from FilterFreelanceEscrowPaymentDisputeRaised and is used to iterate over the raw logs and unpacked data for FreelanceEscrowPaymentDisputeRaised events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowPaymentDisputeRaised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowPaymentDisputeRaised represents a FreelanceEscrowPaymentDisputeRaised event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowPaymentDisputeRaised struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowPaymentDisputeRaised is a free log retrieval operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowPaymentDisputeRaised(opts *bind.FilterOpts) (*FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator, error) {

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__PaymentDisputeRaised")
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowPaymentDisputeRaisedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__PaymentDisputeRaised", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowPaymentDisputeRaised is a free log subscription operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowPaymentDisputeRaised(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowPaymentDisputeRaised) (event.Subscription, error) {

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__PaymentDisputeRaised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__PaymentDisputeRaised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowPaymentDisputeRaised is a log parse operation binding the contract event 0x7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e.
//
// Solidity: event FreelanceEscrow__PaymentDisputeRaised(uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowPaymentDisputeRaised(log types.Log) (*FreelanceEscrowFreelanceEscrowPaymentDisputeRaised, error) {
	event := new(FreelanceEscrowFreelanceEscrowPaymentDisputeRaised)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__PaymentDisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator is returned from FilterFreelanceEscrowRandomDisputeRaised and is used to iterate over the raw logs and unpacked data for FreelanceEscrowRandomDisputeRaised events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator struct {
	Event *FreelanceEscrowFreelanceEscrowRandomDisputeRaised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowRandomDisputeRaised represents a FreelanceEscrowRandomDisputeRaised event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRandomDisputeRaised struct {
	Raiser    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowRandomDisputeRaised is a free log retrieval operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowRandomDisputeRaised(opts *bind.FilterOpts, raiser []common.Address) (*FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator, error) {

	var raiserRule []interface{}
	for _, raiserItem := range raiser {
		raiserRule = append(raiserRule, raiserItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__RandomDisputeRaised", raiserRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowRandomDisputeRaisedIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__RandomDisputeRaised", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowRandomDisputeRaised is a free log subscription operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowRandomDisputeRaised(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowRandomDisputeRaised, raiser []common.Address) (event.Subscription, error) {

	var raiserRule []interface{}
	for _, raiserItem := range raiser {
		raiserRule = append(raiserRule, raiserItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__RandomDisputeRaised", raiserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RandomDisputeRaised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowRandomDisputeRaised is a log parse operation binding the contract event 0x022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c3991.
//
// Solidity: event FreelanceEscrow__RandomDisputeRaised(address indexed raiser, uint256 timestamp)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowRandomDisputeRaised(log types.Log) (*FreelanceEscrowFreelanceEscrowRandomDisputeRaised, error) {
	event := new(FreelanceEscrowFreelanceEscrowRandomDisputeRaised)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RandomDisputeRaised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator is returned from FilterFreelanceEscrowRevertedDealBreak and is used to iterate over the raw logs and unpacked data for FreelanceEscrowRevertedDealBreak events raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator struct {
	Event *FreelanceEscrowFreelanceEscrowRevertedDealBreak // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreelanceEscrowFreelanceEscrowRevertedDealBreak represents a FreelanceEscrowRevertedDealBreak event raised by the FreelanceEscrow contract.
type FreelanceEscrowFreelanceEscrowRevertedDealBreak struct {
	Reverter common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFreelanceEscrowRevertedDealBreak is a free log retrieval operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) FilterFreelanceEscrowRevertedDealBreak(opts *bind.FilterOpts, reverter []common.Address) (*FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator, error) {

	var reverterRule []interface{}
	for _, reverterItem := range reverter {
		reverterRule = append(reverterRule, reverterItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.FilterLogs(opts, "FreelanceEscrow__RevertedDealBreak", reverterRule)
	if err != nil {
		return nil, err
	}
	return &FreelanceEscrowFreelanceEscrowRevertedDealBreakIterator{contract: _FreelanceEscrow.contract, event: "FreelanceEscrow__RevertedDealBreak", logs: logs, sub: sub}, nil
}

// WatchFreelanceEscrowRevertedDealBreak is a free log subscription operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) WatchFreelanceEscrowRevertedDealBreak(opts *bind.WatchOpts, sink chan<- *FreelanceEscrowFreelanceEscrowRevertedDealBreak, reverter []common.Address) (event.Subscription, error) {

	var reverterRule []interface{}
	for _, reverterItem := range reverter {
		reverterRule = append(reverterRule, reverterItem)
	}

	logs, sub, err := _FreelanceEscrow.contract.WatchLogs(opts, "FreelanceEscrow__RevertedDealBreak", reverterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
				if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RevertedDealBreak", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFreelanceEscrowRevertedDealBreak is a log parse operation binding the contract event 0x9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a8.
//
// Solidity: event FreelanceEscrow__RevertedDealBreak(address indexed reverter)
func (_FreelanceEscrow *FreelanceEscrowFilterer) ParseFreelanceEscrowRevertedDealBreak(log types.Log) (*FreelanceEscrowFreelanceEscrowRevertedDealBreak, error) {
	event := new(FreelanceEscrowFreelanceEscrowRevertedDealBreak)
	if err := _FreelanceEscrow.contract.UnpackLog(event, "FreelanceEscrow__RevertedDealBreak", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
