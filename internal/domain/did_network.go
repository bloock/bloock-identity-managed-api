package domain

import identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"

type DidNetwork int32

const (
	DidNetworkGoerli DidNetwork = iota
	DidNetworkMain
	DidNetworkMumbai
	DidNetworkNoNetwork
	DidNetworkUnknownNetwork
)

func NewDidNetwork(network string) (DidNetwork, error) {
	switch network {
	case "goerli":
		return DidNetworkGoerli, nil
	case "main":
		return DidNetworkMain, nil
	case "mumbai":
		return DidNetworkMumbai, nil
	case "no_network":
		return DidNetworkNoNetwork, nil
	case "unknown_network":
		return DidNetworkUnknownNetwork, nil
	default:
		return 0, ErrInvalidBlockchainProvided
	}
}

func (p DidNetwork) ToBloockNetwork() identityEntity.NetworkId {
	switch p {
	case DidNetworkGoerli:
		return identityEntity.ListOfNetworkIds().Goerli
	case DidNetworkMain:
		return identityEntity.ListOfNetworkIds().Main
	case DidNetworkMumbai:
		return identityEntity.ListOfNetworkIds().Mumbai
	case DidNetworkNoNetwork:
		return identityEntity.ListOfNetworkIds().NoNetwork
	case DidNetworkUnknownNetwork:
		return identityEntity.ListOfNetworkIds().UnknownNetwork
	default:
		return identityEntity.ListOfNetworkIds().Mumbai
	}
}
