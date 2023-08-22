package local

import "github.com/rs/zerolog"

type LocalKeyProvider struct {
	logger zerolog.Logger
}

func NewLocalKeyProvider(l zerolog.Logger) LocalKeyProvider {
	return LocalKeyProvider{
		logger: l,
	}
}

func (l LocalKeyProvider) Create() {

}
