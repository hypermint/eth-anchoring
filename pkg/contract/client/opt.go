package client

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"eth-anchoring/pkg/consts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gcrypto "github.com/ethereum/go-ethereum/crypto"
)

func MakeGenTxOpts(prv *ecdsa.PrivateKey) GenTxOpts {
	addr := gcrypto.PubkeyToAddress(prv.PublicKey)
	return func(ctx context.Context) *bind.TransactOpts {
		return &bind.TransactOpts{
			From:     addr,
			GasLimit: consts.DefaultGasLimit,
			Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				if address != addr {
					return nil, errors.New("not authorized to sign this account")
				}
				signature, err := gcrypto.Sign(signer.Hash(tx).Bytes(), prv)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
		}
	}
}
