# Checkpoint Proof Generator for Ethereum Beacon Chain
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#experimental)
![testnet-deployment](https://github.com/rzmahmood/ethereum-pos-testnet/actions/workflows/testnet-deployment.yml/badge.svg) [![Ceasefire Now](https://badge.techforpalestine.org/default)](https://techforpalestine.org/learn-more)


A Go application to generate and verify Merkle proofs for the BeaconState `finalized_checkpoint`.

## Introduction
Imagine if the EVM could introspect its own finalization. Thanks to EIP4788, it can!

This tool serves as an example for how to connect to an Ethereum Beacon Node, fetches a Beacon Block and its state for a specified slot, and generates a Merkle proof for the finalized checkpoint within that block. Then correctly prove it against a smart contract

## Installation
1. Clone the Repository:
```bash
git clone https://github.com/rzmahmood/eip4788-proof-generator.git
cd eip4788-proof-generator
```

2. Run unit tests against fixed data in `internal/checkpointproofs/testdata`
```bash
go test ./...
```

3. Build the Application:
```bash
go build -o proof-generator ./cmd/proof-generator
```

4. Run a Local PoS Network to End to End test against
See my local testnet setup which now supports EIP4788! https://github.com/rzmahmood/ethereum-pos-testnet . You will need to wait a few minutes under the first block is finalized.
You can query the chain's finalized checkpoint via, which will eventually return a non-zero epoch `curl http://127.0.0.1:4100/eth/v2/debug/beacon/states/head | jq .data.finalized_checkpoint`
Alternatively run a Sepolia Beacon Node as RPC providers do not serve Debug APIs which is required. A script I used with local reth/lighthouse binaries can be found [here](https://github.com/rzmahmood/ethereum-pos-testnet/blob/test-environment/sepolia.sh)


5. A verification smart contract to the network. For the local `ethereum-pos-testnet`, this can be done using the following:
```
forge script
```

## Future Work and Limitations
- WARNING: This will only work on Deneb blocks.
- This project does not consider forward compatability of Proof verification smart contracts, i.e. [EIP7688](https://ethereum-magicians.org/t/eip-7688-forward-compatible-consensus-data-structures/19673/7)
- This project is currently limited in that proving an epoch was finalized isn't particularly useful without also correlating it to a block number which is possible
- This project is not flexible to proving other parts of the BeaconState. However, it could be forked to do so with ease. Perhaps I will do this as a future project
  - Refactor to use Generalized Index 

## Learnings and Hurdles
- Prysm Client for Beaconchain data is terrible. It removed the API Key from the path when querying, meaning I had to manually form paths. Eventually abandoned it.
- umbracle/go-eth-consensus ->  https://github.com/rzmahmood/go-eth-consensus
  - Hasn't been updated to support Deneb so I forked it and added Deneb Support by adding KZG commitments
  - I had to add Debug RPC support to the fork
- My local testnet was failing Beacon State root calculation due to an Epoch size of 6. Fixed it by changing Epoch size to 32





