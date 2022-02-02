package date_utils

import (
	"time"

	"maranatha_web/internal/logger"
)

const (
	apiDateLayout = "2 Jan 2006 15:04:05"
	apiDBLayout   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {

	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)

}

func StringToDate(str string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)

	if err != nil {
		logger.Error("Please provide a valid date format.", err)
		logger.GetLogger().Fatal("Please provide a valid date format.")
	}
	return t
}
