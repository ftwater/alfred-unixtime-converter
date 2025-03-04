package date

import (
	"strconv"
	"time"
)

func (formatter unixTimeInputFormatter) ToTime() (time.Time, error) {
	unixtime, _ := strconv.ParseInt(formatter.input, 10, 64)
	t := time.Now()
	if len(formatter.input) > 10 {
		t = time.UnixMilli(unixtime)
	} else {
		t = time.Unix(unixtime, 0)
	}
	return t, nil
}

func (formatter unixTimeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
