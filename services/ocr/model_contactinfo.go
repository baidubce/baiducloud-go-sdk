package ocr

type Contactinfo struct {
	Website     []*Website `json:"website,omitempty"`
	Phonenumber *string    `json:"phonenumber,omitempty"`
	Email       *string    `json:"email,omitempty"`
}
