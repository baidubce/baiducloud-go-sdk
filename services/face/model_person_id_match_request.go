package face

type PersonIdMatchRequest struct {
	IdCardNumber *string `json:"id_card_number,omitempty"`
	Name         *string `json:"name,omitempty"`
}
