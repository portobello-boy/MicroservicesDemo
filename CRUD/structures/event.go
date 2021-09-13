package structures

import "time"

type Event struct {
	Title     string
	URL       string
	StartTime time.Time
	EndTime   time.Time
	Attendees []Person
}

func (e *Event) Duration() time.Duration {
	return e.EndTime.Sub(e.StartTime)
}
