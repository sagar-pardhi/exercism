package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Millisecond * 1000 * 1000000000
	return t.Add(gigasecond)
}
