package schema

import (
	"encoding/json"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type Credential struct {
	ent.Schema
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("credential_id", uuid.UUID{}),
		field.String("credential_type").NotEmpty(),
		field.String("issuer_did").NotEmpty(),
		field.String("holder_did").NotEmpty(),
		field.JSON("credential_data", json.RawMessage{}),
		field.JSON("signature_proof", json.RawMessage{}),
		field.JSON("sparse_mt_proof", json.RawMessage{}).Optional(),
	}
}

// Edges of the Certification.
func (Credential) Edges() []ent.Edge {
	return nil
}

func (Credential) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("credential_id").Unique(),
	}
}
