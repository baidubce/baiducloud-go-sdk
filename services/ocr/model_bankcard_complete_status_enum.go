package ocr

// BankcardCompleteStatusEnum the model 'BankcardCompleteStatusEnum'
type BankcardCompleteStatusEnum int32

// List of BankcardCompleteStatusEnum
const (
	BankcardCompleteStatusEnumValue0 BankcardCompleteStatusEnum = 0
	BankcardCompleteStatusEnumValue1 BankcardCompleteStatusEnum = 1
)

// All allowed values of BankcardCompleteStatusEnum enum
var AllowedBankcardCompleteStatusEnumEnumValues = []BankcardCompleteStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankcardCompleteStatusEnum) IsValid() bool {
	for _, existing := range AllowedBankcardCompleteStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
