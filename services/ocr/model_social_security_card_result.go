package ocr

type SocialSecurityCardResult struct {
	CardNumber           *SocialSecurityCardField `json:"card_number,omitempty"`
	Name                 *SocialSecurityCardField `json:"name,omitempty"`
	Sex                  *SocialSecurityCardField `json:"sex,omitempty"`
	BirthDate            *SocialSecurityCardField `json:"birth_date,omitempty"`
	SocialSecurityNumber *SocialSecurityCardField `json:"social_security_number,omitempty"`
	IssueDate            *SocialSecurityCardField `json:"issue_date,omitempty"`
	BankCardNumber       *SocialSecurityCardField `json:"bank_card_number,omitempty"`
	ExpiryDate           *SocialSecurityCardField `json:"expiry_date,omitempty"`
}
