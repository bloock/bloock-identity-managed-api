package domain

import "github.com/bloock/bloock-sdk-go/v2/entity/identity"

type DidMethod int32

const (
	PolygonID DidMethod = iota
	PolygonIDTest
)

func NewDidMethod(method string) (DidMethod, error) {
	switch method {
	case "polygon_id":
		return PolygonID, nil
	case "polygon_id_test":
		return PolygonIDTest, nil
	default:
		return -1, ErrInvalidDidMethod
	}
}

func (d DidMethod) Str() string {
	switch d {
	case PolygonID:
		return "polygon_id"
	case PolygonIDTest:
		return "polygon_id_test"
	default:
		return ""
	}
}

func (d DidMethod) GetBloockDidMethod() identity.DidMethod {
	switch d {
	case PolygonID:
		return identity.PolygonID
	case PolygonIDTest:
		return identity.PolygonIDTest
	default:
		return identity.PolygonID
	}
}
