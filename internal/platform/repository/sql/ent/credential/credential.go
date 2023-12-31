// Code generated by ent, DO NOT EDIT.

package credential

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the credential type in the database.
	Label = "credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCredentialID holds the string denoting the credential_id field in the database.
	FieldCredentialID = "credential_id"
	// FieldCredentialType holds the string denoting the credential_type field in the database.
	FieldCredentialType = "credential_type"
	// FieldIssuerDid holds the string denoting the issuer_did field in the database.
	FieldIssuerDid = "issuer_did"
	// FieldHolderDid holds the string denoting the holder_did field in the database.
	FieldHolderDid = "holder_did"
	// FieldCredentialData holds the string denoting the credential_data field in the database.
	FieldCredentialData = "credential_data"
	// FieldSignatureProof holds the string denoting the signature_proof field in the database.
	FieldSignatureProof = "signature_proof"
	// FieldSparseMtProof holds the string denoting the sparse_mt_proof field in the database.
	FieldSparseMtProof = "sparse_mt_proof"
	// Table holds the table name of the credential in the database.
	Table = "credentials"
)

// Columns holds all SQL columns for credential fields.
var Columns = []string{
	FieldID,
	FieldCredentialID,
	FieldCredentialType,
	FieldIssuerDid,
	FieldHolderDid,
	FieldCredentialData,
	FieldSignatureProof,
	FieldSparseMtProof,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// CredentialTypeValidator is a validator for the "credential_type" field. It is called by the builders before save.
	CredentialTypeValidator func(string) error
	// IssuerDidValidator is a validator for the "issuer_did" field. It is called by the builders before save.
	IssuerDidValidator func(string) error
	// HolderDidValidator is a validator for the "holder_did" field. It is called by the builders before save.
	HolderDidValidator func(string) error
)

// OrderOption defines the ordering options for the Credential queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCredentialID orders the results by the credential_id field.
func ByCredentialID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredentialID, opts...).ToFunc()
}

// ByCredentialType orders the results by the credential_type field.
func ByCredentialType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCredentialType, opts...).ToFunc()
}

// ByIssuerDid orders the results by the issuer_did field.
func ByIssuerDid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIssuerDid, opts...).ToFunc()
}

// ByHolderDid orders the results by the holder_did field.
func ByHolderDid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHolderDid, opts...).ToFunc()
}
