package ocr

type Permission struct {
	Name     *string `json:"name,omitempty"`
	Province *string `json:"province,omitempty"`
	Liandate *string `json:"liandate,omitempty"`
	Caseno   *string `json:"caseno,omitempty"`
}
