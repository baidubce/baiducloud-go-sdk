package ocr

// BankcardClearStatusEnum the model 'BankcardClearStatusEnum'
type BankcardClearStatusEnum int32

// List of BankcardClearStatusEnum
const (
	BankcardClearStatusEnumValue0 BankcardClearStatusEnum = 0
	BankcardClearStatusEnumValue1 BankcardClearStatusEnum = 1
)

// All allowed values of BankcardClearStatusEnum enum
var AllowedBankcardClearStatusEnumEnumValues = []BankcardClearStatusEnum{
	0,
	1,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankcardClearStatusEnum) IsValid() bool {
	for _, existing := range AllowedBankcardClearStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
