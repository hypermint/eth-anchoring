package cmd

import (
	"context"

	"eth-anchoring/pkg/anchor"
	"eth-anchoring/pkg/chain"
	"eth-anchoring/pkg/logger"
	ptypes "eth-anchoring/pkg/types"
	"eth-anchoring/pkg/wallet"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func MakeRunCMD() *cobra.Command {
	const (
		flagL1Address = "l1.address"
		flagL2Address = "l2.address"
		flagMnemonic  = "mnemonic"
		flagHDWPath   = "hdw-path"
	)
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "run an anchoring service",
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
			cc, err := MakeContractClient(
				viper.GetString(flagMnemonic),
				viper.GetString(flagHDWPath),
				viper.GetString(flagL1Address),
			)
			l2 := chain.NewL2Client(viper.GetString(flagL2Address))

			ctx, cancel := context.WithCancel(context.Background())
			srv := anchor.NewAnchoringService(logger.With("module", "anchoring_service"), l2, cc, ptypes.NewSigner(prv))
			go func() {
				defer cancel()
				if err := srv.Start(ctx); err != nil {
					logger.Error("failed to start", "err", err)
				}
			}()
			select {
			case <-ctx.Done():
				return nil
			}
		},
	}
	runCmd.Flags().String(flagL1Address, "http://localhost:8545", "an address for l1 node")
	runCmd.Flags().String(flagL2Address, "tcp://127.0.0.1:26657", "an address for l2 node")
	runCmd.Flags().String(flagMnemonic, "", "mnemonic string")
	runCmd.Flags().String(flagHDWPath, "", "HD Wallet path")
	return runCmd
}
