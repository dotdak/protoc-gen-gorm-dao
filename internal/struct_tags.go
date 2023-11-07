package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type structTags [][2]string

func (tags structTags) String() string {
	if len(tags) == 0 {
		return ""
	}
	var ss []string
	for _, tag := range tags {
		// NOTE: When quoting the value, we need to make sure the backtick
		// character does not appear. Convert all cases to the escaped hex form.
		key := tag[0]
		val := strings.Replace(strconv.Quote(tag[1]), "`", `\x60`, -1)
		ss = append(ss, fmt.Sprintf("%s:%s", key, val))
	}
	return "`" + strings.Join(ss, " ") + "`"
}
