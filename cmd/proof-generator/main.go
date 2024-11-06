package main

import (
	"context"
	"eip4788-proof-generator/internal/checkpointproofs"
	"eip4788-proof-generator/pkg/consensus"
	"flag"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	// Parse command-line flags
	beaconURL := flag.String("beacon-url", "http://127.0.0.1:4100", "Beacon node endpoint URL")
	slot := flag.Uint64("slot", 282, "Beacon block slot to fetch")
	flag.Parse()

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
	fmt.Println("Proof generation completed successfully")

	// Sanity check the proof is correct
	if err := checkpointproofs.VerifyCheckpointProof(ctx, proofInputs); err != nil {
		log.Fatalf("Proof generation failed: %v", err)
	}
	fmt.Println("Proof sanity check completed successfully")

	// Submit Proof to smart contract
	// TODO
}
