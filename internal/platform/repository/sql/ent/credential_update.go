// Code generated by ent, DO NOT EDIT.

package ent

import (
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/credential"
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/predicate"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CredentialUpdate is the builder for updating Credential entities.
type CredentialUpdate struct {
	config
	hooks    []Hook
	mutation *CredentialMutation
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cu *CredentialUpdate) Where(ps ...predicate.Credential) *CredentialUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCredentialID sets the "credential_id" field.
func (cu *CredentialUpdate) SetCredentialID(u uuid.UUID) *CredentialUpdate {
	cu.mutation.SetCredentialID(u)
	return cu
}

// SetAnchorID sets the "anchor_id" field.
func (cu *CredentialUpdate) SetAnchorID(i int64) *CredentialUpdate {
	cu.mutation.ResetAnchorID()
	cu.mutation.SetAnchorID(i)
	return cu
}

// AddAnchorID adds i to the "anchor_id" field.
func (cu *CredentialUpdate) AddAnchorID(i int64) *CredentialUpdate {
	cu.mutation.AddAnchorID(i)
	return cu
}

// SetSchemaType sets the "schema_type" field.
func (cu *CredentialUpdate) SetSchemaType(s string) *CredentialUpdate {
	cu.mutation.SetSchemaType(s)
	return cu
}

// SetIssuerDid sets the "issuer_did" field.
func (cu *CredentialUpdate) SetIssuerDid(s string) *CredentialUpdate {
	cu.mutation.SetIssuerDid(s)
	return cu
}

// SetHolderDid sets the "holder_did" field.
func (cu *CredentialUpdate) SetHolderDid(s string) *CredentialUpdate {
	cu.mutation.SetHolderDid(s)
	return cu
}

// SetProofType sets the "proof_type" field.
func (cu *CredentialUpdate) SetProofType(s []string) *CredentialUpdate {
	cu.mutation.SetProofType(s)
	return cu
}

// AppendProofType appends s to the "proof_type" field.
func (cu *CredentialUpdate) AppendProofType(s []string) *CredentialUpdate {
	cu.mutation.AppendProofType(s)
	return cu
}

// ClearProofType clears the value of the "proof_type" field.
func (cu *CredentialUpdate) ClearProofType() *CredentialUpdate {
	cu.mutation.ClearProofType()
	return cu
}

// SetCredentialData sets the "credential_data" field.
func (cu *CredentialUpdate) SetCredentialData(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.SetCredentialData(jm)
	return cu
}

// AppendCredentialData appends jm to the "credential_data" field.
func (cu *CredentialUpdate) AppendCredentialData(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.AppendCredentialData(jm)
	return cu
}

// SetSignatureProof sets the "signature_proof" field.
func (cu *CredentialUpdate) SetSignatureProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.SetSignatureProof(jm)
	return cu
}

// AppendSignatureProof appends jm to the "signature_proof" field.
func (cu *CredentialUpdate) AppendSignatureProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.AppendSignatureProof(jm)
	return cu
}

// SetIntegrityProof sets the "integrity_proof" field.
func (cu *CredentialUpdate) SetIntegrityProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.SetIntegrityProof(jm)
	return cu
}

// AppendIntegrityProof appends jm to the "integrity_proof" field.
func (cu *CredentialUpdate) AppendIntegrityProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.AppendIntegrityProof(jm)
	return cu
}

// ClearIntegrityProof clears the value of the "integrity_proof" field.
func (cu *CredentialUpdate) ClearIntegrityProof() *CredentialUpdate {
	cu.mutation.ClearIntegrityProof()
	return cu
}

// SetSparseMtProof sets the "sparse_mt_proof" field.
func (cu *CredentialUpdate) SetSparseMtProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.SetSparseMtProof(jm)
	return cu
}

// AppendSparseMtProof appends jm to the "sparse_mt_proof" field.
func (cu *CredentialUpdate) AppendSparseMtProof(jm json.RawMessage) *CredentialUpdate {
	cu.mutation.AppendSparseMtProof(jm)
	return cu
}

// ClearSparseMtProof clears the value of the "sparse_mt_proof" field.
func (cu *CredentialUpdate) ClearSparseMtProof() *CredentialUpdate {
	cu.mutation.ClearSparseMtProof()
	return cu
}

// Mutation returns the CredentialMutation object of the builder.
func (cu *CredentialUpdate) Mutation() *CredentialMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CredentialUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CredentialUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CredentialUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CredentialUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CredentialUpdate) check() error {
	if v, ok := cu.mutation.SchemaType(); ok {
		if err := credential.SchemaTypeValidator(v); err != nil {
			return &ValidationError{Name: "schema_type", err: fmt.Errorf(`ent: validator failed for field "Credential.schema_type": %w`, err)}
		}
	}
	if v, ok := cu.mutation.IssuerDid(); ok {
		if err := credential.IssuerDidValidator(v); err != nil {
			return &ValidationError{Name: "issuer_did", err: fmt.Errorf(`ent: validator failed for field "Credential.issuer_did": %w`, err)}
		}
	}
	if v, ok := cu.mutation.HolderDid(); ok {
		if err := credential.HolderDidValidator(v); err != nil {
			return &ValidationError{Name: "holder_did", err: fmt.Errorf(`ent: validator failed for field "Credential.holder_did": %w`, err)}
		}
	}
	return nil
}

