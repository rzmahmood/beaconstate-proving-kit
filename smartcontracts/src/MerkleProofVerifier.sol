// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./IMerkleProofVerifier.sol";

/**
 * @title MerkleProofVerifier
 * @notice Contract for verifying Merkle proofs.
 * @dev Implements the `verifyProof` function to verify a Merkle proof for a given leaf value.
 *      Adapted from https://github.com/ethereum/go-ethereum/blob/master/beacon/merkle/merkle.go#L44
 */
contract MerkleProofVerifier is IMerkleProofVerifier {
    /**
     * @notice Verifies a Merkle proof for a given leaf value and Merkle root.
     * @dev This function traverses the proof branch and computes the hash up to the root,
     *      verifying that the provided value is part of the Merkle tree represented by the root.
     * @param root The Merkle root against which the proof is verified.
     * @param proof The proof inputs containing the index, branch (sibling hashes), and value (leaf value).
     *
     * Reverts:
     * - `BranchHasExtraItems` if the proof branch has extra items.
     * - `BranchIsMissingItems` if the proof branch is missing items.
     * - `RootMismatch` if the computed root does not match the provided root.
     */
    function verifyProof(MerkleRoot root, ProofInputs memory proof) public pure {
        for (uint256 i = 0; i < proof.branch.length; i++) {
            bytes32 sibling = proof.branch[i];

            if (proof.index & 1 == 0) {
                proof.value = sha256(abi.encodePacked(proof.value, sibling));
            } else {
                proof.value = sha256(abi.encodePacked(sibling, proof.value));
            }

            proof.index = proof.index >> 1;

            if (proof.index == 0) {
                revert BranchHasExtraItems();
            }
        }

        if (proof.index != 1) {
            revert BranchIsMissingItems();
        }
        if (proof.value != MerkleRoot.unwrap(root)) {
            revert RootMismatch();
        }
    }
}
