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
        BUSY,
        CLIENT_STAKED,
        FREELANCER_STAKED,
        ALL_STAKED_AND_PENDING,
        PENDING_CLIENT_CONFIRMATION,
        JOB_COMPLETED,
        CANCEL_REQUESTED,
        DEAL_BROKEN
    }
    address private immutable i_owner;
    address private immutable i_client;
    address private immutable i_freelancer;
    uint256 private immutable i_jobId;
    uint256 public s_clientStake = 0;
    uint256 public s_freelancerStake = 0;
    EscrowState private currState;
    EscrowState private prevState;
    address private s_dealBreakInitiator = address(0);
    uint256 private s_unilateralCompletionTimestamp;

    error FreelanceEscrow__NotParticipant();
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
    event FreelanceEscrow__JobCompletedRejected(uint256 timestamp);
    event FreelanceEscrow__DealCancelRequested(address indexed initiator);
    event FreelanceEscrow__DealBroken(uint256 timestamp);
    event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
        uint256 timestamp
    );
    event FreelanceEscrow__AggreementCreated(
        address indexed client,
        address indexed freelancer
    );
    event Freelance__RevertedDealBreak(address indexed reverter);

    modifier onlyClient() {
        _onlyClient();
        _;
    }

    constructor(uint256 jobId, address client, address freelancer) {
        i_owner = msg.sender;
        i_jobId = jobId;
        i_client = client;
        i_freelancer = freelancer;
        currState = EscrowState.AGREED;
        prevState = EscrowState.AGREED;
        emit FreelanceEscrow__AggreementCreated(client, freelancer);
    }

    function addClientStake() public payable nonReentrant onlyClient {
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

    function addfreelancerStake() public payable nonReentrant {
        if (msg.sender != i_freelancer) {
            revert FreelanceEscrow__InvalidFreelancer();
        }

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

    function freelancerCompletedAndRequestingFundsUnilaterally()
        public
        nonReentrant
    {
        if (msg.sender != i_freelancer) {
            revert FreelanceEscrow__InvalidFreelancer();
        }
        if (currState != EscrowState.ALL_STAKED_AND_PENDING) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        /** @dev Start a timer of 5 days for Client Confirmations */

        currState = EscrowState.PENDING_CLIENT_CONFIRMATION;
        emit FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
            block.timestamp
        );
    }

    function finaliseUnilateralJob() public nonReentrant {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
        if (block.timestamp < s_unilateralCompletionTimestamp + 5 days) {
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

    function rejectJobCompletion() public nonReentrant onlyClient {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        currState = EscrowState.ALL_STAKED_AND_PENDING;
        emit FreelanceEscrow__JobCompletedRejected(block.timestamp);
    }

    function acceptJobCompletion() public nonReentrant onlyClient {
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
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

    function jobCompletionAgreedBilaterally() public nonReentrant {
        if (currState != EscrowState.ALL_STAKED_AND_PENDING) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }
        if (msg.sender != i_client) {
            revert FreelanceEscrow__InvalidClient();
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

    function breakDeal() public nonReentrant {
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

    function cancelDealBreak() public nonReentrant {
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
        emit Freelance__RevertedDealBreak(msg.sender);
    }

    /** @dev Getter Functions */
    function getEscrowState() public view returns (EscrowState) {
        return currState;
    }

    /** @dev Internal Functions */
    function _onlyClient() public view {
        if (msg.sender != i_client) {
            revert FreelanceEscrow__InvalidClient();
        }
    }
}
