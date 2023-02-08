package proxy

import (
	"github.com/pkg/errors"

	"github.com/zhengjianfeng1103/fbc/libs/tendermint/libs/service"
)

//-----------------------------

// Tendermint's interface to the application consists of multiple connections
type AppConns interface {
	service.Service

	Mempool() AppConnMempool
	Consensus() AppConnConsensus
	Query() AppConnQuery
}

func NewAppConns(clientCreator ClientCreator) AppConns {
	return NewMultiAppConn(clientCreator)
}

//-----------------------------
// multiAppConn implements AppConns

// a multiAppConn is made of a few appConns (mempool, consensus, query)
// and manages their underlying abci clients
// TODO: on app restart, clients must reboot together
type multiAppConn struct {
	service.BaseService

	mempoolConn   AppConnMempool
	consensusConn AppConnConsensus
	queryConn     AppConnQuery

	clientCreator ClientCreator
}

// Make all necessary abci connections to the application
func NewMultiAppConn(clientCreator ClientCreator) AppConns {
	multiAppConn := &multiAppConn{
		clientCreator: clientCreator,
	}
	multiAppConn.BaseService = *service.NewBaseService(nil, "multiAppConn", multiAppConn)
	return multiAppConn
}

// Returns the mempool connection
func (app *multiAppConn) Mempool() AppConnMempool {
	return app.mempoolConn
}

// Returns the consensus Connection
func (app *multiAppConn) Consensus() AppConnConsensus {
	return app.consensusConn
}

// Returns the query Connection
func (app *multiAppConn) Query() AppConnQuery {
	return app.queryConn
}

func (app *multiAppConn) OnStart() error {
	// query connection
	querycli, err := app.clientCreator.NewABCIClient()
	if err != nil {
		return errors.Wrap(err, "Error creating ABCI client (query connection)")
	}
	querycli.SetLogger(app.Logger.With("module", "abci-client", "connection", "query"))
	if err := querycli.Start(); err != nil {
		return errors.Wrap(err, "Error starting ABCI client (query connection)")
	}
	app.queryConn = NewAppConnQuery(querycli)

	// mempool connection
	memcli, err := app.clientCreator.NewABCIClient()
	if err != nil {
		return errors.Wrap(err, "Error creating ABCI client (mempool connection)")
	}
	memcli.SetLogger(app.Logger.With("module", "abci-client", "connection", "mempool"))
	if err := memcli.Start(); err != nil {
		return errors.Wrap(err, "Error starting ABCI client (mempool connection)")
	}
	app.mempoolConn = NewAppConnMempool(memcli)

	// consensus connection
	concli, err := app.clientCreator.NewABCIClient()
	if err != nil {
		return errors.Wrap(err, "Error creating ABCI client (consensus connection)")
	}
	concli.SetLogger(app.Logger.With("module", "abci-client", "connection", "consensus"))
	if err := concli.Start(); err != nil {
		return errors.Wrap(err, "Error starting ABCI client (consensus connection)")
	}
	app.consensusConn = NewAppConnConsensus(concli)

	return nil
}
