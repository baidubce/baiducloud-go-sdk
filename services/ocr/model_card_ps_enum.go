package ocr

// CardPsEnum the model 'CardPsEnum'
type CardPsEnum int32

// List of CardPsEnum
const (
	CardPsEnumValueMinus1 CardPsEnum = -1
	CardPsEnumValue0      CardPsEnum = 0
	CardPsEnumValue1      CardPsEnum = 1
)

// All allowed values of CardPsEnum enum
var AllowedCardPsEnumEnumValues = []CardPsEnum{
	-1,
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CardPsEnum) IsValid() bool {
	for _, existing := range AllowedCardPsEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
