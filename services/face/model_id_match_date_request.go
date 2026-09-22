package face

type IdMatchDateRequest struct {
	Name         *string `json:"name,omitempty"`
	IdCardNumber *string `json:"id_card_number,omitempty"`
	StartDate    *string `json:"start_date,omitempty"`
	EndDate      *string `json:"end_date,omitempty"`
}
