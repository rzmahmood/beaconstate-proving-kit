package consensus

import (
	"context"

	consensus "github.com/umbracle/go-eth-consensus"
	"github.com/umbracle/go-eth-consensus/http"
)

type BeaconClient interface {
	GetBlockWithState(ctx context.Context, slot uint64) (*consensus.BeaconBlockWithStateDeneb, error)
}

// BeaconClientRPC wraps the beacon node HTTP client.
type BeaconClientRPC struct {
	endpoint *http.BeaconEndpoint
}

// NewBeaconClientRPC creates a new BeaconClient.
func NewBeaconClientRPC(url string) *BeaconClientRPC {
	return &BeaconClientRPC{
		endpoint: http.New(url).Beacon(),
	}
}

// GetBlockWithState retrieves a beacon block along with its state.
func (bc *BeaconClientRPC) GetBlockWithState(ctx context.Context, slot uint64) (*consensus.BeaconBlockWithStateDeneb, error) {
	testBlockID := http.Slot(slot)

	var beaconBlock consensus.BeaconBlockDeneb
	_, err := bc.endpoint.GetBlock(testBlockID, &beaconBlock)
	if err != nil {
		return nil, err
	}

	var beaconState consensus.BeaconStateDeneb
	err = bc.endpoint.GetBeaconState(testBlockID, &beaconState)
	if err != nil {
		return nil, err
	}

	beaconBlockWithState := &consensus.BeaconBlockWithStateDeneb{
		Slot:          beaconBlock.Slot,
		ProposerIndex: beaconBlock.ProposerIndex,
		ParentRoot:    beaconBlock.ParentRoot,
		BeaconState:   &beaconState,
		Body:          beaconBlock.Body,
	}

	return beaconBlockWithState, nil
}
