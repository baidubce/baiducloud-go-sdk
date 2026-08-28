package ocr

type Pledge struct {
	Registno      *string `json:"registno,omitempty"`
	Pledgor       *string `json:"pledgor,omitempty"`
	Pledgorno     *string `json:"pledgorno,omitempty"`
	Pledgee       *string `json:"pledgee,omitempty"`
	Pledgeeno     *string `json:"pledgeeno,omitempty"`
	Pledgedamount *string `json:"pledgedamount,omitempty"`
	Regdate       *string `json:"regdate,omitempty"`
	Publicdate    *string `json:"publicdate,omitempty"`
	Status        *string `json:"status,omitempty"`
}
