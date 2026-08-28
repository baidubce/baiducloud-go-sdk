package ocr

type HkMacauTaiwanExitentrypermitResult struct {
	CardNumber        []*HkMacauTaiwanExitentrypermitField `json:"card_number,omitempty"`
	NameChn           []*HkMacauTaiwanExitentrypermitField `json:"name_chn,omitempty"`
	NameEng           []*HkMacauTaiwanExitentrypermitField `json:"name_eng,omitempty"`
	Birthday          []*HkMacauTaiwanExitentrypermitField `json:"birthday,omitempty"`
	Sex               []*HkMacauTaiwanExitentrypermitField `json:"sex,omitempty"`
	ValidDate         []*HkMacauTaiwanExitentrypermitField `json:"valid_date,omitempty"`
	IssueAuthority    []*HkMacauTaiwanExitentrypermitField `json:"issue_authority,omitempty"`
	IssuePlace        []*HkMacauTaiwanExitentrypermitField `json:"issue_place,omitempty"`
	MRZCode           []*HkMacauTaiwanExitentrypermitField `json:"MRZCode,omitempty"`
	HkType            []*HkMacauTaiwanExitentrypermitField `json:"hk_type,omitempty"`
	HkValidDate       []*HkMacauTaiwanExitentrypermitField `json:"hk_valid_date,omitempty"`
	HkRemarks         []*HkMacauTaiwanExitentrypermitField `json:"hk_remarks,omitempty"`
	HkRoundTripNumber []*HkMacauTaiwanExitentrypermitField `json:"hk_round_trip_number,omitempty"`
	McType            []*HkMacauTaiwanExitentrypermitField `json:"mc_type,omitempty"`
	McValidDate       []*HkMacauTaiwanExitentrypermitField `json:"mc_valid_date,omitempty"`
	McRemarks         []*HkMacauTaiwanExitentrypermitField `json:"mc_remarks,omitempty"`
	McRoundTripNumber []*HkMacauTaiwanExitentrypermitField `json:"mc_round_trip_number,omitempty"`
	OcrType           []*HkMacauTaiwanExitentrypermitField `json:"type,omitempty"`
	Remarks           []*HkMacauTaiwanExitentrypermitField `json:"remarks,omitempty"`
	RoundTripNumber   []*HkMacauTaiwanExitentrypermitField `json:"round_trip_number,omitempty"`
	IssueTimes        []*HkMacauTaiwanExitentrypermitField `json:"issue_times,omitempty"`
	IdcardName        []*HkMacauTaiwanExitentrypermitField `json:"idcard_name,omitempty"`
	IdcardNumber      []*HkMacauTaiwanExitentrypermitField `json:"idcard_number,omitempty"`
	MRZCode1          []*HkMacauTaiwanExitentrypermitField `json:"MRZCode1,omitempty"`
	MRZCode2          []*HkMacauTaiwanExitentrypermitField `json:"MRZCode2,omitempty"`
}
