package response

type GetCredentialOfferResponse struct {
	ID       string
	ThreadID string
	Body     GetCredentialOfferBodyResponse
	From     string
	To       string
	Typ      string
	Type     string
}

type GetCredentialOfferBodyResponse struct {
	URL         string
	ID          string
	Description string
}
