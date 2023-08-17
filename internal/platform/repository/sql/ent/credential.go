// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/credential"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Credential is the model entity for the Credential schema.
type Credential struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CredentialID holds the value of the "credential_id" field.
	CredentialID uuid.UUID `json:"credential_id,omitempty"`
	// AnchorID holds the value of the "anchor_id" field.
	AnchorID int64 `json:"anchor_id,omitempty"`
	// SchemaType holds the value of the "schema_type" field.
	SchemaType string `json:"schema_type,omitempty"`
	// IssuerDid holds the value of the "issuer_did" field.
	IssuerDid string `json:"issuer_did,omitempty"`
	// HolderDid holds the value of the "holder_did" field.
	HolderDid string `json:"holder_did,omitempty"`
	// ProofType holds the value of the "proof_type" field.
	ProofType []string `json:"proof_type,omitempty"`
	// CredentialData holds the value of the "credential_data" field.
	CredentialData json.RawMessage `json:"credential_data,omitempty"`
	// SignatureProof holds the value of the "signature_proof" field.
	SignatureProof json.RawMessage `json:"signature_proof,omitempty"`
	// IntegrityProof holds the value of the "integrity_proof" field.
	IntegrityProof json.RawMessage `json:"integrity_proof,omitempty"`
	// SparseMtProof holds the value of the "sparse_mt_proof" field.
	SparseMtProof json.RawMessage `json:"sparse_mt_proof,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Credential) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case credential.FieldProofType, credential.FieldCredentialData, credential.FieldSignatureProof, credential.FieldIntegrityProof, credential.FieldSparseMtProof:
			values[i] = new([]byte)
		case credential.FieldID, credential.FieldAnchorID:
			values[i] = new(sql.NullInt64)
		case credential.FieldSchemaType, credential.FieldIssuerDid, credential.FieldHolderDid:
			values[i] = new(sql.NullString)
		case credential.FieldCredentialID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Credential fields.
func (c *Credential) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case credential.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case credential.FieldCredentialID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field credential_id", values[i])
			} else if value != nil {
				c.CredentialID = *value
			}
		case credential.FieldAnchorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field anchor_id", values[i])
			} else if value.Valid {
				c.AnchorID = value.Int64
			}
		case credential.FieldSchemaType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field schema_type", values[i])
			} else if value.Valid {
				c.SchemaType = value.String
			}
		case credential.FieldIssuerDid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field issuer_did", values[i])
			} else if value.Valid {
				c.IssuerDid = value.String
			}
		case credential.FieldHolderDid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field holder_did", values[i])
			} else if value.Valid {
				c.HolderDid = value.String
			}
		case credential.FieldProofType:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field proof_type", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.ProofType); err != nil {
					return fmt.Errorf("unmarshal field proof_type: %w", err)
				}
			}
		case credential.FieldCredentialData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field credential_data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.CredentialData); err != nil {
					return fmt.Errorf("unmarshal field credential_data: %w", err)
				}
			}
		case credential.FieldSignatureProof:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature_proof", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.SignatureProof); err != nil {
					return fmt.Errorf("unmarshal field signature_proof: %w", err)
				}
			}
		case credential.FieldIntegrityProof:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field integrity_proof", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.IntegrityProof); err != nil {
					return fmt.Errorf("unmarshal field integrity_proof: %w", err)
				}
			}
		case credential.FieldSparseMtProof:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sparse_mt_proof", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.SparseMtProof); err != nil {
					return fmt.Errorf("unmarshal field sparse_mt_proof: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Credential.
// This includes values selected through modifiers, order, etc.
func (c *Credential) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Credential.
// Note that you need to call Credential.Unwrap() before calling this method if this Credential
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Credential) Update() *CredentialUpdateOne {
	return NewCredentialClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Credential entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Credential) Unwrap() *Credential {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Credential is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Credential) String() string {
	var builder strings.Builder
	builder.WriteString("Credential(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("credential_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CredentialID))
	builder.WriteString(", ")
	builder.WriteString("anchor_id=")
	builder.WriteString(fmt.Sprintf("%v", c.AnchorID))
	builder.WriteString(", ")
	builder.WriteString("schema_type=")
	builder.WriteString(c.SchemaType)
	builder.WriteString(", ")
	builder.WriteString("issuer_did=")
	builder.WriteString(c.IssuerDid)
	builder.WriteString(", ")
	builder.WriteString("holder_did=")
	builder.WriteString(c.HolderDid)
	builder.WriteString(", ")
	builder.WriteString("proof_type=")
	builder.WriteString(fmt.Sprintf("%v", c.ProofType))
	builder.WriteString(", ")
	builder.WriteString("credential_data=")
	builder.WriteString(fmt.Sprintf("%v", c.CredentialData))
	builder.WriteString(", ")
	builder.WriteString("signature_proof=")
	builder.WriteString(fmt.Sprintf("%v", c.SignatureProof))
	builder.WriteString(", ")
	builder.WriteString("integrity_proof=")
	builder.WriteString(fmt.Sprintf("%v", c.IntegrityProof))
	builder.WriteString(", ")
	builder.WriteString("sparse_mt_proof=")
	builder.WriteString(fmt.Sprintf("%v", c.SparseMtProof))
	builder.WriteByte(')')
	return builder.String()
}

// Credentials is a parsable slice of Credential.
type Credentials []*Credential