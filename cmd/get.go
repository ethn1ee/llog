/*
Copyright © 2025 Ethan Lee <ethantlee21@gmail.com>
*/
package cmd

import (
	"github.com/ethn1ee/llog/internal/handler"
	"github.com/spf13/cobra"
)

var getOpts handler.GetOpts

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get log entries",
	Long:  `Get log entries from the database. You can get a specific entry by its ID or get all entries.`,
	Args:  cobra.NoArgs,
	RunE:  withLog(get),
}

func get(cmd *cobra.Command, args []string) error {
	handler, err := handler.New(cmd)
	if err != nil {
		return err
	}
	return handler.GetEntry(cmd, args, &getOpts)
}

func init() {
	getCmd.Flags().BoolVarP(&getOpts.Today, "today", "t", false, "get today's entries")
	getCmd.Flags().StringVar(&getOpts.From, "from", "", "date in YYYY-MM-DD format")
	getCmd.Flags().StringVar(&getOpts.To, "to", "", "date in YYYY-MM-DD format")
	rootCmd.AddCommand(getCmd)
}
