package ocr

type Mpledge struct {
	Registerno        *string `json:"registerno,omitempty"`
	Registerdate      *string `json:"registerdate,omitempty"`
	Publicdate        *string `json:"publicdate,omitempty"`
	Registeroffice    *string `json:"registeroffice,omitempty"`
	Debtsecuredamount *string `json:"debtsecuredamount,omitempty"`
	Status            *string `json:"status,omitempty"`
}
