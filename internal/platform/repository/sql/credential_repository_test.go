package sql

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/platform/repository/sql/connection"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSQLCredentialRepository(t *testing.T) {
	entConnector := connection.NewEntConnector(zerolog.Logger{})
	conn, err := connection.NewEntConnection("file:ent?mode=memory&cache=shared&_fk=1", entConnector, zerolog.Logger{})
	require.NoError(t, err)
	err = conn.Migrate(context.Background())
	require.NoError(t, err)

	credentialId, err := uuid.NewUUID()
	require.NoError(t, err)
	credentialType := "DrivingLicense"
	issuerDid := "did:polygonid:polygon:mumbai:2qGg7TzmcoU4Jg3E86wXp4WJcyGUTuafPZxVRxpYQr"
	holderDid := "did:polygonid:polygon:mumbai:2qGg7TzmcoU4Jg3E86wXp4WJcyGUTuafPZxVRxpYQr"
	cred := `{"@context":["https://www.w3.org/2018/credentials/v1","https://schema.iden3.io/core/jsonld/iden3proofs.jsonld","https://api.bloock.dev/hosting/v1/ipfs/QmZ9BzmMGzLv4y9P6djYUm8sgQt47ZjECGAM8nToFW2qvt"],"credentialSchema":{"id":"https://api.bloock.dev/hosting/v1/ipfs/QmTvHzXiegijCdhGC4kgjps8hSi3FP1K17ezrYPgdMU6Ek","type":"JsonSchema2023"},"credentialStatus":{"id":"https://api.bloock.dev/identity/v1/did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U/claims/revocation/status/3825417065","revocationNonce":3825417065,"type":"SparseMerkleTreeProof"},"credentialSubject":{"birth_date":921950325,"country":"Spain","name":"Eduard","type":"DrivingLicense","first_surname":"Tomas","id":"did:polygonid:polygon:mumbai:2qGg7TzmcoU4Jg3E86wXp4WJcyGUTuafPZxVRxpYQr","license_type":1,"nif":"54688188M","second_surname":"Escoruela"},"expirationDate":"2028-06-15T07:07:39Z","id":"https://api.bloock.dev/identity/v1/did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U/claims/5c9b42c2-13c6-4fcf-b76b-57e104ee8f9c","issuer":"did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U","issuanceDate":"2023-07-24T10:29:25.18351605Z","type":["VerifiableCredential","DrivingLicense"]}`
	signProof := `{"type":"BJJSignature2021","issuerData":{"id":"did:polygonid:polygon:mumbai:2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV","state":{"claimsTreeRoot":"dae18071fcd09e561cc5516beb492a2d7f56e33a77a0de3d7f5dde425637d725","value":"0a8a9ad617e03eccc412fcbb404832ba5eaa877aa3741ab8aa0aac0e7ce90926"},"authCoreClaim":"cca3371a6cb1b715004407e325bd993c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000063d54c7b631861b879bf3bc2e69bd3dc6ee5f86963e0a30a417f978e69b05012bf5aaf8b489949e18dfda8a388f3173058104d9d4c0b797e2b2f93a6492d55150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","mtp":{"existence":true,"siblings":[]},"credentialStatus":{"id":"https://api.bloock.dev/identity/v1/did%3Apolygonid%3Apolygon%3Amumbai%3A2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV/claims/revocation/status/0","revocationNonce":0,"type":"SparseMerkleTreeProof"}},"coreClaim":"e8c51220d38a937afdcf86110d6ccdb92a00000000000000000000000000000002125caf312e33a0b0c82d57fdd240b7261d58901a346261c5ce5621136c0b00811608ce01121947fc78f2f14b4ca45c2c5144df24b353a4d317616393f8452100000000000000000000000000000000000000000000000000000000000000009a9c241d00000000ee30c6f30000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signature":"60621fbf374fdc12a1e1d15f1570cb38cb113a130ef60f67b3176900e2a6dd1bbe1903e1dd6dd71f30ad4d085ab1de88e66d3c2777dab1114eb79eba676f3901"}`

	var credentialData json.RawMessage
	err = json.Unmarshal([]byte(cred), &credentialData)
	require.NoError(t, err)
	var signatureProof json.RawMessage
	err = json.Unmarshal([]byte(signProof), &signatureProof)
	require.NoError(t, err)

	credential := domain.Credential{
		CredentialId:   credentialId,
		HolderDid:      holderDid,
		CredentialType: credentialType,
		IssuerDid:      issuerDid,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
	}

	sparseProof := `{"type":"Iden3SparseMerkleTreeProof","issuerData":{"id":"did:polygonid:polygon:mumbai:2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV","state":{"txId":"0x6b075aa63d20a2be49c99c49fce6f11e38539a57f547f618f24f0f45caaa343d","blockTimestamp":1691672214,"blockNumber":38843289,"rootOfRoots":"3a141c020ba9591ec270dad8f67b3f7d04d14fefd9bce847fd847ebf69765a13","claimsTreeRoot":"c3badcaaafb6f5955bcc9b9c1b73f01c4fc9dc5ad1f0610e2db5aca595eaf20a","revocationTreeRoot":"0000000000000000000000000000000000000000000000000000000000000000","value":"f54348d9e88a380a97d740124e1eda1b1880195ce9e886066a85b10de1257720"}},"coreClaim":"e8c51220d38a937afdcf86110d6ccdb92a00000000000000000000000000000002125caf312e33a0b0c82d57fdd240b7261d58901a346261c5ce5621136c0b00811608ce01121947fc78f2f14b4ca45c2c5144df24b353a4d317616393f8452100000000000000000000000000000000000000000000000000000000000000009a9c241d00000000ee30c6f30000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","mtp":{"existence":true,"siblings":["17115829438154782853639910941115220327911765576806113666055455307308952773082"]}}`
	var sparseMtProof json.RawMessage
	err = json.Unmarshal([]byte(sparseProof), &sparseMtProof)
	require.NoError(t, err)

	credentialUpdated := domain.Credential{
		CredentialId:   credentialId,
		HolderDid:      holderDid,
		CredentialType: credentialType,
		IssuerDid:      issuerDid,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
		SparseMtProof:  sparseMtProof,
	}

	cr := NewSQLCredentialRepository(*conn, 5*time.Second, zerolog.Logger{})

	t.Run("Given a credential it should be saved", func(t *testing.T) {
		err = cr.Save(context.Background(), credential)
		assert.NoError(t, err)

		res, err := cr.GetCredentialById(context.Background(), credentialId)
		assert.NoError(t, err)
		assert.Equal(t, credential.CredentialData, res.CredentialData)
		assert.Equal(t, credential.SignatureProof, res.SignatureProof)
		assert.Equal(t, "null", string(res.SparseMtProof))

		err = cr.UpdateSparseMtProof(context.Background(), credential.CredentialId, sparseMtProof)
		assert.NoError(t, err)

		res, err = cr.GetCredentialById(context.Background(), credentialId)
		assert.NoError(t, err)
		assert.Equal(t, credentialUpdated, res)
	})
}
