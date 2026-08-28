package ocr

type HouseholdRegisterRequest struct {
	Image                 *string `json:"image,omitempty"`
	Url                   *string `json:"url,omitempty"`
	HouseholdRegisterSide *string `json:"household_register_side,omitempty"`
}
