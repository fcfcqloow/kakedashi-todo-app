package domain

import (
	"fmt"
	"strings"

	cnv "github.com/fcfcqloow/go-advance/convert"
)

type (
	Date     string
	Log      string
	DateList []Date
)

func FilterYearMonth(d []Date, year, month int) []Date {
	isOk := func(date Date) bool {
		if year > 0 && !strings.Contains(string(date), cnv.MustStr(year)+"_") {
			return false
		}

		if month > 0 && !strings.Contains(string(date), "_"+fmt.Sprintf("%02d", month)) {
			return false
		}

		return true
	}

	return filter(d, isOk)
}
