package ocr

type Change struct {
	Changefield  *string `json:"changefield,omitempty"`
	Changebefore *string `json:"changebefore,omitempty"`
	Changeafter  *string `json:"changeafter,omitempty"`
	Changedate   *string `json:"changedate,omitempty"`
}
