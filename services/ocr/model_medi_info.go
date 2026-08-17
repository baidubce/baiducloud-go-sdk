package ocr

type MediInfo struct {
	MediCode     *string `json:"medi_code,omitempty"`
	MediName     *string `json:"medi_name,omitempty"`
	MediRegister *string `json:"medi_register,omitempty"`
	MediType     *string `json:"medi_type,omitempty"`
	MediRegion   *string `json:"medi_region,omitempty"`
	MediCheck    *int32  `json:"medi_check,omitempty"`
}
