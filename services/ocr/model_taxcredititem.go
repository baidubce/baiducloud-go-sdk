package ocr

type Taxcredititem struct {
	Taxpayerno   *string `json:"taxpayerno,omitempty"`
	Taxpayername *string `json:"taxpayername,omitempty"`
	Year         *string `json:"year,omitempty"`
	Level        *string `json:"level,omitempty"`
}
