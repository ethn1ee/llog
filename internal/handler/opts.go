package handler

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethn1ee/llog/internal/models"
)

type GetOpts struct {
	Today    bool
	From     string
	To       string

	fromTime time.Time
	toTime   time.Time
}

func (o *GetOpts) validate() error {
	var from, to time.Time
	var err error

	if o.Today {
		if o.From != "" {
			return errors.New("flag 'from' cannot be used with flag 'today'")
		}
		if o.To != "" {
			return errors.New("flag 'to' cannot be used with flag 'today'")
		}

		now := time.Now()
		from = now.Truncate(24 * time.Hour)
		to = from.Add(24 * time.Hour)
	} else if o.From != "" {
		from, err = time.Parse(models.DATE_LAYOUT, o.From)
		if err != nil {
			return fmt.Errorf("failed to parse time: %w", err)
		}
	} else if o.To != "" {
		to, err = time.Parse(models.DATE_LAYOUT, o.To)
		if err != nil {
			return fmt.Errorf("failed to parse time: %w", err)
		}
	}

	o.fromTime = from
	o.toTime = to

	return nil
}
