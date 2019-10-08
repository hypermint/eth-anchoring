package chain

import (
	"github.com/tendermint/tendermint/rpc/client"
)

type L2Client struct {
	client.Client
}

func NewL2Client(address string) *L2Client {
	rpc := client.NewHTTP(address, "/websocket")
	return &L2Client{rpc}
}

func (cl *L2Client) GetLatestBlockHeight() (int64, error) {
	res, err := cl.BlockchainInfo(0, 0)
	if err != nil {
		return 0, err
	}
	return res.LastHeight, nil
}
