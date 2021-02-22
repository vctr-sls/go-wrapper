package vctr

import (
	"net/url"
	"strconv"
)

func getLimitOffsetQuery(limit, offset int) url.Values {
	return url.Values{
		"limit":  []string{strconv.Itoa(limit)},
		"offset": []string{strconv.Itoa(offset)},
	}
}
