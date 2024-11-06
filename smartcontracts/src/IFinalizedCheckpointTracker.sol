// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./MerkleProofVerifier.sol";
import "./SSZHelper.sol";

/**
 * @title IFinalizedCheckpointTracker
 * @notice Interface for the FinalizedCheckpointTracker contract.
 * @dev Defines the functions and custom errors for the contract.
 */
interface IFinalizedCheckpointTracker {
    /**
     * @notice Proves that a checkpoint is finalized by verifying the Merkle proof.
     * @dev This function updates the highest finalized epoch and timestamp if the proof is valid
     *      and the epoch is greater than the current highest.
     *      It retrieves the beacon root from the Beacon Roots contract and verifies the proof against it.
     * @param timestamp The timestamp associated with the beacon root.
     * @param proof The Merkle proof inputs containing the value (serialized epoch) and proof data.
     *
     * Requirements:
     * - `timestamp` must be greater than or equal to `highestFinalizedTimestamp`.
     * - `epoch` extracted from `proof.value` must be greater than or equal to `highestFinalizedEpoch`.
     * - The call to `beaconRootsAddress` must succeed.
     * - The Merkle proof must be valid against the retrieved beacon root.
     *
     * Emits:
     * - `FinalizedCheckpointUpdated` event upon successful proof verification.
     *
     * Reverts:
     * - `AlreadyFinalized()` if `timestamp` is less than `highestFinalizedTimestamp`.
     * - `AlreadyFinalized()` if `epoch` is less than `highestFinalizedEpoch`.
     * - `BeaconRootsCallFailed()` if the call to `beaconRootsAddress` fails.
     * - `InvalidProof()` if the Merkle proof verification fails.
     */
    function proveCheckpointFinalized(uint256 timestamp, MerkleProofVerifier.ProofInputs calldata proof) external;

    /// @notice Emitted when a new highest finalized epoch is recorded.
    /// @param epoch The new highest finalized epoch.
    /// @param timestamp The timestamp associated with the epoch.
    event FinalizedCheckpointUpdated(uint256 indexed epoch, uint256 indexed timestamp);

    /// @notice Thrown when attempting to prove a checkpoint that is already finalized.
    error AlreadyFinalized();

    /// @notice Thrown when the call to the Beacon Roots contract fails.
    error BeaconRootsCallFailed();

    /// @notice Thrown when the Merkle proof verification fails.
    error InvalidProof();

    /// @notice Gets the address of the Beacon Roots contract.
    /// @return The address of the Beacon Roots contract.
    function beaconRootsAddress() external view returns (address);

    /// @notice Gets the highest finalized epoch that has been proven so far.
    /// @return The highest finalized epoch.
    function highestFinalizedEpoch() external view returns (uint256);

    /// @notice Gets the timestamp associated with the highest finalized epoch.
    /// @return The timestamp of the highest finalized epoch.
    function highestFinalizedTimestamp() external view returns (uint256);
}
