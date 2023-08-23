package domain

import "github.com/bloock/bloock-sdk-go/v2/entity/identityV2"

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

func (p DidNetwork) ToBloockNetwork() identityV2.NetworkId {
	switch p {
	case DidNetworkGoerli:
		return identityV2.ListOfNetworkIds().Goerli
	case DidNetworkMain:
		return identityV2.ListOfNetworkIds().Main
	case DidNetworkMumbai:
		return identityV2.ListOfNetworkIds().Mumbai
	case DidNetworkNoNetwork:
		return identityV2.ListOfNetworkIds().NoNetwork
	case DidNetworkUnknownNetwork:
		return identityV2.ListOfNetworkIds().UnknownNetwork
	default:
		return identityV2.ListOfNetworkIds().Mumbai
	}
}
