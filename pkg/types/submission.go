package types

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tendermint/tendermint/types"
)

type Submission struct {
	Height       [8]byte
	BlockHeaders []SubmissionBlockHeader
	Signature    []byte
}

func (s Submission) GetHeight() uint64 {
	return binary.BigEndian.Uint64(s.Height[:])
}

func (s Submission) Bytes() []byte {
	b, err := rlp.EncodeToBytes(s)
	if err != nil {
		panic(err)
	}
	return b
}

func (s *Submission) FromBytes(b []byte) error {
	return rlp.DecodeBytes(b, s)
}

func (s *Submission) HeadersHash() []byte {
	var hash []byte
	for _, h := range s.BlockHeaders {
		if hash == nil {
			hash = h.MerkleRootHash.Bytes()
		} else {
			harray := sha256.Sum256(
				append(hash, h.MerkleRootHash.Bytes()...),
			)
			hash = harray[:]
		}
	}
	return hash
}

func (s *Submission) GetMessageHash() (common.Hash, error) {
	var bytesToSign = []byte("\x19Ethereum Signed Message:\n32")
	hash := s.HeadersHash()
	return crypto.Keccak256Hash(append(bytesToSign, hash...)), nil
}

func (s *Submission) SignWith(f func([]byte) ([]byte, error)) error {
	h, err := s.GetMessageHash()
	if err != nil {
		return err
	}
	sig, err := f(h.Bytes())
	if err != nil {
		return err
	}
	s.Signature = sig
	return nil
}

func (s Submission) Validate() error {
	if s.Height == [8]byte{} {
		return errors.New("height must not be empty")
	}
	if len(s.Signature) != 65 {
		return fmt.Errorf("length of signatures must be 65, but got %v", len(s.Signature))
	}
	if len(s.BlockHeaders) == 0 {
		return errors.New("blockHeaders must not be empty")
	}
	return nil
}

type SubmissionBlockHeader struct {
	MerkleRootHash common.Hash
}

func MakeSubmission(height uint64, blocks []*types.Block) (*Submission, error) {
	var hashes []common.Hash
	for _, b := range blocks {
		hashes = append(hashes, common.BytesToHash(b.Hash()))
	}
	return MakeSubmissionWithHashes(height, hashes)
}

func MakeSubmissionWithHashes(height uint64, hashes []common.Hash) (*Submission, error) {
	s := new(Submission)
	binary.BigEndian.PutUint64(s.Height[:], height)
	for _, h := range hashes {
		s.BlockHeaders = append(
			s.BlockHeaders, SubmissionBlockHeader{
				MerkleRootHash: h,
			},
		)
	}
	return s, nil
}
