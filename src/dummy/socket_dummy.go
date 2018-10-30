package dummy

import (
	socket "github.com/andrecronje/lachesis/src/proxy"
	"github.com/sirupsen/logrus"
)

// DummySocketClient is a socket implementation of the dummy app. Lachesis and the
// app run in separate processes and communicate through TCP sockets using
// a SocketLachesisProxy and a SocketAppProxy.
type DummySocketClient struct {
	state         *State
	lachesisProxy *socket.GrpcLachesisProxy
	logger        *logrus.Logger
}

// NewDummySocketClient instantiates a DummySocketClient and starts the
// SocketLachesisProxy
func NewDummySocketClient(clientAddr string, nodeAddr string, logger *logrus.Logger) (*DummySocketClient, error) {
	state := NewState(logger)

	lachesisProxy, err := socket.NewGrpcLachesisProxy(nodeAddr, logger)
	if err != nil {
		return nil, err
	}
	client := &DummySocketClient{
		state:         state,
		lachesisProxy: lachesisProxy,
		logger:        logger,
	}
	return client, nil
}

// SubmitTx sends a transaction to Babble via the SocketProxy
func (c *DummySocketClient) SubmitTx(tx []byte) error {
	return c.lachesisProxy.SubmitTx(tx)
}
