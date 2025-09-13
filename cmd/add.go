/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.GetConfig(cmd)
		if err != nil {
			return fmt.Errorf("failed to get config: %w", err)
		}

		fmt.Println(cfg.Foo)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
