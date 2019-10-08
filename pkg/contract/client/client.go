package client

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

	"eth-anchoring/pkg/chain"
	"eth-anchoring/pkg/consts"
	"eth-anchoring/pkg/contract/blocks"
)

type Client interface {
	SubmitBlockHeaders(ctx context.Context, submissionBytes []byte) (*etypes.Transaction, error)
	GetOperator(ctx context.Context) (common.Address, error)
	GetLastBlockNumber(ctx context.Context) (uint64, error)
	GetHeaderHash(ctx context.Context, height uint64) (common.Hash, error)
}

type client struct {
	genTxOpts GenTxOpts
	ethc      *chain.L1Client
	contract  *blocks.Blocks
}

func (cl client) SubmitBlockHeaders(ctx context.Context, submissionBytes []byte) (*etypes.Transaction, error) {
	return cl.contract.Submit(cl.genTxOpts(ctx), submissionBytes)
}

func (cl client) GetOperator(ctx context.Context) (common.Address, error) {
	opts := cl.genTxOpts(ctx)
	return cl.contract.GetOperator(
		&bind.CallOpts{
			From:    opts.From,
			Context: opts.Context,
		},
	)
}

func (cl client) GetLastBlockNumber(ctx context.Context) (uint64, error) {
	opts := cl.genTxOpts(ctx)
	return cl.contract.GetLastBlockNumber(
		&bind.CallOpts{
			From:    opts.From,
			Context: opts.Context,
		},
	)
}

func (cl client) GetHeaderHash(ctx context.Context, height uint64) (common.Hash, error) {
	opts := cl.genTxOpts(ctx)
	h, err := cl.contract.GetHeaderHash(
		&bind.CallOpts{
			From:    opts.From,
			Context: opts.Context,
		}, height,
	)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(h[:]), nil
}

type GenTxOpts func(ctx context.Context) *bind.TransactOpts

func NewClient(ethc *chain.L1Client, genTxOpts GenTxOpts) (Client, error) {
	contract, err := blocks.NewBlocks(common.HexToAddress(consts.BlocksAddress), ethc)
	if err != nil {
		return nil, err
	}
	return &client{ethc: ethc, contract: contract, genTxOpts: genTxOpts}, nil
}
