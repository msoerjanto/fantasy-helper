package gql

import (
	"fmt"
	"time"
)

func ConvertStringToTime(timeStr string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, timeStr)

	if err != nil {
		fmt.Println(err)
	}
	return t
}
