// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Freelance-Escrow
 * @author D-Pixel-Crime
 * @notice Decentralised solution for trustless freelancing
 */
contract FreelanceEscrow {
    enum EscrowState {
        AGREED,
        BUSY,
        CLIENT_STAKED,
        FREELANCER_STAKED,
        ALL_STAKED_AND_PENDING,
        PENDING_CLIENT_CONFIRMATION,
        JOB_COMPLETED,
        DEAL_BROKEN
    }
    address private immutable i_owner;
    address private immutable i_client;
    address private immutable i_freelancer;
    uint256 private immutable i_jobId;
    uint256 public s_clientStake = 0;
    uint256 public s_freelancerStake = 0;
    EscrowState private currState;

    error FreelanceEscrow__InvalidOwner();
    error FreelanceEscrow__InvalidClient();
    error FreelanceEscrow__InvalidFreelancer();
    error FreelanceEscrow__ClientAlreadyStaked(EscrowState);
    error FreelanceEscrow__FreelancerAlreadyStaked(EscrowState);
    error FreenlanceEscrow__DealAlreadyBroken(EscrowState);
    error FreenlanceEscrow__Busy(EscrowState);
    error FreelanceEscrow__RefundError(address, uint256);
    error FreelanceEscrow__ProcessNotAllowed(EscrowState);
    error FreelanceEscrow__PaymentError(uint256, uint256);

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
    event FreelanceEscrow__DealBroken(uint256 timestamp);
    event FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
        uint256 timestamp
    );

    modifier onlyOwner() {
        _onlyOwner();
        _;
    }

    modifier checkEscrowBusy() {
        _checkEscrowBusy();
        _;
    }

    constructor(uint256 jobId, address client, address freelancer) {
        i_owner = msg.sender;
        i_jobId = jobId;
        i_client = client;
        i_freelancer = freelancer;
        currState = EscrowState.AGREED;
    }

    function addClientStake(
        address client
    ) public payable onlyOwner checkEscrowBusy {
        if (client != i_client) {
            revert FreelanceEscrow__InvalidClient();
        }

        currState = EscrowState.BUSY;

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

    function addfreelancerStake(
        address freelancer
    ) public payable onlyOwner checkEscrowBusy {
        if (freelancer != i_freelancer) {
            revert FreelanceEscrow__InvalidFreelancer();
        }

        currState = EscrowState.BUSY;

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

    function breakDeal() public onlyOwner checkEscrowBusy {
        if (currState == EscrowState.DEAL_BROKEN) {
            revert FreenlanceEscrow__DealAlreadyBroken(currState);
        }

        currState = EscrowState.BUSY;

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

    function freelancerCompleted(
        address freelancer
    ) public onlyOwner checkEscrowBusy {
        if (freelancer != i_freelancer) {
            revert FreelanceEscrow__InvalidFreelancer();
        }
        if (currState != EscrowState.ALL_STAKED_AND_PENDING) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        currState = EscrowState.BUSY;

        /** @dev Start a timer of 5 days for Client Confirmations */

        currState = EscrowState.PENDING_CLIENT_CONFIRMATION;
        emit FreelanceEscrow__FreelancerCompletedAndClientConfirmationPending(
            block.timestamp
        );
    }

    function rejectJobCompletion(
        address client
    ) public onlyOwner checkEscrowBusy {
        if (client != i_client) {
            revert FreelanceEscrow__InvalidClient();
        }
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        currState = EscrowState.ALL_STAKED_AND_PENDING;
        emit FreelanceEscrow__JobCompletedRejected(block.timestamp);
    }

    function acceptJobCompletion(
        address client
    ) public onlyOwner checkEscrowBusy {
        if (client != i_client) {
            revert FreelanceEscrow__InvalidClient();
        }
        if (currState != EscrowState.PENDING_CLIENT_CONFIRMATION) {
            revert FreelanceEscrow__ProcessNotAllowed(currState);
        }

        currState = EscrowState.BUSY;
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

    /** @dev Getter Functions */
    function getEscrowState() public view returns (EscrowState) {
        return currState;
    }

    function getOwner() public view returns (address) {
        return i_owner;
    }

    /** @dev Internal Functions */
    function _checkEscrowBusy() internal view {
        if (currState == EscrowState.BUSY) {
            revert FreenlanceEscrow__Busy(currState);
        }
    }

    function _onlyOwner() internal view {
        if (msg.sender != i_owner) {
            revert FreelanceEscrow__InvalidOwner();
        }
    }
}
