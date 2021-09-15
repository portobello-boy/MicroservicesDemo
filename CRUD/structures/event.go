package structures

import (
	"net/http"
	"time"
)

type Event struct {
	ObjectID  string    `json:"_id"        bson:"_id"`
	Title     string    `json:"title"      bson:"title"`
	URL       string    `json:"url"        bson:"url"`
	StartTime time.Time `json:"startTime"  bson:"starttime"`
	EndTime   time.Time `json:"endTime"    bson:"endtime"`
	AllDay    bool      `json:"allDay"     bson:"allday"`
	Attendees []Person  `json:"attendees"  bson:"attendees"`
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
