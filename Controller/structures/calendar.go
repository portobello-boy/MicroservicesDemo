package structures

type Calendar struct {
	Events []*Event
}

func (c *Calendar) AddEvent(event *Event) {
	c.Events = append(c.Events, event)
}
