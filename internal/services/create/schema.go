package create

import (
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"github.com/rs/zerolog"
)

type Schema struct {
	logger zerolog.Logger
}

func NewSchema(l zerolog.Logger) *Schema {
	return &Schema{
		logger: l,
	}
}

func (s Schema) Create(ctx context.Context, req request.CreateSchemaRequest) (interface{}, error) {

	//TODO call create schema sdk function

	return "schema_id", nil
}
