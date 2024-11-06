package checkpointsubmitter

import (
	"context"
	"eip4788-proof-generator/internal/bindings"
	"eip4788-proof-generator/internal/checkpointproofs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	consensus "github.com/umbracle/go-eth-consensus"
	"github.com/umbracle/go-eth-consensus/http"
	"log"
	"math/big"
	"strings"
)

type CheckpointClient struct {
	evmClient        *ethclient.Client
	beaconClient     *http.BeaconEndpoint
	checkpointClient *bindings.Bindings
	transactor       *bind.TransactOpts
}

func New(ctx context.Context, evmRPCURL string, beaconURL string, checkpointTrackerAddress string, privateKey string) (*CheckpointClient, error) {
	evmClient, err := ethclient.DialContext(ctx, evmRPCURL)
	if err != nil {
		return nil, err
	}

	beaconClient := http.New(beaconURL).Beacon()

	checkpointClient, err := bindings.NewBindings(common.HexToAddress(checkpointTrackerAddress), evmClient)
	if err != nil {
		log.Fatalf("Failed to bind: %v", err)
	}

	pkey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return nil, err
	}

	chainID, err := evmClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(pkey, chainID)
	if err != nil {
		return nil, err
	}
	return &CheckpointClient{
		evmClient:        evmClient,
		beaconClient:     beaconClient,
		checkpointClient: checkpointClient,
		transactor:       transactor,
	}, nil
}

type stringBlockID struct {
	s string
}

func (s stringBlockID) BlockID() string {
	return s.s
}

func (c *CheckpointClient) SubmitProof(ctx context.Context, proof *checkpointproofs.ProofInputs) (*types.Receipt, error) {
	log := logrus.WithContext(ctx)

	// Fetch the block based on root
	id := stringBlockID{s: proof.Root.String()}
	header, err := c.beaconClient.GetBlockHeader(id)
	if err != nil {
		return nil, err
	}
	// We add 1 as the beacon root in the chain is parentbeaconroot
	slotToQueryForTimestamp := header.Header.Message.Slot + 1

	var beaconState consensus.BeaconStateDeneb
	if err = c.beaconClient.GetBeaconState(http.Slot(slotToQueryForTimestamp), &beaconState); err != nil {
		return nil, err
	}

	timestampToSubmit := beaconState.LatestExecutionPayloadHeader.Timestamp

	branch := make([][32]byte, len(proof.Branch))
	for i, v := range proof.Branch {
		branch[i] = v
	}

	// TODO: Dangerous casting
	inputTimestamp := big.NewInt(int64(timestampToSubmit))
	proofInput := bindings.IMerkleProofVerifierProofInputs{
		// TODO: Dangerous casting
		Index:  big.NewInt(int64(proof.Index)),
		Branch: branch,
		Value:  proof.Value,
	}
	log.Infof("Using timestamp %s for submission\n", inputTimestamp.String())
	transaction, err := c.checkpointClient.ProveCheckpointFinalized(c.transactor, inputTimestamp, proofInput)
	if err != nil {
		return nil, err
	}

	receipt, err := bind.WaitMined(ctx, c.evmClient, transaction)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (c *CheckpointClient) LatestFinalizedTimestamp(ctx context.Context) (*big.Int, error) {
	timestamp, err := c.checkpointClient.HighestFinalizedTimestamp(nil)
	if err != nil {
		return nil, err
	}
	return timestamp, nil
}
