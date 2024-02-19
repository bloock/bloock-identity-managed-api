package domain

import (
	"errors"
	identityEntity "github.com/bloock/bloock-sdk-go/v2/entity/identity"
)

type PublishIntervalMinutes int

var ErrInvalidPublishIntervalMinutes = errors.New("publish interval minutes not supported")

const (
	PublishIntervalMinutes5 PublishIntervalMinutes = iota
	PublishIntervalMinutes15
	PublishIntervalMinutes60
)

func NewPublishIntervalMinutes(_type int) (PublishIntervalMinutes, error) {
	switch _type {
	case 5:
		return PublishIntervalMinutes5, nil
	case 15:
		return PublishIntervalMinutes15, nil
	case 60:
		return PublishIntervalMinutes60, nil
	default:
		return -1, ErrInvalidPublishIntervalMinutes
	}
}

func (p PublishIntervalMinutes) Params() identityEntity.PublishIntervalParams {
	switch p {
	case PublishIntervalMinutes5:
		return identityEntity.Interval5
	case PublishIntervalMinutes15:
		return identityEntity.Interval15
	case PublishIntervalMinutes60:
		return identityEntity.Interval60
	default:
		return identityEntity.Interval60
	}
}
