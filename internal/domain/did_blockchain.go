package domain

import identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"

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

func (p DidBlockchain) ToBloockBlockchain() identityEntity.Blockchain {
	switch p {
	case DidBlockchainPolygon:
		return identityEntity.ListOfBlockchains().Polygon
	case DidBlockchainEthereum:
		return identityEntity.ListOfBlockchains().Ethereum
	case DidBlockchainNoChain:
		return identityEntity.ListOfBlockchains().NoChain
	case DidBlockchainUnknownChain:
		return identityEntity.ListOfBlockchains().UnknownChain
	default:
		return identityEntity.ListOfBlockchains().Polygon
	}
}
