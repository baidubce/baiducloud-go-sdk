package ocr

type Branch struct {
	Companycode *string `json:"companycode,omitempty"`
	Companyname *string `json:"companyname,omitempty"`
	Authority   *string `json:"authority,omitempty"`
	Creditno    *string `json:"creditno,omitempty"`
	Legalperson *string `json:"legalperson,omitempty"`
}
