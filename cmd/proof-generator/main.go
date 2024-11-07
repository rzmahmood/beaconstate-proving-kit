package main

import (
	"context"
	"eip4788-proof-generator/internal/checkpointproofs"
	"eip4788-proof-generator/internal/checkpointsubmitter"
	"eip4788-proof-generator/pkg/consensus"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	ctx := context.Background()
	log := logrus.WithContext(ctx)

	// Parse and validate command-line flags
	beaconURL := flag.String("beacon-url", "http://127.0.0.1:4100", "Beacon node endpoint URL")
	evmURL := flag.String("evm-url", "http://127.0.0.1:8000", "EVM RPC URL to submit")
	checkpointTrackerAddress := flag.String("checkpoint-tracker-address", "", "Checkpoint tracker contract address (required)")
	slot := flag.Uint64("slot", 230, "Beacon block slot to prove finalized")
	privateKey := flag.String("pkey", "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "private key to submit proof with")
	submission := flag.Bool("submit", true, "Enable submission to chain")

	flag.Parse()

	if *checkpointTrackerAddress == "" {
		fmt.Println("Error: --checkpoint-tracker-address is required")
		flag.Usage()
		os.Exit(1)
	}

	if !common.IsHexAddress(*checkpointTrackerAddress) {
		log.Fatalf("CheckpointTrackerAddress is not a valid address")
	}

	// Initialize the beacon client
	beaconClient := consensus.NewBeaconClientRPC(*beaconURL)

	// Fetch the beacon block with state replacing the state root
	// This requires access to a debug endpoint on the CL which most RPC providers don't serve due to size
	beaconBlockWithState, err := beaconClient.GetBlockWithState(ctx, *slot)
	if err != nil {
		log.Fatalf("Failed to get beacon block with state: %v", err)
	}

	// Generate the proof inputs
	proofInputs, err := checkpointproofs.GenerateCheckpointProof(ctx, beaconBlockWithState)
	if err != nil {
		log.Fatalf("Proof generation failed: %v", err)
	}
	log.Println("Proof generation completed successfully. Proving branch is length: ", len(proofInputs.Branch))
	log.Printf("Node to be proven is at tree index : %d with finalized checkpoint epoch value %d \n",
		proofInputs.Index, binary.LittleEndian.Uint64(proofInputs.Value[:]))

	// Sanity check the proof is correct
	if err := checkpointproofs.VerifyCheckpointProof(ctx, proofInputs); err != nil {
		log.Fatalf("Proof generation failed: %v", err)
	}
	log.Println("Local proof verification check completed successfully")

	if !*submission {
		return
	}
	// Submit Proof to smart contract
	log.Println("Submitting proof to smart contract")

	checkpointSubmitter, err := checkpointsubmitter.New(ctx, *evmURL, *beaconURL, *checkpointTrackerAddress, *privateKey)
	if err != nil {
		log.Fatalf("Failed to create checkpoint submitter: %v", err)
	}

	receipt, err := checkpointSubmitter.SubmitProof(ctx, proofInputs)
	if err != nil {
		log.Fatalf("Failed to submit proof: %v. Did you already submit a proof for this checkpoint epoch?", err)
	}
	log.Println("Submitted in tx hash: ", receipt.TxHash.String())

	if receipt.Status == 0 {
		log.Warnln("Transaction reverted! ", receipt.TxHash.String())
	} else {
		log.Infoln("Transaction was successul! ", receipt.TxHash.String())
	}

	lastFinalizedTimestamp, err := checkpointSubmitter.LatestFinalizedTimestamp(ctx)
	if err != nil {
		log.Fatalf("Failed to get lastFinalizedTimestamp: %v", err)
	}

	log.Println("Latest FinalizedTimestamp is: ", lastFinalizedTimestamp)
}
