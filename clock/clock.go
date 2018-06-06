/*
Package clock implements a clock that handles times without dates.
You should be able to add and subtract minutes to it. Two clocks
that represent the same time should be equal to each other.
*/
package clock

import (
	"fmt"
)

// Clock is defined by the Hours and Minutes which are both int.
type Clock struct {
	Hours, Minutes int
}

// Define the constants for the program.
const (
	minutesInHour = 60
	hoursInDay    = 24
)

// String implements the Stringer interface. Returns a zero padded time.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

// Add adds the total number of minutes to the current time.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.totalMinutes()+minutes)
}

// Subtract substracts the total number of minutes to the current time.
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.totalMinutes()-minutes)
}

// New creates a new value of type Clock.
func New(hours, minutes int) Clock {
	// Normalize the total number of minutes.
	if minutes < 0 {
		m := minutes * -1
		hours -= 1 + (m / minutesInHour)
		minutes = minutesInHour - (m % minutesInHour)
	} else {
		hours += minutes / minutesInHour
		minutes %= minutesInHour
	}

	// Normalize the total number of hours.
	hours = hoursInDay + (hours % hoursInDay)
	hours %= hoursInDay

	return Clock{hours, minutes}
}

// totalMinutes calculates the total minutes of a Clock.
func (c Clock) totalMinutes() int {
	return c.Minutes + (c.Hours * minutesInHour)
}
