package entity

import "strconv"

type Event struct {
	Type        string `json:"type,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      uint   `json:"amount,omitempty"`
	Destination string `json:"destination,omitempty"`
}

func (e Event) GetOrigin() uint {
	origin, _ := strconv.Atoi(e.Origin)
	return uint(origin)
}

func (e Event) GetDestination() uint {
	destination, _ := strconv.Atoi(e.Destination)
	return uint(destination)
}
