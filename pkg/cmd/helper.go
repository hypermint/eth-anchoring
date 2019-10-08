package cmd

import (
	"eth-anchoring/pkg/chain"
	"eth-anchoring/pkg/contract/client"
	"eth-anchoring/pkg/wallet"
)

func MakeContractClient(mnemonic, hdwpath, l1addr string) (client.Client, error) {
	prv, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(
		mnemonic,
		hdwpath,
	)
	if err != nil {
		return nil, err
	}
	l1, err := chain.NewL1Client(l1addr)
	if err != nil {
		return nil, err
	}
	return client.NewClient(l1, client.MakeGenTxOpts(prv))
}
