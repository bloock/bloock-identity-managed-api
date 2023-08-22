package managed

import "github.com/rs/zerolog"

type ManagedKeyProvider struct {
	logger zerolog.Logger
}

func NewManagedKeyProvider(l zerolog.Logger) ManagedKeyProvider {
	return ManagedKeyProvider{
		logger: l,
	}
}

func (m ManagedKeyProvider) Create() {

}
