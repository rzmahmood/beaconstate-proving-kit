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

    address public immutable beaconRootsAddress;

    uint256 public immutable leafIndex;

    uint256 public highestFinalizedEpoch;

    uint256 public highestFinalizedTimestamp;

    /**
     * @notice Initializes the contract with the Beacon Roots contract address.
     * @param beaconRootsAddress_ The address of the Beacon Roots contract.
     * @param leafIndex_ The expected leaf index for the value we want to prove.
     * We cannot trust the user to give us a valid value
     */
    constructor(address beaconRootsAddress_, uint256 leafIndex_) {
        beaconRootsAddress = beaconRootsAddress_;
        leafIndex = leafIndex_;
    }

    /**
     * @inheritdoc IFinalizedCheckpointTracker
     */
    function proveCheckpointFinalized(uint256 timestamp, ProofInputs calldata proof) external override {
        require(timestamp > highestFinalizedTimestamp, AlreadyFinalized());
        require(proof.index == leafIndex, InvalidProof());

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
