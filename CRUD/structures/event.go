package structures

import (
	"net/http"
	"time"
)

type Event struct {
	Title     string    `json:title`
	URL       string    `json:url`
	StartTime time.Time `json:startTime`
	EndTime   time.Time `json:endTime`
	Attendees []Person  `json:attendees`
}

func (e *Event) Duration() time.Duration {
	return e.EndTime.Sub(e.StartTime)
}

func (e *Event) AddAttendee(person Person) {
	e.Attendees = append(e.Attendees, person)
}

func CreateEventFromRequest(request *http.Request) *Event {
	// var event Event
	// err := json.NewDecoder(request.Body).Decode(&event)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// return &event
	return nil
}
