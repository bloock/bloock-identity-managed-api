package domain

import (
	"errors"
	"github.com/bloock/bloock-sdk-go/v2/entity/identityV2"
)

type PublishIntervalMinutes int

var ErrInvalidPublishIntervalMinutes = errors.New("publish interval minutes not supported")

const (
	PublishIntervalMinutes1 PublishIntervalMinutes = iota
	PublishIntervalMinutes5
	PublishIntervalMinutes15
	PublishIntervalMinutes60
)

func NewPublishIntervalMinutes(_type int) (PublishIntervalMinutes, error) {
	switch _type {
	case 1:
		return PublishIntervalMinutes1, nil
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

func (p PublishIntervalMinutes) Params() identityV2.PublishIntervalParams {
	switch p {
	case PublishIntervalMinutes1:
		return identityV2.Interval1
	case PublishIntervalMinutes5:
		return identityV2.Interval5
	case PublishIntervalMinutes15:
		return identityV2.Interval15
	case PublishIntervalMinutes60:
		return identityV2.Interval60
	default:
		return identityV2.Interval60
	}
}
