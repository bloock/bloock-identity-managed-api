package create

import (
	"bloock-identity-managed-api/internal/domain"
	"bloock-identity-managed-api/internal/domain/repository"
	"bloock-identity-managed-api/internal/services/create/request"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Credential struct {
	credentialRepository repository.CredentialRepository
	identityRepository   repository.IdentityRepository
	logger               zerolog.Logger
}

func NewCredential(cr repository.CredentialRepository, ir repository.IdentityRepository, l zerolog.Logger) *Credential {
	return &Credential{
		credentialRepository: cr,
		identityRepository:   ir,
		logger:               l,
	}
}

func (c Credential) Create(ctx context.Context, req request.CredentialRequest) (interface{}, error) {
	credentialId, _ := uuid.NewUUID()

	schemaType := "DrivingLicense"
	issuerDid := "did:polygonid:polygon:mumbai:2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV"

	holderDid := "did:polygonid:polygon:mumbai:2qGg7TzmcoU4Jg3E86wXp4WJcyGUTuafPZxVRxpYQr"
	cred := `{"@context":["https://www.w3.org/2018/credentials/v1","https://schema.iden3.io/core/jsonld/iden3proofs.jsonld","https://api.bloock.dev/hosting/v1/ipfs/QmZ9BzmMGzLv4y9P6djYUm8sgQt47ZjECGAM8nToFW2qvt"],"credentialSchema":{"id":"https://api.bloock.dev/hosting/v1/ipfs/QmTvHzXiegijCdhGC4kgjps8hSi3FP1K17ezrYPgdMU6Ek","type":"JsonSchema2023"},"credentialStatus":{"id":"https://api.bloock.dev/identity/v1/did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U/claims/revocation/status/3825417065","revocationNonce":3825417065,"type":"SparseMerkleTreeProof"},"credentialSubject":{"birth_date":921950325,"country":"Spain","name":"Eduard","type":"DrivingLicense","first_surname":"Tomas","id":"did:polygonid:polygon:mumbai:2qGg7TzmcoU4Jg3E86wXp4WJcyGUTuafPZxVRxpYQr","license_type":1,"nif":"54688188M","second_surname":"Escoruela"},"expirationDate":"2028-06-15T07:07:39Z","id":"https://api.bloock.dev/identity/v1/did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U/claims/5c9b42c2-13c6-4fcf-b76b-57e104ee8f9c","issuer":"did:polygonid:polygon:mumbai:2qGUovMWDMyoXKLWiRMBRTyMfKcdrUg958QcCDkC9U","issuanceDate":"2023-07-24T10:29:25.18351605Z","type":["VerifiableCredential","DrivingLicense"]}`
	signProof := `{"type":"BJJSignature2021","issuerData":{"id":"did:polygonid:polygon:mumbai:2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV","state":{"claimsTreeRoot":"dae18071fcd09e561cc5516beb492a2d7f56e33a77a0de3d7f5dde425637d725","value":"0a8a9ad617e03eccc412fcbb404832ba5eaa877aa3741ab8aa0aac0e7ce90926"},"authCoreClaim":"cca3371a6cb1b715004407e325bd993c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000063d54c7b631861b879bf3bc2e69bd3dc6ee5f86963e0a30a417f978e69b05012bf5aaf8b489949e18dfda8a388f3173058104d9d4c0b797e2b2f93a6492d55150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","mtp":{"existence":true,"siblings":[]},"credentialStatus":{"id":"https://api.bloock.dev/identity/v1/did%3Apolygonid%3Apolygon%3Amumbai%3A2qNeWsY3DuGhARS3mXsHYw3tmB48tUoVq7pSR5jBmV/claims/revocation/status/0","revocationNonce":0,"type":"SparseMerkleTreeProof"}},"coreClaim":"e8c51220d38a937afdcf86110d6ccdb92a00000000000000000000000000000002125caf312e33a0b0c82d57fdd240b7261d58901a346261c5ce5621136c0b00811608ce01121947fc78f2f14b4ca45c2c5144df24b353a4d317616393f8452100000000000000000000000000000000000000000000000000000000000000009a9c241d00000000ee30c6f30000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","signature":"60621fbf374fdc12a1e1d15f1570cb38cb113a130ef60f67b3176900e2a6dd1bbe1903e1dd6dd71f30ad4d085ab1de88e66d3c2777dab1114eb79eba676f3901"}`
	iProof := `{"leaves":["232e8173357afd71a8d02928ac91a3b35dd91b942364cccd155fc569a5e6345e"],"nodes":["3e8a823763984f5466236ddd3e65d792d1182c38a685e080911b67092ac33921","e3f9c004f09f91a67964895451d4e56e51d07da534af0baa91648f667efecd8d","ea7ef1e3c8d78165135796d5cf20797131626b67732964b79af889befa2ff16e","c1f286de5d9ef9a2647f1439c70a1b81402f3a09c35f87e86e39326a2b948066","5bfd064435e9f185e60fca87354d11d2e48a55927e3aedea7bacc1a9790615dc","74cc77f8375511298a8b3f830f4dcba0e0377cbbfecab2137a2c1bd7d583a423","edd3f4d85b41c7c3afbe5b53e647680d230514f1d58aa9a765137c81906d73fb","eecba25253ce73acb03823cd9b3c44b5e0f3a01b1b984755e5d97af00d2a8f80","e6c4c0978941bae8d1751bb1b05743422043a7b86d6a11e6e467b6dcf1598886"],"depth":"0001000200030005000600090009000800070004","bitmap":"fbc0","anchor":{"anchor_id":117969,"networks":[{"name":"ethereum_goerli","state":"Pending","tx_hash":"0x2dd3c8a9dd2957691389d96841eba54f534c6c99d52e9c9cc26ce8e2633dd7de"},{"name":"bloock_chain","state":"Confirmed","tx_hash":"0x78225aa21a2092d1d0ee5ef86f36b3474ed3f75836dd142acff5a3f24f4be589"}],"root":"014fd762294735c76511ad8b95be8248c5ed0c5189892e547fb71512ce30a9c1","status":"Processing"},"type":"BloockIntegrityProof"}`
	proofTypes := []string{"bloock_proof", "polygon_proof"}

	for _, p := range proofTypes {
		_, err := domain.NewProofType(p)
		if err != nil {
			c.logger.Error().Err(err).Msg("")
			return nil, err
		}
	}

	var credentialData json.RawMessage
	_ = json.Unmarshal([]byte(cred), &credentialData)
	var signatureProof json.RawMessage
	_ = json.Unmarshal([]byte(signProof), &signatureProof)
	var integrityProof json.RawMessage
	_ = json.Unmarshal([]byte(iProof), &integrityProof)

	credential := domain.Credential{
		CredentialId:   credentialId,
		IssuerDid:      issuerDid,
		HolderDid:      holderDid,
		SchemaType:     schemaType,
		ProofType:      proofTypes,
		CredentialData: credentialData,
		SignatureProof: signatureProof,
		IntegrityProof: integrityProof,
	}

	if err := c.credentialRepository.Save(ctx, credential); err != nil {
		return nil, err
	}

	return credentialId.String(), nil
}
