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
