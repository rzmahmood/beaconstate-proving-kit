// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import {Test, console, Vm} from "forge-std/Test.sol";
import {MerkleProofVerifier, IMerkleProofVerifier} from "../src/MerkleProofVerifier.sol";

contract MerkleProofVerifierTest is Test, IMerkleProofVerifier {
    MerkleProofVerifier mpv;

    function setUp() public {
        mpv = new MerkleProofVerifier();
    }

    function test_validProofLeftTree_success() public view {
        // Simple tree, proving Tree Index 2
        //                              c7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef
        //                     /                                                                \
        //    8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae     d8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f
        MerkleRoot root = MerkleRoot.wrap(0xc7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef);

        ProofInputs memory input;
        input.index = 2;
        bytes32[] memory branch = new bytes32[](1);
        branch[0] = 0xd8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f;
        input.branch = branch;
        input.value = 0x8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae;

        mpv.verifyProof(root, input);
    }

    function test_validProofRightTree_success() public view {
        // Simple tree, proving Tree Index 3
        //                              c7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef
        //                     /                                                                \
        //    8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae     d8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f
        MerkleRoot root = MerkleRoot.wrap(0xc7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef);

        ProofInputs memory input;
        input.index = 3;
        bytes32[] memory branch = new bytes32[](1);
        branch[0] = 0x8442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae;
        input.branch = branch;
        input.value = 0xd8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f;

        mpv.verifyProof(root, input);
    }

    function test_invalidProof_reverts() public {
        // Simple tree, proving Tree Index 3 - The Proof is incorrect as it has been twiddled
        //                              c7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef
        //                     /                                                                \
        //    7442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae     d8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f
        MerkleRoot root = MerkleRoot.wrap(0xc7ceaeb9fe27c2542d2ea23342de897d3379c81807591fb8869499823d8de1ef);

        ProofInputs memory input;
        input.index = 3;
        bytes32[] memory branch = new bytes32[](1);
        branch[0] = 0x7442eef091c609959286e4f8f558dfe5bcbce55ceab1b15bdd02c4bfa7cfe1ae;
        input.branch = branch;
        input.value = 0xd8f6d8e58fb41aa9da0e5355aa759571e85086b3bd425046abce97b518a7282f;

        vm.expectRevert(IMerkleProofVerifier.RootMismatch.selector);
        mpv.verifyProof(root, input);
    }
}
