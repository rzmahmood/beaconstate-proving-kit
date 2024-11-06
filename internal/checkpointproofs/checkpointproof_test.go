package checkpointproofs

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	consensus "github.com/umbracle/go-eth-consensus"
	"github.com/umbracle/go-eth-consensus/http"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateCheckpointProof_ValidBlockValidState_VerifiesCorrectly(t *testing.T) {
	// -- Arrange
	beaconBlockPath := filepath.Join("testdata", "slot282_beaconblock_valid.json")
	beaconStatePath := filepath.Join("testdata", "slot282_beaconstate_valid.json")


	beaconBlock := getBeaconBlockFromPath(t, beaconBlockPath)
	require.Equal(t, uint64(282), beaconBlock.Slot)
	beaconState := getBeaconStateFromPath(t, beaconStatePath)
	require.Equal(t, uint64(282), beaconState.Slot)

	beaconBlockWithState := &consensus.BeaconBlockWithStateDeneb{
		Slot:          beaconBlock.Slot,
		ProposerIndex: beaconBlock.ProposerIndex,
		ParentRoot:    beaconBlock.ParentRoot,
		BeaconState:   &beaconState,
		Body:          beaconBlock.Body,
	}

	ctx := context.Background()
	// -- Act
	proofInputs, err := GenerateCheckpointProof(ctx, beaconBlockWithState)
	require.NoError(t, err)
	err = VerifyCheckpointProof(ctx, proofInputs)

	// -- Assert
	tree, err := beaconBlock.GetTree()
	require.NoError(t, err)
	// The block root hash should equal the one calculated
	require.True(t, bytes.Equal(proofInputs.Root.Bytes(), tree.Hash()))
}

func TestGenerateCheckpointProof_ValidBlockInvalidState_FailsVerification(t *testing.T) {
	// -- Arrange
	beaconBlockPath := filepath.Join("testdata", "slot282_beaconblock_valid.json")
	// blob gas used was change 0 -> 1 to simulate invalid body to header
	beaconStatePath := filepath.Join("testdata", "slot282_beaconstate_invalid.json")


	beaconBlock := getBeaconBlockFromPath(t, beaconBlockPath)
	require.Equal(t, uint64(282), beaconBlock.Slot)
	beaconState := getBeaconStateFromPath(t, beaconStatePath)
	require.Equal(t, uint64(282), beaconState.Slot)

	beaconBlockWithState := &consensus.BeaconBlockWithStateDeneb{
		Slot:          beaconBlock.Slot,
		ProposerIndex: beaconBlock.ProposerIndex,
		ParentRoot:    beaconBlock.ParentRoot,
		BeaconState:   &beaconState,
		Body:          beaconBlock.Body,
	}

	ctx := context.Background()
	// -- Act
	proofInputs, err := GenerateCheckpointProof(ctx, beaconBlockWithState)
	require.NoError(t, err)
	err = VerifyCheckpointProof(ctx, proofInputs)

	// -- Assert
	tree, err := beaconBlock.GetTree()
	require.NoError(t, err)
	// The block root hash should NOT equal the one calculated
	require.False(t, bytes.Equal(proofInputs.Root.Bytes(), tree.Hash()))
}

func getBeaconBlockFromPath(t *testing.T, path string) consensus.BeaconBlockDeneb {
	// Read the JSON file
	data, err := os.ReadFile(path)
	assert.NoError(t, err)

	// Now unmarshal the Message field into beaconBlock
	var beaconBlock consensus.BeaconBlockDeneb
	out := &http.Block{
		Message: &beaconBlock,
	}
	err = http.Unmarshal(data, &out, false)
	assert.NoError(t, err)

	return beaconBlock
}

func getBeaconStateFromPath(t *testing.T, path string) consensus.BeaconStateDeneb {
	// Read the JSON file
	data, err := os.ReadFile(path)
	assert.NoError(t, err)

	// Now unmarshal the Message field into beaconBlock
	var beaconState consensus.BeaconStateDeneb
	err = http.Unmarshal(data, &beaconState, true)
	assert.NoError(t, err)

	return beaconState
}