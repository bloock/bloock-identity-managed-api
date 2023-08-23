package domain

import "github.com/bloock/bloock-sdk-go/v2/entity/identityV2"

type DidBlockchain int32

const (
	DidBlockchainPolygon DidBlockchain = iota
	DidBlockchainEthereum
	DidBlockchainNoChain
	DidBlockchainUnknownChain
)

func NewDidBlockchain(blockchain string) (DidBlockchain, error) {
	switch blockchain {
	case "polygon":
		return DidBlockchainPolygon, nil
	case "ethereum":
		return DidBlockchainEthereum, nil
	case "no_chain":
		return DidBlockchainNoChain, nil
	case "unknown_chain":
		return DidBlockchainUnknownChain, nil
	default:
		return 0, ErrInvalidBlockchainProvided
	}
}

func (p DidBlockchain) ToBloockBlockchain() identityV2.Blockchain {
	switch p {
	case DidBlockchainPolygon:
		return identityV2.ListOfBlockchains().Polygon
	case DidBlockchainEthereum:
		return identityV2.ListOfBlockchains().Ethereum
	case DidBlockchainNoChain:
		return identityV2.ListOfBlockchains().NoChain
	case DidBlockchainUnknownChain:
		return identityV2.ListOfBlockchains().UnknownChain
	default:
		return identityV2.ListOfBlockchains().Polygon
	}
}
