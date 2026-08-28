package ocr

// BankcardDirectionEnum the model 'BankcardDirectionEnum'
type BankcardDirectionEnum int32

// List of BankcardDirectionEnum
const (
	BankcardDirectionEnumValueMinus1 BankcardDirectionEnum = -1
	BankcardDirectionEnumValue0      BankcardDirectionEnum = 0
	BankcardDirectionEnumValue1      BankcardDirectionEnum = 1
	BankcardDirectionEnumValue2      BankcardDirectionEnum = 2
	BankcardDirectionEnumValue3      BankcardDirectionEnum = 3
)

// All allowed values of BankcardDirectionEnum enum
var AllowedBankcardDirectionEnumEnumValues = []BankcardDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankcardDirectionEnum) IsValid() bool {
	for _, existing := range AllowedBankcardDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
