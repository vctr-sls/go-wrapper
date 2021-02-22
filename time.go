package vctr

import (
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const formatRFC3339NanoWOTimeZone = "2006-01-02T15:04:05.999999999"

func (t *Time) UnmarshalJSON(buf []byte) (err error) {
	data := strings.Trim(string(buf), "\"")
	t.Time, err = time.Parse(time.RFC3339Nano, data)
	if _, ok := err.(*time.ParseError); ok {
		t.Time, err = time.Parse(formatRFC3339NanoWOTimeZone, data)
	}
	return
}
