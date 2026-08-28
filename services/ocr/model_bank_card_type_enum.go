package ocr

// BankCardTypeEnum the model 'BankCardTypeEnum'
type BankCardTypeEnum int32

// List of BankCardTypeEnum
const (
	BankCardTypeEnumValue0 BankCardTypeEnum = 0
	BankCardTypeEnumValue1 BankCardTypeEnum = 1
	BankCardTypeEnumValue2 BankCardTypeEnum = 2
	BankCardTypeEnumValue3 BankCardTypeEnum = 3
	BankCardTypeEnumValue4 BankCardTypeEnum = 4
)

// All allowed values of BankCardTypeEnum enum
var AllowedBankCardTypeEnumEnumValues = []BankCardTypeEnum{
	0,
	1,
	2,
	3,
	4,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BankCardTypeEnum) IsValid() bool {
	for _, existing := range AllowedBankCardTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
