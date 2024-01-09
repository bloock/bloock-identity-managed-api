package request

type CreateIssuerRequest struct {
	Key             string
	DidMetadata     DidMetadataRequest
	Name            string
	Description     string
	Image           string
	PublishInterval int
}

type DidMetadataRequest struct {
	Blockchain string
	Method     string
	Network    string
}
