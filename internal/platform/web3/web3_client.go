package web3

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/iden3/contracts-abi/state/go/abi"
	"github.com/rs/zerolog"
)

type ClientWeb3 struct {
	client          *ethclient.Client
	contractAddress string
	logger          zerolog.Logger
}

func NewClientWeb3(ctx context.Context, provider string, apiKey string, contractAddress string, logger zerolog.Logger) (ClientWeb3, error) {
	headers := rpc.WithHeader("x-api-key", apiKey)
	rpcClient, err := rpc.DialOptions(ctx, provider, headers)
	if err != nil {
		err = fmt.Errorf("error: connecting with provider: %s", provider)
		logger.Error().Err(err).Msg("")
		return ClientWeb3{}, err
	}

	return ClientWeb3{
		client:          ethclient.NewClient(rpcClient),
		contractAddress: contractAddress,
		logger:          logger,
	}, nil
}

func (e ClientWeb3) GetAbiState() (*abi.State, error) {
	stateContractInstance, err := abi.NewState(common.HexToAddress(e.contractAddress), e.client)
	if err != nil {
		err = fmt.Errorf("error failed create state contract client: %w", err)
		e.logger.Error().Err(err).Msg("")
		return nil, err
	}

	return stateContractInstance, nil
}
