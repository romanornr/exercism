package gigasecond

// import path for the time package from the standard library
import "time"

const (
    testVersion = 4
    giga = 1000000000
)

func AddGigasecond(t time.Time) time.Time {
    return t.Add(time.Duration(giga) * time.Second)
}
