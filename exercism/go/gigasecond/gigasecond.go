package gigasecond

import (
	"math"
	"time"
)

// AddGigasecondV1 calculate the moment when someone has lived for 10^9 seconds
func AddGigasecondV1(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(math.Pow10(9)))
	return t
}

// AddGigasecond calculate the moment when someone has lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9)
	return t
}
