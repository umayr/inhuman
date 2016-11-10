package inhuman

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	regNSpc = regexp.MustCompile(`(\d+)([a-zA-Z]+)`)
	regSpc  = regexp.MustCompile(`\s+`)
	regHr   = regexp.MustCompile("^(h|hr|hrs|hour|hours)$")
	regMin  = regexp.MustCompile("^(m|min|mins|minute|minutes)$")
	regSec  = regexp.MustCompile("^(s|sec|secs|second|seconds)$")
)

// Parse converts humanized timestamps into nanoseconds
// so that machines could understand them and help kree empire to rule over earth
func Parse(str string) (t time.Duration, err error) {
	raw := regNSpc.ReplaceAllString(str, "$1 $2 ")
	parts := regSpc.Split(strings.TrimSpace(raw), -1)

calc:
	d, err := calc(parts[0:2])
	if err != nil {
		return
	}
	t += d

	if len(parts) > 2 {
		parts = parts[2:]
		goto calc
	}

	return
}

func calc(parts []string) (t time.Duration, err error) {
	var factor time.Duration

	if regHr.MatchString(parts[1]) {
		factor = time.Hour
	}

	if regMin.MatchString(parts[1]) {
		factor = time.Minute
	}

	if regSec.MatchString(parts[1]) {
		factor = time.Second
	}

	return parse(parts[0], factor)
}

func parse(str string, factor time.Duration) (t time.Duration, err error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}

	return time.Duration(i) * factor, nil
}
