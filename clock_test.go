package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock_Time(t *testing.T) {
	tests := []struct {
		name string
		c    Clock
		want time.Time
	}{
		// {name: "Default", c: Default, want: time.Now().UTC().Truncate(time.Second)},
		{name: "Epoch", c: Epoch, want: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
		{name: "Format", c: Format, want: time.Date(2006, 1, 2, 15, 4, 5, 987654321, time.UTC)},
		{name: "Playground", c: Playground, want: time.Date(2009, 11, 10, 11, 0, 0, 0, time.UTC)},
		{name: "NineElevent", c: NineEleven, want: time.Date(2001, 9, 11, 12, 46, 40, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.c.MustTime()()

			assert.Equal(t, c.Year(), tt.want.Year())
			assert.Equal(t, c.Month(), tt.want.Month())
			assert.Equal(t, c.Day(), tt.want.Day())
			assert.Equal(t, c.Hour(), tt.want.Hour())
			assert.Equal(t, c.Minute(), tt.want.Minute())
			assert.Equal(t, c.Second(), tt.want.Second())
			if tt.want.Nanosecond() != 0 {
				assert.Equal(t, c.Nanosecond(), tt.want.Nanosecond())
			}
			assert.Equal(t, c.Location().String(), tt.want.Location().String())
		})
	}
}

func TestClock_Time_Error(t *testing.T) {
	tests := []struct {
		name string
		time string
	}{
		{name: "Y", time: "1970"},
		{name: "YM", time: "1970-01"},
		{name: "YMD", time: "1970-01-02"},
		{name: "YMDH", time: "1970-01-02 15"},
		{name: "YMDHM", time: "1970-01-02 15:04"},
		{name: "YMDHMS", time: "1970-01-02 15:04:05"},
		{name: "YMDHMSN", time: "1970-01-02 15:04:05.987654321"},
		{name: "Z", time: "-0000"},
		{name: "YZ", time: "1970 -0000"},
		{name: "YMZ", time: "1970-01 -0000"},
		{name: "YMDZ", time: "1970-01-02 -0000"},
		{name: "YMDHZ", time: "1970-01-02 15 -0000"},
		{name: "YMDHMZ", time: "1970-01-02 15:04 -0000"},
		{name: "YMDHMSZ", time: "1970-01-02 15:04:05 -0000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Clock(tt.time).Time()
			assert.Error(t, err, "should have failed")
		})
	}
}
