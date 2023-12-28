package web3

import (
	"bloock-identity-managed-api/internal/config"
	"bloock-identity-managed-api/internal/pkg"
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

func NewClientWeb3(ctx context.Context, logger zerolog.Logger) (ClientWeb3, error) {
	var headers rpc.ClientOption
	provider := config.Configuration.Blockchain.Provider

	if pkg.GetApiKeyFromContext(ctx) == "" {
		headers = rpc.WithHeader("x-api-key", pkg.GetApiKeyFromContext(ctx))
		provider = config.PublicPolygonProvider
	}

	rpcClient, err := rpc.DialOptions(ctx, provider, headers)
	if err != nil {
		err = fmt.Errorf("error: connecting with provider: %s", provider)
		logger.Error().Err(err).Msg("")
		return ClientWeb3{}, err
	}

	return ClientWeb3{
		client:          ethclient.NewClient(rpcClient),
		contractAddress: config.Configuration.Blockchain.SmartContract,
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
