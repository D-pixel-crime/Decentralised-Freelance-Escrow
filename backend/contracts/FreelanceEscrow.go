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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"jobId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"arbitrator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"confirmationPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"_noCancel\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_noDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyArbitrator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyClient\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_onlyFreelancer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addClientStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addfreelancerStake\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"breakDeal\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDealBreak\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finaliseUnilateralJob\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getClientStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmationPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEscrowState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFreelancerStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseDispute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rejectJobCompletion\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestPayment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolveDispute\",\"inputs\":[{\"name\":\"clientPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"freelancerPayment\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FreelanceEscrow__AggreementCreated\",\"inputs\":[{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"freelancer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__BothPartyStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__ClientStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealBroken\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DealCancelRequested\",\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__DisputeResolved\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__FreelancerStakeCompleted\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletedAndFreelancerPaid\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__JobCompletionRejected\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__PaymentDisputeRaised\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RandomDisputeRaised\",\"inputs\":[{\"name\":\"raiser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FreelanceEscrow__RevertedDealBreak\",\"inputs\":[{\"name\":\"reverter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ActiveConfirmationTimePeriod\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ClientAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__FreelancerAlreadyStaked\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidArbitrator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidClient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFreelancer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__InvalidFundsDistribution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NoCancelRequestedYet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__PaymentError\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__ProcessNotAllowed\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreelanceEscrow__RefundError\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__Busy\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"FreenlanceEscrow__DealAlreadyBroken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumFreelanceEscrow.EscrowState\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6101206040525f5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f6001555f600255348015610057575f5ffd5b506040516127f83803806127f8833981810160405281019061007991906102f7565b600161009761008c61023460201b60201c565b61025d60201b60201c565b5f01819055503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508461010081815250508373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660e08173ffffffffffffffffffffffffffffffffffffffff16815250505f60055f6101000a81548160ff0219169083600981111561019a5761019961036e565b5b02179055505f600560016101000a81548160ff021916908360098111156101c4576101c361036e565b5b0217905550806004819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fdbbe197ec6a15efb7d4ad9294dc52184ff29c79bef8778db27fbdae139c1d01f60405160405180910390a3505050505061039b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f5ffd5b5f819050919050565b61027c8161026a565b8114610286575f5ffd5b50565b5f8151905061029781610273565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102c68261029d565b9050919050565b6102d6816102bc565b81146102e0575f5ffd5b50565b5f815190506102f1816102cd565b92915050565b5f5f5f5f5f60a086880312156103105761030f610266565b5b5f61031d88828901610289565b955050602061032e888289016102e3565b945050604061033f888289016102e3565b9350506060610350888289016102e3565b925050608061036188828901610289565b9150509295509295909350565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60805160a05160c05160e051610100516123b66104425f395f50505f61133a01525f8181610535015281816106b501528181610ac201528181610d8b01528181610e1901528181610f55015281816110fe0152818161196801528181611aad0152611da301525f818161030501528181610a6b01528181610c9401528181610d2001528181610efe015281816110a70152818161186a01526118f801525f50506123b65ff3fe608060405260043610610113575f3560e01c806387d40b5a1161009f578063bdc84ac311610063578063bdc84ac31461026f578063d31cdf8914610297578063dd33869f146102ad578063df44be2c146102c3578063ee4b7ab8146102ed57610113565b806387d40b5a146101d1578063939ae273146101fb578063957448691461020557806397c9522e1461021b578063ac755a631461024557610113565b806363bdb94b116100e657806363bdb94b146101635780636977e176146101795780636daa2d441461018f57806375f84b40146101a55780637f25e289146101bb57610113565b806308e0771e1461011757806324dcdf641461012d5780633e216754146101375780634e07c9d11461014d575b5f5ffd5b348015610122575f5ffd5b5061012b610303565b005b61013561038a565b005b348015610142575f5ffd5b5061014b610533565b005b348015610158575f5ffd5b506101616105ba565b005b34801561016e575f5ffd5b506101776107e1565b005b348015610184575f5ffd5b5061018d6108f8565b005b34801561019a575f5ffd5b506101a3610ee4565b005b3480156101b0575f5ffd5b506101b9611095565b005b3480156101c6575f5ffd5b506101cf611338565b005b3480156101dc575f5ffd5b506101e56113bf565b6040516101f2919061210e565b60405180910390f35b6102036113c8565b005b348015610210575f5ffd5b50610219611571565b005b348015610226575f5ffd5b5061022f611635565b60405161023c919061210e565b60405180910390f35b348015610250575f5ffd5b5061025961163e565b604051610266919061210e565b60405180910390f35b34801561027a575f5ffd5b5061029560048036038101906102909190612155565b611647565b005b3480156102a2575f5ffd5b506102ab611be2565b005b3480156102b8575f5ffd5b506102c1611ca7565b005b3480156102ce575f5ffd5b506102d7611ecf565b6040516102e49190612206565b60405180910390f35b3480156102f8575f5ffd5b50610301611ee4565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610388576040517fe81d887300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b610392610533565b61039a61202b565b6103a2611571565b6103aa611be2565b5f60098111156103bd576103bc612193565b5b60055f9054906101000a900460ff1660098111156103de576103dd612193565b5b0361041257600260055f6101000a81548160ff0219169083600981111561040857610407612193565b5b02179055506104f4565b6001600981111561042657610425612193565b5b60055f9054906101000a900460ff16600981111561044757610446612193565b5b036104a857600360055f6101000a81548160ff0219169083600981111561047157610470612193565b5b0217905550427f84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da60405160405180910390a26104f3565b60055f9054906101000a900460ff166040517f16a523f80000000000000000000000000000000000000000000000000000000081526004016104ea9190612206565b60405180910390fd5b5b3460028190555034427f4eb16fadca2c78d42b73cba26d8a0f26724618ec841c29f92c59d8130c151c8e60405160405180910390a361053161204d565b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b8576040517f2240971400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6105c261202b565b6105ca610303565b6105d2611571565b6105da611be2565b600460098111156105ee576105ed612193565b5b60055f9054906101000a900460ff16600981111561060f5761060e612193565b5b1415801561065057506003600981111561062c5761062b612193565b5b60055f9054906101000a900460ff16600981111561064d5761064c612193565b5b14155b156106a05760055f9054906101000a900460ff166040517f361057570000000000000000000000000000000000000000000000000000000081526004016106979190612206565b60405180910390fd5b5f4790505f6001819055505f6002819055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16826040516106f79061224c565b5f6040518083038185875af1925050503d805f8114610731576040519150601f19603f3d011682016040523d82523d5f602084013e610736565b606091505b505090508061077e5781426040517f3f678334000000000000000000000000000000000000000000000000000000008152600401610775929190612260565b60405180910390fd5b42827f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a36005805f6101000a81548160ff021916908360098111156107d0576107cf612193565b5b021790555050506107df61204d565b565b6107e9610533565b6107f161202b565b6107f9611571565b610801611be2565b6003600981111561081557610814612193565b5b60055f9054906101000a900460ff16600981111561083657610835612193565b5b146108865760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161087d9190612206565b60405180910390fd5b600460055f6101000a81548160ff021916908360098111156108ab576108aa612193565b5b0217905550426003819055507ffbbb65879a5d311cf11666d7dbec54042f92fc07f89dec8b32cd7b1992dc4dc6426040516108e6919061210e565b60405180910390a16108f661204d565b565b61090061202b565b610908611571565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061099457506005600981111561097157610970612193565b5b60055f9054906101000a900460ff16600981111561099257610991612193565b5b145b156109e45760055f9054906101000a900460ff166040517f361057570000000000000000000000000000000000000000000000000000000081526004016109db9190612206565b60405180910390fd5b600760098111156109f8576109f7612193565b5b60055f9054906101000a900460ff166009811115610a1957610a18612193565b5b03610a695760055f9054906101000a900460ff166040517f9153e0e2000000000000000000000000000000000000000000000000000000008152600401610a609190612206565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610b1157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610b48576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff165f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c845760055f9054906101000a900460ff16600560016101000a81548160ff02191690836009811115610bce57610bcd612193565b5b0217905550335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600660055f6101000a81548160ff02191690836009811115610c3757610c36612193565b5b02179055503373ffffffffffffffffffffffffffffffffffffffff167ff0fd788aa2e0080a47bc875cc82d8c44adb2386bf6e439e99ac2c0c8b70501ce60405160405180910390a2610eda565b5f60015490505f6001819055505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051610cd69061224c565b5f6040518083038185875af1925050503d805f8114610d10576040519150601f19603f3d011682016040523d82523d5f602084013e610d15565b606091505b5050905080610d7d577f0000000000000000000000000000000000000000000000000000000000000000426040517fc3940718000000000000000000000000000000000000000000000000000000008152600401610d749291906122c6565b60405180910390fd5b60025491505f6002819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051610dcd9061224c565b5f6040518083038185875af1925050503d805f8114610e07576040519150601f19603f3d011682016040523d82523d5f602084013e610e0c565b606091505b50508091505080610e76577f0000000000000000000000000000000000000000000000000000000000000000426040517fc3940718000000000000000000000000000000000000000000000000000000008152600401610e6d9291906122c6565b60405180910390fd5b600760055f6101000a81548160ff02191690836009811115610e9b57610e9a612193565b5b02179055507fa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf2442604051610ecf919061210e565b60405180910390a150505b610ee261204d565b565b610eec61202b565b610ef4611571565b610efc611be2565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610fa457507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610fdb576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055f9054906101000a900460ff16600560016101000a81548160ff0219169083600981111561100e5761100d612193565b5b0217905550600860055f6101000a81548160ff0219169083600981111561103857611037612193565b5b02179055503373ffffffffffffffffffffffffffffffffffffffff167f022951c4cd0cc7b418ba2031cdba03ad982b223294cef2bfde4f2035fc7c399142604051611083919061210e565b60405180910390a261109361204d565b565b61109d61202b565b6110a5611571565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561114d57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15611184576040517f9206f47c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006600981111561119857611197612193565b5b60055f9054906101000a900460ff1660098111156111b9576111b8612193565b5b146112095760055f9054906101000a900460ff166040517feffe824c0000000000000000000000000000000000000000000000000000000081526004016112009190612206565b60405180910390fd5b6004600981111561121d5761121c612193565b5b600560019054906101000a900460ff16600981111561123f5761123e612193565b5b0361127357600360055f6101000a81548160ff0219169083600981111561126957611268612193565b5b02179055506112ac565b600560019054906101000a900460ff1660055f6101000a81548160ff021916908360098111156112a6576112a5612193565b5b02179055505b5f5f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f9669e06168d73c2ceacb7ff50b80de3e1d0db4a33c5da3a4c561aaac9669e9a860405160405180910390a261133661204d565b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146113bd576040517f5bffd58900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f600154905090565b6113d061202b565b6113d8610303565b6113e0611571565b6113e8611be2565b5f60098111156113fb576113fa612193565b5b60055f9054906101000a900460ff16600981111561141c5761141b612193565b5b0361145057600160055f6101000a81548160ff0219169083600981111561144657611445612193565b5b0217905550611532565b6002600981111561146457611463612193565b5b60055f9054906101000a900460ff16600981111561148557611484612193565b5b036114e657600360055f6101000a81548160ff021916908360098111156114af576114ae612193565b5b0217905550427f84eaa55e7294716fa360bacd01a05e5e37ddfa6761cd37f21af95e414f4360da60405160405180910390a2611531565b60055f9054906101000a900460ff166040517f92b538e00000000000000000000000000000000000000000000000000000000081526004016115289190612206565b60405180910390fd5b5b3460018190555034427f218bcd452fadafc4b7375c6aace19cfa7b70e56a522a386e72e3337f6a6512dc60405160405180910390a361156f61204d565b565b6008600981111561158557611584612193565b5b60055f9054906101000a900460ff1660098111156115a6576115a5612193565b5b14806115e357506009808111156115c0576115bf612193565b5b60055f9054906101000a900460ff1660098111156115e1576115e0612193565b5b145b156116335760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161162a9190612206565b60405180910390fd5b565b5f600454905090565b5f600254905090565b61164f61202b565b611657611338565b61165f611be2565b6008600981111561167357611672612193565b5b60055f9054906101000a900460ff16600981111561169457611693612193565b5b141580156116d457506009808111156116b0576116af612193565b5b60055f9054906101000a900460ff1660098111156116d1576116d0612193565b5b14155b156117245760055f9054906101000a900460ff166040517f3610575700000000000000000000000000000000000000000000000000000000815260040161171b9190612206565b60405180910390fd5b5f4790508183611734919061231a565b81101561176d576040517f4aa5857a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8314801561177b57505f82145b15611860576004600981111561179457611793612193565b5b600560019054906101000a900460ff1660098111156117b6576117b5612193565b5b036117ea57600360055f6101000a81548160ff021916908360098111156117e0576117df612193565b5b0217905550611823565b600560019054906101000a900460ff1660055f6101000a81548160ff0219169083600981111561181d5761181c612193565b5b02179055505b7f0939ea4648b30019633687bbe9bcd0c90c5a678c6801f6a55610eeaa0399477e42604051611852919061210e565b60405180910390a150611bd6565b5f5f8414611a9c577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16846040516118ac9061224c565b5f6040518083038185875af1925050503d805f81146118e6576040519150601f19603f3d011682016040523d82523d5f602084013e6118eb565b606091505b50508091505080611957577f00000000000000000000000000000000000000000000000000000000000000006001546040517fc394071800000000000000000000000000000000000000000000000000000000815260040161194e9291906122c6565b60405180910390fd5b5f8483611964919061234d565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16816040516119aa9061224c565b5f6040518083038185875af1925050503d805f81146119e4576040519150601f19603f3d011682016040523d82523d5f602084013e6119e9565b606091505b50508092505081611a3557806001546040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611a2c929190612260565b60405180910390fd5b600760055f6101000a81548160ff02191690836009811115611a5a57611a59612193565b5b02179055507fa5e2e04f3b86e38e6841f1530d9f4f01ecbed17e03c2303d0625ae615fecbf2442604051611a8e919061210e565b60405180910390a150611bd3565b5f8483611aa9919061234d565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681604051611aef9061224c565b5f6040518083038185875af1925050503d805f8114611b29576040519150601f19603f3d011682016040523d82523d5f602084013e611b2e565b606091505b50508092505081611b7a57806001546040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611b71929190612260565b60405180910390fd5b6005805f6101000a81548160ff02191690836009811115611b9e57611b9d612193565b5b021790555042817f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a3505b50505b611bde61204d565b5050565b60066009811115611bf657611bf5612193565b5b60055f9054906101000a900460ff166009811115611c1757611c16612193565b5b1480611c55575060076009811115611c3257611c31612193565b5b60055f9054906101000a900460ff166009811115611c5357611c52612193565b5b145b15611ca55760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611c9c9190612206565b60405180910390fd5b565b611caf61202b565b611cb7611571565b611cbf611be2565b60046009811115611cd357611cd2612193565b5b60055f9054906101000a900460ff166009811115611cf457611cf3612193565b5b14611d445760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611d3b9190612206565b60405180910390fd5b600454600354611d54919061231a565b421015611d9c57600354426040517fc3abe3e2000000000000000000000000000000000000000000000000000000008152600401611d93929190612260565b60405180910390fd5b5f4790505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1682604051611de59061224c565b5f6040518083038185875af1925050503d805f8114611e1f576040519150601f19603f3d011682016040523d82523d5f602084013e611e24565b606091505b5050905080611e6c5781426040517f3f678334000000000000000000000000000000000000000000000000000000008152600401611e63929190612260565b60405180910390fd5b42827f4522aa30f768574398e49d722c5bc0969cb2e6c5fdffaef4dbf5cb7d6017e98160405160405180910390a36005805f6101000a81548160ff02191690836009811115611ebe57611ebd612193565b5b02179055505050611ecd61204d565b565b5f60055f9054906101000a900460ff16905090565b611eec61202b565b611ef4610303565b611efc611571565b611f04611be2565b60046009811115611f1857611f17612193565b5b60055f9054906101000a900460ff166009811115611f3957611f38612193565b5b14611f895760055f9054906101000a900460ff166040517f36105757000000000000000000000000000000000000000000000000000000008152600401611f809190612206565b60405180910390fd5b600960055f6101000a81548160ff02191690836009811115611fae57611fad612193565b5b02179055507f52308313ade0f30bcaf4b85b70e27441ea4b21c0dfefae04d1135fc92a87eadd42604051611fe2919061210e565b60405180910390a17f7c4bbc177b774c93b3bff2cac72521a5be4b3bc65faa8a60bd05afe89571324e42604051612019919061210e565b60405180910390a161202961204d565b565b612033612067565b60026120456120406120a8565b6120d1565b5f0181905550565b600161205f61205a6120a8565b6120d1565b5f0181905550565b61206f6120da565b156120a6576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f60026120ed6120e86120a8565b6120d1565b5f015414905090565b5f819050919050565b612108816120f6565b82525050565b5f6020820190506121215f8301846120ff565b92915050565b5f5ffd5b612134816120f6565b811461213e575f5ffd5b50565b5f8135905061214f8161212b565b92915050565b5f5f6040838503121561216b5761216a612127565b5b5f61217885828601612141565b925050602061218985828601612141565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600a81106121d1576121d0612193565b5b50565b5f8190506121e1826121c0565b919050565b5f6121f0826121d4565b9050919050565b612200816121e6565b82525050565b5f6020820190506122195f8301846121f7565b92915050565b5f81905092915050565b50565b5f6122375f8361221f565b915061224282612229565b5f82019050919050565b5f6122568261222c565b9150819050919050565b5f6040820190506122735f8301856120ff565b61228060208301846120ff565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6122b082612287565b9050919050565b6122c0816122a6565b82525050565b5f6040820190506122d95f8301856122b7565b6122e660208301846120ff565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612324826120f6565b915061232f836120f6565b9250828201905080821115612347576123466122ed565b5b92915050565b5f612357826120f6565b9150612362836120f6565b925082820390508181111561237a576123796122ed565b5b9291505056fea2646970667358221220c7f553d702565398335b1ec559bf2ff919c24da2ccbcd4b0363ce87bff261b1a64736f6c634300081e0033",
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
// Solidity: function acceptJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowTransactor) AcceptJobCompletion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreelanceEscrow.contract.Transact(opts, "acceptJobCompletion")
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() returns()
func (_FreelanceEscrow *FreelanceEscrowSession) AcceptJobCompletion() (*types.Transaction, error) {
	return _FreelanceEscrow.Contract.AcceptJobCompletion(&_FreelanceEscrow.TransactOpts)
}

// AcceptJobCompletion is a paid mutator transaction binding the contract method 0x4e07c9d1.
//
// Solidity: function acceptJobCompletion() returns()
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
