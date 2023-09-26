package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/iden3/contracts-abi/state/go/abi"
	ethState "github.com/iden3/go-iden3-auth/v2/state"
	"math/big"
)

// ETHResolver resolver for eth blockchains
type BloockNodeResolver struct {
	RPCUrl          string
	ApiKey          string
	ContractAddress common.Address
}

// NewETHResolver create ETH resolver for state.
func NewBloockNodeResolver(url, apiKey, contract string) *BloockNodeResolver {
	return &BloockNodeResolver{
		RPCUrl:          url,
		ApiKey:          apiKey,
		ContractAddress: common.HexToAddress(contract),
	}
}

// Resolve returns Resolved state from blockchain
func (r BloockNodeResolver) Resolve(ctx context.Context, id, state *big.Int) (*ethState.ResolvedState, error) {
	headers := rpc.WithHeader("x-api-key", r.ApiKey)
	rpcClient, err := rpc.DialOptions(ctx, r.RPCUrl, headers)
	if err != nil {
		err = fmt.Errorf("error: connecting with provider: %s", r.RPCUrl)
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)
	defer client.Close()

	getter, err := abi.NewStateCaller(r.ContractAddress, client)
	if err != nil {
		return nil, err
	}
	return ethState.Resolve(ctx, getter, id, state)
}

// ResolveGlobalRoot returns Resolved global state from blockchain
func (r BloockNodeResolver) ResolveGlobalRoot(ctx context.Context, state *big.Int) (*ethState.ResolvedState, error) {
	headers := rpc.WithHeader("x-api-key", r.ApiKey)
	rpcClient, err := rpc.DialOptions(ctx, r.RPCUrl, headers)
	if err != nil {
		err = fmt.Errorf("error: connecting with provider: %s", r.RPCUrl)
		return nil, err
	}

	client := ethclient.NewClient(rpcClient)
	defer client.Close()

	getter, err := abi.NewStateCaller(r.ContractAddress, client)
	if err != nil {
		return nil, err
	}
	return ethState.ResolveGlobalRoot(ctx, getter, state)
}
