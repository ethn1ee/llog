package handler

import (
	"fmt"

	"github.com/ethn1ee/llog/internal/config"
	_db "github.com/ethn1ee/llog/internal/db"
	"github.com/spf13/cobra"
)

func Summarize(cfg *config.Config, db *_db.DB, opts *SummarizeOpts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Println("summarize called")
		return nil
	}
}

type SummarizeOpts struct {
	Time timeOpts
}

func (o *SummarizeOpts) applyFlags(cmd *cobra.Command) {
	o.Time.applyFlags(cmd)
}

func (o *SummarizeOpts) validate(cfg *config.Config, cmd *cobra.Command, args []string) error {
	return o.Time.validate(cfg, cmd, args)
}
