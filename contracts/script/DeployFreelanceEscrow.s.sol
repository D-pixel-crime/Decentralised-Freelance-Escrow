// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Script} from "forge-std/Script.sol";
import {FreelanceEscrow} from "../src/FreelanceEscrow.sol";

contract DeployFreelanceEscrow is Script {
    function run() external {
        uint256 deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        vm.startBroadcast(deployerPrivateKey);

        // Account 1
        address client = address(0x70997970C51812dc3A010C7d01b50e0d17dc79C8);
        // Account 2
        address freelancer = address(0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC);
        // Account 0 (Deployer)
        address arbitrator = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);

        FreelanceEscrow escrow = new FreelanceEscrow(
            1, // jobId
            client,
            freelancer,
            arbitrator,
            5 days // confirmationPeriod
        );

        vm.stopBroadcast();
    }
}
