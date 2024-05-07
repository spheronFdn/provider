package client

import (
	"errors"

	"github.com/akash-network/provider/spheron"
)

var (
	ErrInvalidClient = errors.New("invalid client")
)

func DiscoverClient() *spheron.Client {
	return spheron.NewClient()
}
