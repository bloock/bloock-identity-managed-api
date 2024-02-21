package request

type CreateIssuerRequest struct {
	Key             string
	Name            string
	Description     string
	Image           string
	PublishInterval int
}
