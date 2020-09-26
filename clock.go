// A clock allows for using time in tests by providing a time.Time that is
// pinned to a fixed time.
package clock

import (
	"time"
)

//
type TimeFunction func() time.Time

// The Clock type is just a string that is formatted
type Clock string

const (
	// RequiredClockFormat is the format string used for parsing a clock string.
	RequiredClockFormat = "2006-01-02 15:04:05.999999999 -0700 MST"
	// Default will result in a `time.Now()`.
	Default    = Clock("")

	// A Clock pinned to the Epoch.
	Epoch      = Clock("1970-01-01 00:00:00.0 -0000 UTC")
	// A Clock pinned to the Go format string.
	Format     = Clock("2006-01-02 15:04:05.999999999 -0000 UTC")
	// A Clock pinned to the same time as the Go Playground.
	Playground = Clock("2009-11-10 11:00:00.0 -0000 UTC")
	// A Clock pinned to 9/11.
	NineEleven = Clock("2001-09-11 12:46:40.0 -0000 UTC")
	// A Clock pinned to (almost) all 1s.
	AllOnes = Clock("2111-11-11 11:11:11.111111111 -1100 UTC")
)

//
func Now(c Clock) TimeFunction {
	return c.MustTime()
}

//
func (c Clock) MustTime() TimeFunction {
	f, err := c.Time()
	if err != nil {
		panic(err)
	}
	return f
}

//
func (c Clock) Time() (TimeFunction, error) {
	if c == "" {
		return time.Now, nil
	}
	t, err := time.Parse(RequiredClockFormat, string(c))
	if err != nil {
		return time.Now, err
	}
	return func() time.Time {
		return t
	}, nil
}

//
func (c Clock) UTC() TimeFunction {
	return c.MustTime()().UTC
}
