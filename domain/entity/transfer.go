package entity

type Transfer struct {
	Origin      Origin      `json:"origin"`
	Destination Destination `json:"destination"`
}