func (cu *CredentialUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(credential.Table, credential.Columns, sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CredentialID(); ok {
		_spec.SetField(credential.FieldCredentialID, field.TypeUUID, value)
	}
	if value, ok := cu.mutation.AnchorID(); ok {
		_spec.SetField(credential.FieldAnchorID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedAnchorID(); ok {
		_spec.AddField(credential.FieldAnchorID, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.SchemaType(); ok {
		_spec.SetField(credential.FieldSchemaType, field.TypeString, value)
	}
	if value, ok := cu.mutation.IssuerDid(); ok {
		_spec.SetField(credential.FieldIssuerDid, field.TypeString, value)
	}
	if value, ok := cu.mutation.HolderDid(); ok {
		_spec.SetField(credential.FieldHolderDid, field.TypeString, value)
	}
	if value, ok := cu.mutation.ProofType(); ok {
		_spec.SetField(credential.FieldProofType, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedProofType(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldProofType, value)
		})
	}
	if cu.mutation.ProofTypeCleared() {
		_spec.ClearField(credential.FieldProofType, field.TypeJSON)
	}
	if value, ok := cu.mutation.CredentialData(); ok {
		_spec.SetField(credential.FieldCredentialData, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedCredentialData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldCredentialData, value)
		})
	}
	if value, ok := cu.mutation.SignatureProof(); ok {
		_spec.SetField(credential.FieldSignatureProof, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedSignatureProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldSignatureProof, value)
		})
	}
	if value, ok := cu.mutation.IntegrityProof(); ok {
		_spec.SetField(credential.FieldIntegrityProof, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedIntegrityProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldIntegrityProof, value)
		})
	}
	if cu.mutation.IntegrityProofCleared() {
		_spec.ClearField(credential.FieldIntegrityProof, field.TypeJSON)
	}
	if value, ok := cu.mutation.SparseMtProof(); ok {
		_spec.SetField(credential.FieldSparseMtProof, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedSparseMtProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldSparseMtProof, value)
		})
	}
	if cu.mutation.SparseMtProofCleared() {
		_spec.ClearField(credential.FieldSparseMtProof, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CredentialUpdateOne is the builder for updating a single Credential entity.
type CredentialUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CredentialMutation
}

// SetCredentialID sets the "credential_id" field.
func (cuo *CredentialUpdateOne) SetCredentialID(u uuid.UUID) *CredentialUpdateOne {
	cuo.mutation.SetCredentialID(u)
	return cuo
}

// SetAnchorID sets the "anchor_id" field.
func (cuo *CredentialUpdateOne) SetAnchorID(i int64) *CredentialUpdateOne {
	cuo.mutation.ResetAnchorID()
	cuo.mutation.SetAnchorID(i)
	return cuo
}

// AddAnchorID adds i to the "anchor_id" field.
func (cuo *CredentialUpdateOne) AddAnchorID(i int64) *CredentialUpdateOne {
	cuo.mutation.AddAnchorID(i)
	return cuo
}

// SetSchemaType sets the "schema_type" field.
func (cuo *CredentialUpdateOne) SetSchemaType(s string) *CredentialUpdateOne {
	cuo.mutation.SetSchemaType(s)
	return cuo
}

// SetIssuerDid sets the "issuer_did" field.
func (cuo *CredentialUpdateOne) SetIssuerDid(s string) *CredentialUpdateOne {
	cuo.mutation.SetIssuerDid(s)
	return cuo
}

// SetHolderDid sets the "holder_did" field.
func (cuo *CredentialUpdateOne) SetHolderDid(s string) *CredentialUpdateOne {
	cuo.mutation.SetHolderDid(s)
	return cuo
}

// SetProofType sets the "proof_type" field.
func (cuo *CredentialUpdateOne) SetProofType(s []string) *CredentialUpdateOne {
	cuo.mutation.SetProofType(s)
	return cuo
}

// AppendProofType appends s to the "proof_type" field.
func (cuo *CredentialUpdateOne) AppendProofType(s []string) *CredentialUpdateOne {
	cuo.mutation.AppendProofType(s)
	return cuo
}

// ClearProofType clears the value of the "proof_type" field.
func (cuo *CredentialUpdateOne) ClearProofType() *CredentialUpdateOne {
	cuo.mutation.ClearProofType()
	return cuo
}

// SetCredentialData sets the "credential_data" field.
func (cuo *CredentialUpdateOne) SetCredentialData(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.SetCredentialData(jm)
	return cuo
}

// AppendCredentialData appends jm to the "credential_data" field.
func (cuo *CredentialUpdateOne) AppendCredentialData(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.AppendCredentialData(jm)
	return cuo
}

// SetSignatureProof sets the "signature_proof" field.
func (cuo *CredentialUpdateOne) SetSignatureProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.SetSignatureProof(jm)
	return cuo
}

