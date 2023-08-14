package response

type RedeemCredentialResponse struct {
	ID       string
	ThreadID string
	Body     map[string]interface{}
	From     string
	To       string
	Typ      string
	Type     string
}
