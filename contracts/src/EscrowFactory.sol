// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {FreelanceEscrow} from "./FreelanceEscrow.sol";

contract EscrowFactory {
    event EscrowCreated(
        uint256 indexed jobId,
        address escrowAddress,
        address indexed client,
        address indexed freelancer
    );

    function createEscrow(
        uint256 _jobId,
        address _client,
        address _freelancer,
        address _arbitrator,
        uint256 _confirmationPeriod
    ) public returns (address) {
        FreelanceEscrow newEscrow = new FreelanceEscrow(
            _jobId,
            _client,
            _freelancer,
            _arbitrator,
            _confirmationPeriod
        );

        emit EscrowCreated(_jobId, address(newEscrow), _client, _freelancer);

        return address(newEscrow);
    }
}
