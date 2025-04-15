package podcast

import (
	"strconv"
	"strings"
)

type ITunesID string

func (id ITunesID) Int() int {
	// strip prefix if present
	idStr := strings.TrimPrefix(string(id), "id")
	i, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return i
}
