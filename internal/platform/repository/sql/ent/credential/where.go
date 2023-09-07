// Code generated by ent, DO NOT EDIT.

package credential

import (
	"bloock-identity-managed-api/internal/platform/repository/sql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldID, id))
}

// CredentialID applies equality check predicate on the "credential_id" field. It's identical to CredentialIDEQ.
func CredentialID(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCredentialID, v))
}

// AnchorID applies equality check predicate on the "anchor_id" field. It's identical to AnchorIDEQ.
func AnchorID(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldAnchorID, v))
}

// CredentialType applies equality check predicate on the "credential_type" field. It's identical to CredentialTypeEQ.
func CredentialType(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCredentialType, v))
}

// HolderDid applies equality check predicate on the "holder_did" field. It's identical to HolderDidEQ.
func HolderDid(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldHolderDid, v))
}

// CredentialIDEQ applies the EQ predicate on the "credential_id" field.
func CredentialIDEQ(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCredentialID, v))
}

// CredentialIDNEQ applies the NEQ predicate on the "credential_id" field.
func CredentialIDNEQ(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldCredentialID, v))
}

// CredentialIDIn applies the In predicate on the "credential_id" field.
func CredentialIDIn(vs ...uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldCredentialID, vs...))
}

// CredentialIDNotIn applies the NotIn predicate on the "credential_id" field.
func CredentialIDNotIn(vs ...uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldCredentialID, vs...))
}

// CredentialIDGT applies the GT predicate on the "credential_id" field.
func CredentialIDGT(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldCredentialID, v))
}

// CredentialIDGTE applies the GTE predicate on the "credential_id" field.
func CredentialIDGTE(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldCredentialID, v))
}

// CredentialIDLT applies the LT predicate on the "credential_id" field.
func CredentialIDLT(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldCredentialID, v))
}

// CredentialIDLTE applies the LTE predicate on the "credential_id" field.
func CredentialIDLTE(v uuid.UUID) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldCredentialID, v))
}

// AnchorIDEQ applies the EQ predicate on the "anchor_id" field.
func AnchorIDEQ(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldAnchorID, v))
}

// AnchorIDNEQ applies the NEQ predicate on the "anchor_id" field.
func AnchorIDNEQ(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldAnchorID, v))
}

// AnchorIDIn applies the In predicate on the "anchor_id" field.
func AnchorIDIn(vs ...int64) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldAnchorID, vs...))
}

// AnchorIDNotIn applies the NotIn predicate on the "anchor_id" field.
func AnchorIDNotIn(vs ...int64) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldAnchorID, vs...))
}

// AnchorIDGT applies the GT predicate on the "anchor_id" field.
func AnchorIDGT(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldAnchorID, v))
}

// AnchorIDGTE applies the GTE predicate on the "anchor_id" field.
func AnchorIDGTE(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldAnchorID, v))
}

// AnchorIDLT applies the LT predicate on the "anchor_id" field.
func AnchorIDLT(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldAnchorID, v))
}

// AnchorIDLTE applies the LTE predicate on the "anchor_id" field.
func AnchorIDLTE(v int64) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldAnchorID, v))
}

// CredentialTypeEQ applies the EQ predicate on the "credential_type" field.
func CredentialTypeEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldCredentialType, v))
}

// CredentialTypeNEQ applies the NEQ predicate on the "credential_type" field.
func CredentialTypeNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldCredentialType, v))
}

// CredentialTypeIn applies the In predicate on the "credential_type" field.
func CredentialTypeIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldCredentialType, vs...))
}

// CredentialTypeNotIn applies the NotIn predicate on the "credential_type" field.
func CredentialTypeNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldCredentialType, vs...))
}

// CredentialTypeGT applies the GT predicate on the "credential_type" field.
func CredentialTypeGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldCredentialType, v))
}

// CredentialTypeGTE applies the GTE predicate on the "credential_type" field.
func CredentialTypeGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldCredentialType, v))
}

// CredentialTypeLT applies the LT predicate on the "credential_type" field.
func CredentialTypeLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldCredentialType, v))
}

// CredentialTypeLTE applies the LTE predicate on the "credential_type" field.
func CredentialTypeLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldCredentialType, v))
}

