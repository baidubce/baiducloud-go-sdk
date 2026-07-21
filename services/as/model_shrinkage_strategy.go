package as

// ShrinkageStrategy the model 'ShrinkageStrategy'
type ShrinkageStrategy string

// List of ShrinkageStrategy
const (
	ShrinkageStrategyEarlier ShrinkageStrategy = "Earlier"
	ShrinkageStrategyLater   ShrinkageStrategy = "Later"
)

// All allowed values of ShrinkageStrategy enum
var AllowedShrinkageStrategyEnumValues = []ShrinkageStrategy{
	"Earlier",
	"Later",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShrinkageStrategy) IsValid() bool {
	for _, existing := range AllowedShrinkageStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
