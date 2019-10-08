package anchor

import (
	"context"
	"errors"
	"time"

	"eth-anchoring/pkg/chain"
	"eth-anchoring/pkg/contract/client"
	"eth-anchoring/pkg/logger"
	"eth-anchoring/pkg/types"

	tmtypes "github.com/tendermint/tendermint/types"
)

type AnchoringService struct {
	logger logger.Logger
	l2     *chain.L2Client
	cc     client.Client
	signer types.Signer

	option
}

func NewAnchoringService(
	logger logger.Logger,
	l2 *chain.L2Client,
	cc client.Client,
	signer types.Signer,
	opts ...Option,
) *AnchoringService {
	option := defaultOption()
	for _, opt := range opts {
		opt(&option)
	}
	return &AnchoringService{logger: logger, l2: l2, cc: cc, signer: signer, option: option}
}

// DoOneShot performs a one-shot block submission
func (a *AnchoringService) DoOneShot(ctx context.Context, height int64) error {
	blocks, err := a.getBlocks(height, 1)
	if err != nil {
		return err
	}
	state := anchorState{
		nextHeight: height,
		blocks:     blocks,
	}
	if err := a.enterSubmit(ctx, &state); err != nil {
		return err
	}
	return nil
}

// Start starts a anchoring service. This service observe block store at L2, and submit its updates to l1 contract
/*
AnchoringService follows bellow steps:
1. get a height of last submitted block from contract state
2. lookup blocks from l2 store
3. submit blocks to l1 contract

Important: This service doesn't take into account that reorg occurs at ethereum
*/
func (a *AnchoringService) Start(ctx context.Context) error {
	state, err := a.getState(ctx)
	if err != nil {
		return err
	}
	return a.loop(ctx, state)
}

func (a *AnchoringService) loop(ctx context.Context, initialState anchorState) error {
	state := initialState
	a.logger.Info("start to loop", "next", state.nextHeight)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if err := a.enterCollectBlocks(ctx, &state); err != nil {
				return err
			}
			if err := a.enterSubmit(ctx, &state); err != nil {
				return err
			}
			state = updateState(state)
			time.Sleep(5 * time.Second)
		}
	}
}

func (a *AnchoringService) enterCollectBlocks(ctx context.Context, state *anchorState) error {
	a.logger.Info("start to enterCollectBlocks", "next", state.nextHeight)
	for {
		latest, err := a.l2.GetLatestBlockHeight()
		if err != nil {
			a.logger.Error("failed to get a latest block height", "err", err.Error())
			continue
		}
		if state.nextHeight >= latest {
			a.logger.Info("wait until new block is created", "next", state.nextHeight, "latest", latest)
			time.Sleep(10 * time.Second)
			continue
		}

		bn := decideBlocksNumber(state.nextHeight, latest, a.option.maxSize)
		if bn <= 0 {
			panic("decideBlocksNumber: unexpected inputs")
		}
		blocks, err := a.getBlocks(state.nextHeight, bn)
		if err != nil {
			a.logger.Error("failed to get blocks", "err", err.Error())
			continue
		}
		state.blocks = blocks
		return nil
	}
}

func decideBlocksNumber(next, latest, max int64) int64 {
	n := latest - next
	if n >= max {
		return max
	} else {
		return n
	}
}

func (a *AnchoringService) getBlocks(next, num int64) ([]*tmtypes.Block, error) {
	blocks := make([]*tmtypes.Block, 0, num)

	for h := next; h < next+num; h++ {
		blk, err := a.l2.Block(&h)
		if err != nil {
			a.logger.Error("failed to get block", "height", h)
			return nil, err
		}
		blocks = append(blocks, blk.Block)
	}

	return blocks, nil
}

func (a *AnchoringService) enterSubmit(ctx context.Context, state *anchorState) error {
	a.logger.Info("start to enterSubmit", "next", state.nextHeight, "num", len(state.blocks))
	s, err := types.MakeSubmission(uint64(state.nextHeight), state.blocks)
	if err != nil {
		return err
	}
	if err := s.SignWith(a.signer.Sign); err != nil {
		return err
	}
	if err := s.Validate(); err != nil {
		return err
	}
	return a.doSubmit(ctx, s)
}

func (a *AnchoringService) doSubmit(ctx context.Context, submission *types.Submission) error {
	a.logger.Info("start to doSubmit", "height", submission.GetHeight(), "num", len(submission.BlockHeaders))
	tx, err := a.cc.SubmitBlockHeaders(ctx, submission.Bytes())
	if err != nil {
		return err
	}
	if tx != nil {
		a.logger.Info("submission is completed", "hash", tx.Hash().Hex(), "start", submission.GetHeight(), "num", len(submission.BlockHeaders))
	}
	if err != nil {
		return err
	}
	if tx == nil {
		return errors.New("contract execution was reverted?")
	}
	return nil
}

func (a *AnchoringService) getState(ctx context.Context) (anchorState, error) {
	h, err := a.cc.GetLastBlockNumber(ctx)
	if err != nil {
		return anchorState{}, err
	}
	return anchorState{nextHeight: int64(h) + 1}, nil
}

func updateState(state anchorState) anchorState {
	return anchorState{
		nextHeight: state.nextHeight + int64(len(state.blocks)),
	}
}

type anchorState struct {
	nextHeight int64
	blocks     []*tmtypes.Block
}

/// Option definition ///

type Option func(*option)

type option struct {
	maxSize       int64         // maxSize is max number of blocks length per one submission.
	submitTimeout time.Duration // submitTimeout is a timeout for calling SubmitBlockHeaders on doSubmit().
}

func defaultOption() option {
	return option{
		maxSize:       5,
		submitTimeout: time.Minute,
	}
}

func WithMaxSize(s int64) Option {
	return func(opt *option) {
		opt.maxSize = s
	}
}
