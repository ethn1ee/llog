package handler

import (
	"fmt"
	"time"

	"github.com/ethn1ee/llog/internal/config"
	"github.com/spf13/cobra"
)

type Opts interface {
	Validate(cfg *config.Config) error
}

func ValidateOptions(cfg *config.Config, opts Opts) HandlerFunc {
	return func(cmd *cobra.Command, args []string) error {
		return opts.Validate(cfg)
	}
}

type GetOpts struct {
	Today     bool
	Yesterday bool
	From      string
	To        string

	fromTime time.Time
	toTime   time.Time
}

const (
	mutexOptError = "option '%s' cannot be used with option '%s'"
	timeParseError = "failed to parse time: %w"
)

func (o *GetOpts) Validate(cfg *config.Config) error {
	// mutual exclusion checks
	if o.Today && o.Yesterday {
		return fmt.Errorf(mutexOptError, "today", "yesterday")
	}
	if o.Today && o.From != "" {
		return fmt.Errorf(mutexOptError, "today", "from")
	}
	if o.Today && o.To != "" {
		return fmt.Errorf(mutexOptError, "today", "to")
	}
	if o.Yesterday && o.From != "" {
		return fmt.Errorf(mutexOptError, "yesterday", "from")
	}
	if o.Yesterday && o.To != "" {
		return fmt.Errorf(mutexOptError, "yesterday", "to")
	}

	// set fromTime and toTime
	if o.Today {
		now := time.Now()
		o.fromTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		o.toTime = o.fromTime.Add(24 * time.Hour)
		return nil
	}

	if o.Yesterday {
		now := time.Now()
		yesterday := now.Add(-24 * time.Hour)
		o.fromTime = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
		o.toTime = o.fromTime.Add(24 * time.Hour)
		return nil
	}

	if o.From != "" {
		from, err := time.Parse(cfg.DateLayout, o.From)
		if err != nil {
			return fmt.Errorf(timeParseError, err)
		}
		o.fromTime = from
	}

	if o.To != "" {
		to, err := time.Parse(cfg.DateLayout, o.To)
		if err != nil {
			return fmt.Errorf(timeParseError, err)
		}
		o.toTime = to
	}

	return nil
}

type AddOpts struct{}

func (o *AddOpts) Validate(cfg *config.Config) error {
	return nil
}
