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
        CLIENT_STAKED,
        FREELANCER_STAKED,
        ALL_STAKED_AND_PENDING,
        PENDING_CLIENT_CONFIRMATION
    }
    address private immutable i_client;
    address private immutable i_freelancer;
    uint256 private immutable i_jobId;
    uint256 public s_clientStake = 0;
    uint256 public s_freelancerStake = 0;
    EscrowState private currState;

    error FreelanceEscrow_invalidClient();
    error FreelanceEscrow_invalidFreelancer();
    error FreelanceEscrow_clientAlreadyStaked(EscrowState);
    error FreelanceEscrow_freelancerAlreadyStaked(EscrowState);

    event FreelanceEscrow_clientStakeCompleted(
        uint256 indexed timestamp,
        uint256 indexed amount
    );
    event FreelanceEscrow_freelancerStakeCompleted(
        uint256 indexed timestamp,
        uint256 indexed amount
    );
    event FreelanceEscrow_bothPartyStakeCompleted(uint256 indexed timestamp);

    constructor(uint256 jobId, address client, address freelancer) {
        i_jobId = jobId;
        i_client = client;
        i_freelancer = freelancer;
        currState = EscrowState.AGREED;
    }

    function addClientStake() public payable {
        if (msg.sender != i_client) {
            revert FreelanceEscrow_invalidClient();
        }
        if (currState == EscrowState.AGREED) {
            currState = EscrowState.CLIENT_STAKED;
        } else if (currState == EscrowState.FREELANCER_STAKED) {
            currState = EscrowState.ALL_STAKED_AND_PENDING;
            emit FreelanceEscrow_bothPartyStakeCompleted(block.timestamp);
        } else {
            revert FreelanceEscrow_clientAlreadyStaked(currState);
        }
        s_clientStake = msg.value;
        emit FreelanceEscrow_clientStakeCompleted(block.timestamp, msg.value);
    }

    function addfreelancerStake() public payable {
        if (msg.sender != i_freelancer) {
            revert FreelanceEscrow_invalidFreelancer();
        }
        if (currState == EscrowState.AGREED) {
            currState = EscrowState.FREELANCER_STAKED;
        } else if (currState == EscrowState.CLIENT_STAKED) {
            currState = EscrowState.ALL_STAKED_AND_PENDING;
            emit FreelanceEscrow_bothPartyStakeCompleted(block.timestamp);
        } else {
            revert FreelanceEscrow_freelancerAlreadyStaked(currState);
        }
        s_freelancerStake = msg.value;
        emit FreelanceEscrow_freelancerStakeCompleted(
            block.timestamp,
            msg.value
        );
    }

    /** @dev Getter Functions */
    function getEscrowState() public view returns (EscrowState) {
        return currState;
    }
}
