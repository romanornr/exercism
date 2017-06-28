package clock

import "fmt"

const testVersion = 4

type Clock struct {
    hour    int
    minute  int
}

func New(hour, minute int) Clock {
    nMinute := minute % 60
    if nMinute < 0 {
        nMinute = nMinute + 60
        hour--
    }
    nHour := (hour + (minute / 60)) % 24
    if nHour < 0 {
        nHour = nHour + 24
    }
    return Clock{ hour: nHour, minute: nMinute }

}

// Returns a string represenation of the clock
func (c *Clock) String() string {

    return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
    return New(c.hour, c.minute + minutes)
}
