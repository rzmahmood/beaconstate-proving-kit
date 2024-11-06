// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title IMerkleProofVerifier
 * @notice Interface for the MerkleProofVerifier contract.
 * @dev Provides the function definition and types for Merkle proof verification.
 */
interface IMerkleProofVerifier {
    /// @notice Represents a Merkle root hash.
    type MerkleRoot is bytes32;

    /**
     * @notice Structure containing inputs required for Merkle proof verification.
     * @dev The `index` is the generalized index of the leaf in the tree,
     *      `branch` is an array of sibling hashes constituting the Merkle proof,
     *      and `value` is the leaf value being verified.
     */
    struct ProofInputs {
        /// @notice Generalized index of the leaf node in the Merkle tree.
        uint256 index;
        /// @notice Array of sibling hashes for the Merkle proof.
        bytes32[] branch;
        /// @notice The value of the leaf node being proved.
        bytes32 value;
    }

    /// @notice Error indicating that the proof branch has extra items.
    error BranchHasExtraItems();

    /// @notice Error indicating that the proof branch is missing items.
    error BranchIsMissingItems();

    /// @notice Error indicating that the computed root does not match the expected root.
    error RootMismatch();
}
