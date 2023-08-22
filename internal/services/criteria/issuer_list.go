package criteria

import (
	"context"
	"github.com/rs/zerolog"
)

type IssuerList struct {
	logger zerolog.Logger
}

func NewIssuerList(l zerolog.Logger) *IssuerList {
	return &IssuerList{
		logger: l,
	}
}

func (i IssuerList) Get(ctx context.Context) (interface{}, error) {
	//TODO call sdk function get issuer list

	return []string{"did", "did"}, nil
}
