// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {FinalizedCheckpointTracker} from "../src/FinalizedCheckpointTracker.sol";

contract DeployFinalizedCheckpointTracker is Script {
    address constant beaconRootsContract = 0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02;

    FinalizedCheckpointTracker public fct;

    function setUp() public {}

    // forge script script/DeployerFCT.s.sol:DeployFinalizedCheckpointTracker --fork-url $RPC --broadcast -vvvv
    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // TODO: manage the generalized index better
        fct = new FinalizedCheckpointTracker(beaconRootsContract, 744);

        console.log(address(fct));
        vm.stopBroadcast();
    }
}
