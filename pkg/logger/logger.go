package logger

import (
	"os"

	"github.com/spf13/viper"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
)

type Logger = log.Logger

func GetLogger(lv string) Logger {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	logger, err := tmflags.ParseLogLevel(lv, logger, cfg.DefaultLogLevel())
	if err != nil {
		panic(err)
	}
	if viper.GetBool(cli.TraceFlag) {
		logger = log.NewTracingLogger(logger)
	}
	return logger.With("module", "main")
}
