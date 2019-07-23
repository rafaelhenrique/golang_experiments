package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond calculate the moment when someone has lived for 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(math.Pow10(9)))
	return t
}
