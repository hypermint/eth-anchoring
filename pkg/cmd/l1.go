package cmd

import (
	"bytes"
	"context"
	"eth-anchoring/pkg/chain"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func MakeL1CMD() *cobra.Command {
	const (
		flagL1Address = "l1.address"
		flagL2Address = "l2.address"

		flagMnemonic = "mnemonic"
		flagHDWPath  = "hdw-path"
	)

	var l1Cmd = &cobra.Command{
		Use:   "l1",
		Short: "commands for l1",
	}

	const (
		flagHeight = "height"
	)

	var latestCmd = &cobra.Command{
		Use:   "latest",
		Short: "show latest height",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.BindPFlags(cmd.Flags())
			cc, err := MakeContractClient(viper.GetString(flagMnemonic), viper.GetString(flagHDWPath), viper.GetString(flagL1Address))
			if err != nil {
				return err
			}
			ctx := context.Background()
			h, err := cc.GetLastBlockNumber(ctx)
			if err != nil {
				return err
			}
			fmt.Println(h)
			return nil
		},
	}

	var blockCmd = &cobra.Command{
		Use:   "block",
		Short: "show block hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.BindPFlags(cmd.Flags())
			cc, err := MakeContractClient(viper.GetString(flagMnemonic), viper.GetString(flagHDWPath), viper.GetString(flagL1Address))
			if err != nil {
				return err
			}
			ctx := context.Background()
			height := viper.GetUint64(flagHeight)
			if height == 0 {
				return fmt.Errorf("height must be greater than 0")
			}
			h, err := cc.GetHeaderHash(ctx, height)
			if err != nil {
				return err
			}
			fmt.Println(h.Hex())
			return nil
		},
	}
	blockCmd.Flags().Uint64(flagHeight, 0, "")

	var verifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "check if l1 has a valid block hash",
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.BindPFlags(cmd.Flags())
			cc, err := MakeContractClient(viper.GetString(flagMnemonic), viper.GetString(flagHDWPath), viper.GetString(flagL1Address))
			if err != nil {
				return err
			}
			l2 := chain.NewL2Client(viper.GetString(flagL2Address))
			ctx := context.Background()
			height := viper.GetUint64(flagHeight)
			if height == 0 {
				return fmt.Errorf("height must be greater than 0")
			}
			h, err := cc.GetHeaderHash(ctx, height)
			if err != nil {
				return err
			}
			ih := int64(height)
			blk, err := l2.Block(&ih)
			if err != nil {
				return err
			}
			if !bytes.Equal(h.Bytes(), blk.Block.Hash().Bytes()) {
				return fmt.Errorf("%X != %X", h.Bytes(), blk.Block.Hash().Bytes())
			}
			fmt.Println("ok")
			return nil
		},
	}
	verifyCmd.Flags().Uint64(flagHeight, 0, "")

	l1Cmd.AddCommand(latestCmd, blockCmd, verifyCmd)
	l1Cmd.PersistentFlags().String(flagL1Address, "http://localhost:8545", "an address for l1 node")
	l1Cmd.PersistentFlags().String(flagL2Address, "tcp://127.0.0.1:26657", "an address for l2 node")
	l1Cmd.PersistentFlags().String(flagMnemonic, "", "mnemonic string")
	l1Cmd.PersistentFlags().String(flagHDWPath, "", "HD Wallet path")

	return l1Cmd
}
