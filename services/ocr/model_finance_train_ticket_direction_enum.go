package ocr

// FinanceTrainTicketDirectionEnum the model 'FinanceTrainTicketDirectionEnum'
type FinanceTrainTicketDirectionEnum int32

// List of FinanceTrainTicketDirectionEnum
const (
	FinanceTrainTicketDirectionEnumValueMinus1 FinanceTrainTicketDirectionEnum = -1
	FinanceTrainTicketDirectionEnumValue0      FinanceTrainTicketDirectionEnum = 0
	FinanceTrainTicketDirectionEnumValue1      FinanceTrainTicketDirectionEnum = 1
	FinanceTrainTicketDirectionEnumValue2      FinanceTrainTicketDirectionEnum = 2
	FinanceTrainTicketDirectionEnumValue3      FinanceTrainTicketDirectionEnum = 3
)

// All allowed values of FinanceTrainTicketDirectionEnum enum
var AllowedFinanceTrainTicketDirectionEnumEnumValues = []FinanceTrainTicketDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FinanceTrainTicketDirectionEnum) IsValid() bool {
	for _, existing := range AllowedFinanceTrainTicketDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
