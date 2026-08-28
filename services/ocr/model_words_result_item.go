package ocr

type WordsResultItem struct {
	CardType    *string                    `json:"card_type,omitempty"`
	Direction   *int32                     `json:"direction,omitempty"`
	Probability *float32                   `json:"probability,omitempty"`
	Location    *MixedMultiVehicleLocation `json:"location,omitempty"`
	LicenseInfo []*LicenseInfoItem         `json:"license_info,omitempty"`
}
