package checkpointproofs

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/beacon/merkle"
	"github.com/ethereum/go-ethereum/common"
	ssz "github.com/ferranbt/fastssz"
	consensus "github.com/umbracle/go-eth-consensus"
)

// ProofInputs represents the inputs required for proof verification.
type ProofInputs struct {
	Root   common.Hash
	Index  uint64
	Branch merkle.Values
	Value  merkle.Value
}

// GenerateCheckpointProof generates and verifies the proof.
func GenerateCheckpointProof(ctx context.Context, beaconBlockWithState *consensus.BeaconBlockWithStateDeneb) (*ProofInputs, error) {
	// Generate proof inputs
	proofInputs, err := generateCheckpointProofInputs(ctx, beaconBlockWithState)
	if err != nil {
		return nil, err
	}
	return proofInputs, nil
}

// VerifyCheckpointProof verifies the proof.
func VerifyCheckpointProof(ctx context.Context, proofInputs *ProofInputs) error {
	// Verify the proof
	if err := merkle.VerifyProof(proofInputs.Root, proofInputs.Index, proofInputs.Branch, proofInputs.Value); err != nil {
		return err
	}
	return nil
}

func generateCheckpointProofInputs(ctx context.Context, beaconBlock *consensus.BeaconBlockWithStateDeneb) (*ProofInputs, error) {
	leafIndexToProve, err := getLeafIndexForFinalizedCheckpoint(ctx, beaconBlock)
	if err != nil {
		return nil, err
	}

	beaconTree, err := beaconBlock.GetTree()
	if err != nil {
		return nil, err
	}

	// Generate the Leaf Value Input and store in leafValue
	finalizedEpochLeaf := make([]byte, 32)
	finalizedEpoch := beaconBlock.BeaconState.FinalizedCheckpoint.Epoch
	if finalizedEpoch == 0 {
		return nil, errors.New("cannot prove as no epochs have been finalized at this slot")
	}
	binary.LittleEndian.PutUint64(finalizedEpochLeaf, finalizedEpoch)
	var leafValue merkle.Value
	copy(leafValue[:], finalizedEpochLeaf)

	// Generate proof and transform into correct type
	proof, err := beaconTree.Prove(leafIndexToProve)
	if err != nil {
		return nil, err
	}
	var proofValues []merkle.Value
	for i := range proof.Hashes {
		var result merkle.Value
		copy(result[:], proof.Hashes[i])
		proofValues = append(proofValues, result)
	}

	// Generate the Root which we must check against
	trustedRoot, err := ssz.HashWithDefaultHasher(beaconBlock)
	if err != nil {
		return nil, err
	}

	return &ProofInputs{
		Root:   common.BytesToHash(trustedRoot[:]),
		Index:  uint64(leafIndexToProve),
		Branch: proofValues,
		Value:  leafValue,
	}, nil
}

func getLeafIndexForFinalizedCheckpoint(ctx context.Context, beaconBlock *consensus.BeaconBlockWithStateDeneb) (int, error) {
	beaconBlockTree, err := beaconBlock.GetTree()
	if err != nil {
		return 0, err
	}

	checkpointTree, err := beaconBlock.BeaconState.FinalizedCheckpoint.GetTree()
	if err != nil {
		return 0, err
	}

	// Brute force search to find the index of the tree that matches the checkpoint tree
	for i := 1; i < 1_000_000; i++ {
		get, err := beaconBlockTree.Get(i)
		if err != nil {
			continue
		}
		if hex.EncodeToString(get.Hash()) == hex.EncodeToString(checkpointTree.Hash()) {
			// The finalized slot is stored in the left leaf of the checkpoint tree
			leafIndex := 2 * i
			return leafIndex, nil
		}
	}
	return 0, errors.New("failed to find finalized checkpoint hash in tree")
}
