// Package gigasecond calculates the moment when someone has lived for
// 1 gigasecond. A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the passed Time and returns it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
