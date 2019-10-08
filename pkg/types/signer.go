package types

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Signer interface {
	Sign([]byte) ([]byte, error)
}

type signer struct {
	prv *ecdsa.PrivateKey
}

func (s *signer) Sign(bs []byte) ([]byte, error) {
	return crypto.Sign(bs, s.prv)
}

func NewSigner(prv *ecdsa.PrivateKey) Signer {
	return &signer{prv: prv}
}
