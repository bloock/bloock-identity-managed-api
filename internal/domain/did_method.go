package domain

import "github.com/bloock/bloock-sdk-go/v2/entity/identityV2"

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

func (p DidMethod) ToBloockMethod() identityV2.Method {
	switch p {
	case DidMethodIden3:
		return identityV2.ListOfMethods().Iden3
	case DidMethodPolygonID:
		return identityV2.ListOfMethods().PolygonId
	default:
		return identityV2.ListOfMethods().PolygonId
	}
}

func GetIssuerParams(method, blockchain, network string) (identityV2.IssuerParams, error) {
	params := identityV2.NewIssuerParams()
	if method != "" && blockchain != "" && network != "" {
		m, err := NewDidMethod(method)
		if err != nil {
			return identityV2.IssuerParams{}, err
		}
		b, err := NewDidBlockchain(blockchain)
		if err != nil {
			return identityV2.IssuerParams{}, err
		}
		n, err := NewDidNetwork(network)
		if err != nil {
			return identityV2.IssuerParams{}, err
		}
		params.Method = m.ToBloockMethod()
		params.Blockchain = b.ToBloockBlockchain()
		params.NetworkId = n.ToBloockNetwork()
	}

	return params, nil
}
