package chain

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type L1Client struct {
	endpoint string

	conn *rpc.Client
	*ethclient.Client
}

func NewL1Client(endpoint string) (*L1Client, error) {
	conn, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return &L1Client{
		endpoint: endpoint,
		conn:     conn,
		Client:   ethclient.NewClient(conn),
	}, nil
}
