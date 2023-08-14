package schema

import (
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
		field.String("schema_type").NotEmpty(),
		field.String("issuer_did").NotEmpty(),
		field.String("holder_did").NotEmpty(),
		field.JSON("credential_data", map[string]interface{}{}),
		field.JSON("proofs", map[string]interface{}{}),
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
