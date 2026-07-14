package as

// PeriodType the model 'PeriodType'
type PeriodType string

// List of PeriodType
const (
	PeriodTypeDay            PeriodType = "DAY"
	PeriodTypeWeek           PeriodType = "WEEK"
	PeriodTypeMonth          PeriodType = "MONTH"
	PeriodTypeCronexpression PeriodType = "CronExpression"
)

// All allowed values of PeriodType enum
var AllowedPeriodTypeEnumValues = []PeriodType{
	"DAY",
	"WEEK",
	"MONTH",
	"CronExpression",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PeriodType) IsValid() bool {
	for _, existing := range AllowedPeriodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
