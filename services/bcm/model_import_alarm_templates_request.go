package bcm

type ImportAlarmTemplatesRequest struct {
	Overwrite *bool       `json:"overwrite,omitempty"`
	Templates []*Template `json:"templates,omitempty"`
}
