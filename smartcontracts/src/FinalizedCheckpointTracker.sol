// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "./MerkleProofVerifier.sol";
import "./SSZHelper.sol";
import "./IFinalizedCheckpointTracker.sol";

/**
 * @title FinalizedCheckpointTracker
 * @notice Tracks the highest finalized epoch and timestamp using proofs of finalized checkpoints.
 * @dev Inherits from MerkleProofVerifier and implements IFinalizedCheckpointTracker.
 */
contract FinalizedCheckpointTracker is MerkleProofVerifier, IFinalizedCheckpointTracker {
    using SSZHelper for bytes32;

    /// @inheritdoc IFinalizedCheckpointTracker
    address public immutable override beaconRootsAddress;

    /// @inheritdoc IFinalizedCheckpointTracker
    uint256 public override highestFinalizedEpoch;

    /// @inheritdoc IFinalizedCheckpointTracker
    uint256 public override highestFinalizedTimestamp;

    /**
     * @notice Initializes the contract with the Beacon Roots contract address.
     * @param beaconRootsAddress_ The address of the Beacon Roots contract.
     */
    constructor(address beaconRootsAddress_) {
        beaconRootsAddress = beaconRootsAddress_;
    }

    /**
     * @inheritdoc IFinalizedCheckpointTracker
     */
    function proveCheckpointFinalized(uint256 timestamp, ProofInputs calldata proof) external override {
        require(timestamp > highestFinalizedTimestamp, AlreadyFinalized());

        uint256 epoch = proof.value.fromLittleEndian();
        require(epoch > highestFinalizedEpoch, AlreadyFinalized());

        (bool success, bytes memory data) = beaconRootsAddress.call(bytes.concat(bytes32(timestamp)));
        require(success, BeaconRootsCallFailed());

        bytes32 beaconRoot = abi.decode(data, (bytes32));

        verifyProof(MerkleRoot.wrap(beaconRoot), proof);

        highestFinalizedEpoch = epoch;
        highestFinalizedTimestamp = timestamp;

        emit FinalizedCheckpointUpdated(epoch, timestamp);
    }
}
