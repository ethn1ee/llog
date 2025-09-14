package models

import (
	"fmt"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

const (
	TIME_LAYOUT = "2006-01-02 15:04"
	DATE_LAYOUT = "2006-01-02"
)

type Entry struct {
	gorm.Model
	Body string
}

func (e Entry) String() string {
	time := color.HiBlackString(e.CreatedAt.Format(TIME_LAYOUT))

	return fmt.Sprintf("%s %s", time, e.Body)
}
