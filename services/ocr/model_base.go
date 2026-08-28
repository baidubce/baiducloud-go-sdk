package ocr

type Base struct {
	Legalperson      *string `json:"legalperson,omitempty"`
	Establishdate    *string `json:"establishdate,omitempty"`
	Revokedate       *string `json:"revokedate,omitempty"`
	Companystatus    *string `json:"companystatus,omitempty"`
	Province         *string `json:"province,omitempty"`
	Creditno         *string `json:"creditno,omitempty"`
	Capital          *string `json:"capital,omitempty"`
	Companytype      *string `json:"companytype,omitempty"`
	Companyaddress   *string `json:"companyaddress,omitempty"`
	Businessscope    *string `json:"businessscope,omitempty"`
	Businessdatefrom *string `json:"businessdatefrom,omitempty"`
	Businessdateto   *string `json:"businessdateto,omitempty"`
	Issuedate        *string `json:"issuedate,omitempty"`
	Orgcode          *string `json:"orgcode,omitempty"`
	Isonstock        *string `json:"isonstock,omitempty"`
	Stocknumber      *string `json:"stocknumber,omitempty"`
	Stocktype        *string `json:"stocktype,omitempty"`
	Keyno            *string `json:"keyno,omitempty"`
	Companyname      *string `json:"companyname,omitempty"`
	Companycode      *string `json:"companycode,omitempty"`
	Authority        *string `json:"authority,omitempty"`
	Regcapcur        *string `json:"regcapcur,omitempty"`
}
