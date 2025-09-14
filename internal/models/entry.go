package models

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

const (
	TIME_LAYOUT = "2006-01-02 15:04"
)

type Entry struct {
	Time int64
	Body string
}

func (e Entry) String() string {
	time := color.HiBlackString(time.Unix(e.Time, 0).Format(TIME_LAYOUT))

	return fmt.Sprintf("%s %s", time, e.Body)
}
