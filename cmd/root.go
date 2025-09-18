/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/ethn1ee/llog/internal/logger"
	"github.com/spf13/cobra"
)

var (
	cfg = &config.Config{}
	db  = &_db.DB{}
	lg  = &logger.Logger{}
)

var rootCmd = &cobra.Command{
	Use:               "llog",
	Short:             "Life log",
	Long:              `Record your fleeting moments with llog.`,
	PersistentPreRunE: handler.Init(cfg, db, lg),
	SilenceUsage:      true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	_ = lg.Close()
}
