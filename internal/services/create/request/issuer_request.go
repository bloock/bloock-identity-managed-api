package request

type CreateIssuerRequest struct {
	DidMetadata DidMetadataRequest
}

type DidMetadataRequest struct {
	Blockchain string
	Method     string
	Network    string
}