// AppendSignatureProof appends jm to the "signature_proof" field.
func (cuo *CredentialUpdateOne) AppendSignatureProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.AppendSignatureProof(jm)
	return cuo
}

// SetIntegrityProof sets the "integrity_proof" field.
func (cuo *CredentialUpdateOne) SetIntegrityProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.SetIntegrityProof(jm)
	return cuo
}

// AppendIntegrityProof appends jm to the "integrity_proof" field.
func (cuo *CredentialUpdateOne) AppendIntegrityProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.AppendIntegrityProof(jm)
	return cuo
}

// ClearIntegrityProof clears the value of the "integrity_proof" field.
func (cuo *CredentialUpdateOne) ClearIntegrityProof() *CredentialUpdateOne {
	cuo.mutation.ClearIntegrityProof()
	return cuo
}

// SetSparseMtProof sets the "sparse_mt_proof" field.
func (cuo *CredentialUpdateOne) SetSparseMtProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.SetSparseMtProof(jm)
	return cuo
}

// AppendSparseMtProof appends jm to the "sparse_mt_proof" field.
func (cuo *CredentialUpdateOne) AppendSparseMtProof(jm json.RawMessage) *CredentialUpdateOne {
	cuo.mutation.AppendSparseMtProof(jm)
	return cuo
}

// ClearSparseMtProof clears the value of the "sparse_mt_proof" field.
func (cuo *CredentialUpdateOne) ClearSparseMtProof() *CredentialUpdateOne {
	cuo.mutation.ClearSparseMtProof()
	return cuo
}

// Mutation returns the CredentialMutation object of the builder.
func (cuo *CredentialUpdateOne) Mutation() *CredentialMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CredentialUpdate builder.
func (cuo *CredentialUpdateOne) Where(ps ...predicate.Credential) *CredentialUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CredentialUpdateOne) Select(field string, fields ...string) *CredentialUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Credential entity.
func (cuo *CredentialUpdateOne) Save(ctx context.Context) (*Credential, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CredentialUpdateOne) SaveX(ctx context.Context) *Credential {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CredentialUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CredentialUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CredentialUpdateOne) check() error {
	if v, ok := cuo.mutation.SchemaType(); ok {
		if err := credential.SchemaTypeValidator(v); err != nil {
			return &ValidationError{Name: "schema_type", err: fmt.Errorf(`ent: validator failed for field "Credential.schema_type": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.IssuerDid(); ok {
		if err := credential.IssuerDidValidator(v); err != nil {
			return &ValidationError{Name: "issuer_did", err: fmt.Errorf(`ent: validator failed for field "Credential.issuer_did": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.HolderDid(); ok {
		if err := credential.HolderDidValidator(v); err != nil {
			return &ValidationError{Name: "holder_did", err: fmt.Errorf(`ent: validator failed for field "Credential.holder_did": %w`, err)}
		}
	}
	return nil
}

func (cuo *CredentialUpdateOne) sqlSave(ctx context.Context) (_node *Credential, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(credential.Table, credential.Columns, sqlgraph.NewFieldSpec(credential.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Credential.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, credential.FieldID)
		for _, f := range fields {
			if !credential.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != credential.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CredentialID(); ok {
		_spec.SetField(credential.FieldCredentialID, field.TypeUUID, value)
	}
	if value, ok := cuo.mutation.AnchorID(); ok {
		_spec.SetField(credential.FieldAnchorID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedAnchorID(); ok {
		_spec.AddField(credential.FieldAnchorID, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.SchemaType(); ok {
		_spec.SetField(credential.FieldSchemaType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IssuerDid(); ok {
		_spec.SetField(credential.FieldIssuerDid, field.TypeString, value)
	}
	if value, ok := cuo.mutation.HolderDid(); ok {
		_spec.SetField(credential.FieldHolderDid, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ProofType(); ok {
		_spec.SetField(credential.FieldProofType, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedProofType(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldProofType, value)
		})
	}
	if cuo.mutation.ProofTypeCleared() {
		_spec.ClearField(credential.FieldProofType, field.TypeJSON)
	}
	if value, ok := cuo.mutation.CredentialData(); ok {
		_spec.SetField(credential.FieldCredentialData, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedCredentialData(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldCredentialData, value)
		})
	}
	if value, ok := cuo.mutation.SignatureProof(); ok {
		_spec.SetField(credential.FieldSignatureProof, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedSignatureProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldSignatureProof, value)
		})
	}
	if value, ok := cuo.mutation.IntegrityProof(); ok {
		_spec.SetField(credential.FieldIntegrityProof, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedIntegrityProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldIntegrityProof, value)
		})
	}
	if cuo.mutation.IntegrityProofCleared() {
		_spec.ClearField(credential.FieldIntegrityProof, field.TypeJSON)
	}
	if value, ok := cuo.mutation.SparseMtProof(); ok {
		_spec.SetField(credential.FieldSparseMtProof, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedSparseMtProof(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, credential.FieldSparseMtProof, value)
		})
	}
	if cuo.mutation.SparseMtProofCleared() {
		_spec.ClearField(credential.FieldSparseMtProof, field.TypeJSON)
	}
	_node = &Credential{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{credential.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}