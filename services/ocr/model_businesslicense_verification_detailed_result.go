package ocr

type BusinesslicenseVerificationDetailedResult struct {
	Base           *Base            `json:"base,omitempty"`
	Branches       []*Branch        `json:"branches,omitempty"`
	Changes        []*Change        `json:"changes,omitempty"`
	Taxcredititems []*Taxcredititem `json:"taxcredititems,omitempty"`
	Contactinfo    *Contactinfo     `json:"contactinfo,omitempty"`
	Employees      []*Employee      `json:"employees,omitempty"`
	Exceptions     []*ExceptionInfo `json:"exceptions,omitempty"`
	Industry       []*Industry      `json:"industry,omitempty"`
	Liquidation    *Liquidation     `json:"liquidation,omitempty"`
	Mpledges       []*Mpledge       `json:"mpledges,omitempty"`
	Originalname   []*Originalname  `json:"originalname,omitempty"`
	Partners       []*Partner       `json:"partners,omitempty"`
	Penalties      []*Penalty       `json:"penalties,omitempty"`
	Permissions    []*Permission    `json:"permissions,omitempty"`
	Pledges        []*Pledge        `json:"pledges,omitempty"`
	Spotchecks     []*Spotcheck     `json:"spotchecks,omitempty"`
	Shixinitems    []*Shixinitem    `json:"shixinitems,omitempty"`
	Zhixingitems   []*Zhixingitem   `json:"zhixingitems,omitempty"`
}
