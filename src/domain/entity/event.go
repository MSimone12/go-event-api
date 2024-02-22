package entity

import "strconv"

// Event is the struct defined for the json body from the original request
type Event struct {
	Type        string `json:"type,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      uint   `json:"amount,omitempty"`
	Destination string `json:"destination,omitempty"`
}

// GetOrigin gets and parses origin from the Event struct
func (e Event) GetOrigin() uint {
	origin, _ := strconv.Atoi(e.Origin)
	return uint(origin)
}

// GetDestination gets and parses destination from the Event struct
func (e Event) GetDestination() uint {
	destination, _ := strconv.Atoi(e.Destination)
	return uint(destination)
}
