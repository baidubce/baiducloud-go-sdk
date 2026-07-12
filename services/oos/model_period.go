package oos

type Period struct {
	Frequency       *string  `json:"frequency,omitempty"`
	DaysOfWeek      []*int32 `json:"daysOfWeek,omitempty"`
	Dates           []*int32 `json:"dates,omitempty"`
	LastDateOfMonth *bool    `json:"lastDateOfMonth,omitempty"`
	Hour            *int32   `json:"hour,omitempty"`
	Minute          *int32   `json:"minute,omitempty"`
}
