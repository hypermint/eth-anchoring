package cmd

import (
	"context"
	"eth-anchoring/pkg/anchor"
	"eth-anchoring/pkg/chain"
	"eth-anchoring/pkg/logger"
	ptypes "eth-anchoring/pkg/types"
	"eth-anchoring/pkg/wallet"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func MakeSubmitCMD() *cobra.Command {
	const (
		flagL1Address = "l1.address"
		flagL2Address = "l2.address"
		flagMnemonic  = "mnemonic"
		flagHDWPath   = "hdw-path"

		flagHeight = "height"
		flagHashes = "hashes"
	)

	var submitCmd = &cobra.Command{
		Use:   "submit",
		Short: "submit a l2 block to l1",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.BindPFlags(cmd.Flags())
			logger := logger.GetLogger("*:debug")
			prv, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(
				viper.GetString(flagMnemonic),
				viper.GetString(flagHDWPath),
			)
			if err != nil {
				return err
			}
			height := viper.GetInt64(flagHeight)
			if height == 0 {
				return fmt.Errorf("height must be greater than 0")
			}
			cc, err := MakeContractClient(
				viper.GetString(flagMnemonic),
				viper.GetString(flagHDWPath),
				viper.GetString(flagL1Address),
			)
			if err != nil {
				return err
			}
			l2 := chain.NewL2Client(viper.GetString(flagL2Address))

			ctx := context.Background()
			srv := anchor.NewAnchoringService(logger.With("module", "anchoring_service"), l2, cc, ptypes.NewSigner(prv))

			if hashes := viper.GetStringSlice(flagHashes); len(hashes) == 0 {
				return srv.DoAnchorWithHeight(ctx, height)
			} else {
				return srv.DoAnchorWithHashes(ctx, height, hexStrSliceToHashes(hashes))
			}
		},
	}

	submitCmd.Flags().String(flagL1Address, "http://localhost:8545", "an address for l1 node")
	submitCmd.Flags().String(flagL2Address, "tcp://127.0.0.1:26657", "an address for l2 node")
	submitCmd.Flags().String(flagMnemonic, "", "mnemonic string")
	submitCmd.Flags().String(flagHDWPath, "", "HD Wallet path")

	submitCmd.Flags().Uint64(flagHeight, 0, "block height")
	submitCmd.Flags().StringSlice(flagHashes, nil, "array of block hash")
	return submitCmd
}

func hexStrSliceToHashes(strs []string) []common.Hash {
	var hs []common.Hash
	for _, s := range strs {
		hs = append(hs, common.HexToHash(s))
	}
	return hs
}
