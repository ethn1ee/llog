/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/ethn1ee/llog/internal/log"

	"github.com/spf13/cobra"
)

type runFn func(cmd *cobra.Command, args []string) error

var rootCmd = &cobra.Command{
	Use:               "llog",
	Short:             "Life log",
	Long:              `Record your fleeting moments with llog.`,
	PersistentPreRunE: setUp,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func withLog(fn runFn) runFn {
	return func(cmd *cobra.Command, args []string) error {
		cmdAttr := slog.String("command", cmd.Name())
		slog.Info("command executing", cmdAttr)
		if err := fn(cmd, args); err != nil {
			slog.Error("command failed", cmdAttr, slog.Any("error", err))
			return err
		}
		slog.Info("command succeeded", cmdAttr)
		return nil
	}
}

func setUp(cmd *cobra.Command, args []string) error {
	if err := config.Init(cmd); err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	cfg, err := config.FromCmd(cmd)
	if err != nil {
		return fmt.Errorf("failed to get config from context: %w", err)
	}

	if err := log.Init(cfg); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	return nil
}
