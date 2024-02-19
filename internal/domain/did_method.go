package domain

import (
	"bloock-identity-managed-api/internal/config"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
)

type DidMethod int32

const (
	DidMethodIden3 DidMethod = iota
	DidMethodPolygonID
)

func NewDidMethod(method string) (DidMethod, error) {
	switch method {
	case "iden3":
		return DidMethodIden3, nil
	case "polygon_id":
		return DidMethodPolygonID, nil
	default:
		return 0, ErrInvalidMethodProvided
	}
}

func (p DidMethod) ToBloockMethod() identityEntity.Method {
	switch p {
	case DidMethodIden3:
		return identityEntity.ListOfMethods().Iden3
	case DidMethodPolygonID:
		return identityEntity.ListOfMethods().PolygonId
	default:
		return identityEntity.ListOfMethods().PolygonId
	}
}

func GetDidType(method, blockchain, network string) (identityEntity.DidType, error) {
	didType := identityEntity.NewDidType()
	if method != "" && blockchain != "" && network != "" {
		m, err := NewDidMethod(method)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		b, err := NewDidBlockchain(blockchain)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		n, err := NewDidNetwork(network)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		didType.Method = m.ToBloockMethod()
		didType.Blockchain = b.ToBloockBlockchain()
		didType.NetworkId = n.ToBloockNetwork()
	} else if config.Configuration.Issuer.DidMetadata.Method != "" {
		m, err := NewDidMethod(config.Configuration.Issuer.DidMetadata.Method)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		b, err := NewDidBlockchain(config.Configuration.Issuer.DidMetadata.Blockchain)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		n, err := NewDidNetwork(config.Configuration.Issuer.DidMetadata.Network)
		if err != nil {
			return identityEntity.DidType{}, err
		}
		didType.Method = m.ToBloockMethod()
		didType.Method = m.ToBloockMethod()
		didType.Blockchain = b.ToBloockBlockchain()
		didType.NetworkId = n.ToBloockNetwork()
	}

	return didType, nil
}
