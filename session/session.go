package session

import (
	"github.com/tendermint/tendermint/libs/log"

	ptypes "github.com/akash-network/akash-api/go/node/provider/v1beta3"
	"github.com/akash-network/provider/spheron"
)

// Session interface wraps Log, Client, Provider and ForModule methods
type Session interface {
	Log() log.Logger
	Client() *spheron.Client
	Provider() *ptypes.Provider
	ForModule(string) Session
	CreatedAtBlockHeight() int64
}

func New(log log.Logger, provider *ptypes.Provider, sphClient *spheron.Client, createdAtBlockHeight int64) Session {
	return session{
		provider:             provider,
		log:                  log,
		client:               sphClient,
		createdAtBlockHeight: createdAtBlockHeight,
	}
}

type session struct {
	client               *spheron.Client
	provider             *ptypes.Provider
	log                  log.Logger
	createdAtBlockHeight int64
}

func (s session) Log() log.Logger {
	return s.log
}

func (s session) Client() *spheron.Client {
	return s.client
}

func (s session) Provider() *ptypes.Provider {
	return s.provider
}

func (s session) ForModule(name string) Session {
	s.log = s.log.With("module", name)
	return s
}

func (s session) CreatedAtBlockHeight() int64 {
	return s.createdAtBlockHeight
}
