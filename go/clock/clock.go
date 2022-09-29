package clock

import (
	"fmt"
)

// Clock that handles times without dates
type Clock struct {
	minutes int
}

// New will create new Clock
func New(h, m int) Clock {
	minutes := (h*60 + m) % 1440
	if minutes < 0 {
		minutes += 1440
	}
	return Clock{minutes}
}

// Add will add minutes to Clock
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

// Subtract will subtract minutes to Clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}

// String present string of Clock
func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
