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
		field.Int64("anchor_id"),
		field.String("credential_type").NotEmpty(),
		field.String("holder_did").NotEmpty(),
		field.Strings("proof_type").Optional(),
		field.JSON("credential_data", json.RawMessage{}),
		field.JSON("signature_proof", json.RawMessage{}),
		field.JSON("integrity_proof", json.RawMessage{}).Optional(),
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
