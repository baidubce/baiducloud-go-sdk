package ocr

type BusinesslicenseVerificationStandardResult struct {
	Companyname           *string   `json:"companyname,omitempty"`
	Companytype           *string   `json:"companytype,omitempty"`
	Legalperson           *string   `json:"legalperson,omitempty"`
	Capital               *string   `json:"capital,omitempty"`
	Companycode           *string   `json:"companycode,omitempty"`
	Companyaddress        *string   `json:"companyaddress,omitempty"`
	Businessscope         *string   `json:"businessscope,omitempty"`
	Authority             *string   `json:"authority,omitempty"`
	Companystatus         *string   `json:"companystatus,omitempty"`
	Establishdate         *string   `json:"establishdate,omitempty"`
	Creditno              *string   `json:"creditno,omitempty"`
	Operationstartdate    *string   `json:"operationstartdate,omitempty"`
	Operationenddate      *string   `json:"operationenddate,omitempty"`
	Issuedate             *string   `json:"issuedate,omitempty"`
	Province              *string   `json:"province,omitempty"`
	Provincecode          *string   `json:"provincecode,omitempty"`
	City                  *string   `json:"city,omitempty"`
	Citycode              *string   `json:"citycode,omitempty"`
	District              *string   `json:"district,omitempty"`
	Districtcode          *string   `json:"districtcode,omitempty"`
	Regcapcur             *string   `json:"regcapcur,omitempty"`
	Orgcode               *string   `json:"orgcode,omitempty"`
	Licensedbusinessscope *string   `json:"licensedbusinessscope,omitempty"`
	Companyenglishname    *string   `json:"companyenglishname,omitempty"`
	Onceusedname          []*string `json:"onceusedname,omitempty"`
	Orgcompanycode        *string   `json:"orgcompanycode,omitempty"`
	Paidincapital         *string   `json:"paidincapital,omitempty"`
	Revokedate            *string   `json:"revokedate,omitempty"`
	Logoffdate            *string   `json:"logoffdate,omitempty"`
}
