// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console, Vm} from "forge-std/Test.sol";
import {IMerkleProofVerifier} from "../src/MerkleProofVerifier.sol";
import {FinalizedCheckpointTracker} from "../src/FinalizedCheckpointTracker.sol";

contract FinalizedCheckpointTrackerTest is Test, IMerkleProofVerifier {
    FinalizedCheckpointTracker fct;

    address constant beaconRootsContract = 0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02;

    function setUp() public {
        fct = new FinalizedCheckpointTracker(beaconRootsContract, 2);
    }

    function test_finalizedCheckpointValid_success() public {
        // Simple tree, proving Tree Index 2
        //                              c7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef
        //                     /                                                                \
        //    8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae     d8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f
        MerkleRoot root = MerkleRoot.wrap(0xc7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef);

        uint256 timestamp = 1730919182;
        bytes[] memory returnData = new bytes[](1);
        returnData[0] = bytes.concat(MerkleRoot.unwrap(root));

        // Mock the call to the beacon roots contract so that ir returns a fixed MerkleRoot
        vm.mockCalls(beaconRootsContract, bytes.concat(bytes32(timestamp)), returnData);

        // Form the inputs
        ProofInputs memory input;
        input.index = 2;
        bytes32[] memory branch = new bytes32[](1);
        branch[0] = 0xd8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f;
        input.branch = branch;
        input.value = 0x8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae;

        // Call Prove
        fct.proveCheckpointFinalized(timestamp, input);

        // Assert that state is updated correctly
        assertEq(
            fct.highestFinalizedEpoch(), 79101409427063649147342557613925153990311921204171468719376173058498201797252
        );
        assertEq(fct.highestFinalizedTimestamp(), timestamp);
    }

    function test_finalizedCheckpointInvalid_reverts() public {
        // Simple tree, proving Tree Index 2. Root returned from BeaconRoots is different
        //                              c7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef
        //                     /                                                                \
        //    8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae     d8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f

        uint256 timestamp = 1730919182;
        bytes[] memory returnData = new bytes[](1);
        returnData[0] = bytes.concat(bytes32(0xc038a39160ee3676e068e223eefb10f59c0f21c1e86d94ccb5755d95f5c05ba7));

        // Mock the call to the beacon roots contract so that ir returns a fixed MerkleRoot
        // The returned merkle root is purposely different to the one being proven against
        vm.mockCalls(beaconRootsContract, bytes.concat(bytes32(timestamp)), returnData);

        // Form the inputs
        ProofInputs memory input;
        input.index = 2;
        bytes32[] memory branch = new bytes32[](1);
        branch[0] = 0xd8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f;
        input.branch = branch;
        input.value = 0x8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae;

        // Call Prove expecting a revert
        vm.expectRevert(IMerkleProofVerifier.RootMismatch.selector);
        fct.proveCheckpointFinalized(timestamp, input);
    }
}