// CredentialTypeContains applies the Contains predicate on the "credential_type" field.
func CredentialTypeContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldCredentialType, v))
}

// CredentialTypeHasPrefix applies the HasPrefix predicate on the "credential_type" field.
func CredentialTypeHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldCredentialType, v))
}

// CredentialTypeHasSuffix applies the HasSuffix predicate on the "credential_type" field.
func CredentialTypeHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldCredentialType, v))
}

// CredentialTypeEqualFold applies the EqualFold predicate on the "credential_type" field.
func CredentialTypeEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldCredentialType, v))
}

// CredentialTypeContainsFold applies the ContainsFold predicate on the "credential_type" field.
func CredentialTypeContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldCredentialType, v))
}

// HolderDidEQ applies the EQ predicate on the "holder_did" field.
func HolderDidEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEQ(FieldHolderDid, v))
}

// HolderDidNEQ applies the NEQ predicate on the "holder_did" field.
func HolderDidNEQ(v string) predicate.Credential {
	return predicate.Credential(sql.FieldNEQ(FieldHolderDid, v))
}

// HolderDidIn applies the In predicate on the "holder_did" field.
func HolderDidIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldIn(FieldHolderDid, vs...))
}

// HolderDidNotIn applies the NotIn predicate on the "holder_did" field.
func HolderDidNotIn(vs ...string) predicate.Credential {
	return predicate.Credential(sql.FieldNotIn(FieldHolderDid, vs...))
}

// HolderDidGT applies the GT predicate on the "holder_did" field.
func HolderDidGT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGT(FieldHolderDid, v))
}

// HolderDidGTE applies the GTE predicate on the "holder_did" field.
func HolderDidGTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldGTE(FieldHolderDid, v))
}

// HolderDidLT applies the LT predicate on the "holder_did" field.
func HolderDidLT(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLT(FieldHolderDid, v))
}

// HolderDidLTE applies the LTE predicate on the "holder_did" field.
func HolderDidLTE(v string) predicate.Credential {
	return predicate.Credential(sql.FieldLTE(FieldHolderDid, v))
}

// HolderDidContains applies the Contains predicate on the "holder_did" field.
func HolderDidContains(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContains(FieldHolderDid, v))
}

// HolderDidHasPrefix applies the HasPrefix predicate on the "holder_did" field.
func HolderDidHasPrefix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasPrefix(FieldHolderDid, v))
}

// HolderDidHasSuffix applies the HasSuffix predicate on the "holder_did" field.
func HolderDidHasSuffix(v string) predicate.Credential {
	return predicate.Credential(sql.FieldHasSuffix(FieldHolderDid, v))
}

// HolderDidEqualFold applies the EqualFold predicate on the "holder_did" field.
func HolderDidEqualFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldEqualFold(FieldHolderDid, v))
}

// HolderDidContainsFold applies the ContainsFold predicate on the "holder_did" field.
func HolderDidContainsFold(v string) predicate.Credential {
	return predicate.Credential(sql.FieldContainsFold(FieldHolderDid, v))
}

// ProofTypeIsNil applies the IsNil predicate on the "proof_type" field.
func ProofTypeIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldProofType))
}

// ProofTypeNotNil applies the NotNil predicate on the "proof_type" field.
func ProofTypeNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldProofType))
}

// IntegrityProofIsNil applies the IsNil predicate on the "integrity_proof" field.
func IntegrityProofIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldIntegrityProof))
}

// IntegrityProofNotNil applies the NotNil predicate on the "integrity_proof" field.
func IntegrityProofNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldIntegrityProof))
}

// SparseMtProofIsNil applies the IsNil predicate on the "sparse_mt_proof" field.
func SparseMtProofIsNil() predicate.Credential {
	return predicate.Credential(sql.FieldIsNull(FieldSparseMtProof))
}

// SparseMtProofNotNil applies the NotNil predicate on the "sparse_mt_proof" field.
func SparseMtProofNotNil() predicate.Credential {
	return predicate.Credential(sql.FieldNotNull(FieldSparseMtProof))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Credential) predicate.Credential {
	return predicate.Credential(func(s *sql.Selector) {
		p(s.Not())
	})
}
