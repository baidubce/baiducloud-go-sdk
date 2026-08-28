package ocr

type VehicleRegisterCertificationwordsResult struct {
	Number                *VehicleRegistrationCertificateWordsItem `json:"number,omitempty"`
	NameIdcardNo          *VehicleRegistrationCertificateWordsItem `json:"name_idcard_no,omitempty"`
	RegistrationAuthority *VehicleRegistrationCertificateWordsItem `json:"registration_authority,omitempty"`
	RegistrationDate      *VehicleRegistrationCertificateWordsItem `json:"registration_date,omitempty"`
	RegistrationNum       *VehicleRegistrationCertificateWordsItem `json:"registration_num,omitempty"`
	VehicleModel          *VehicleRegistrationCertificateWordsItem `json:"vehicle_model,omitempty"`
	VehicleType           *VehicleRegistrationCertificateWordsItem `json:"vehicle_type,omitempty"`
	Vin                   *VehicleRegistrationCertificateWordsItem `json:"vin,omitempty"`
	EngineNum             *VehicleRegistrationCertificateWordsItem `json:"engine_num,omitempty"`
	SeatingCapacity       *VehicleRegistrationCertificateWordsItem `json:"seating_capacity,omitempty"`
	BodyColor             *VehicleRegistrationCertificateWordsItem `json:"body_color,omitempty"`
	NatureOfUse           *VehicleRegistrationCertificateWordsItem `json:"nature_of_use,omitempty"`
	DateOfProduction      *VehicleRegistrationCertificateWordsItem `json:"date_of_production,omitempty"`
	DateOfIssue           *VehicleRegistrationCertificateWordsItem `json:"date_of_issue,omitempty"`
	SealOfIssueAuthority  *VehicleRegistrationCertificateWordsItem `json:"seal_of_issue_authority,omitempty"`
}
