package image

// SegmentMethodEnum the model 'SegmentMethodEnum'
type SegmentMethodEnum string

// List of SegmentMethodEnum
const (
	SegmentMethodEnumAuto    SegmentMethodEnum = "auto"
	SegmentMethodEnumControl SegmentMethodEnum = "control"
)

// All allowed values of SegmentMethodEnum enum
var AllowedSegmentMethodEnumEnumValues = []SegmentMethodEnum{
	"auto",
	"control",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SegmentMethodEnum) IsValid() bool {
	for _, existing := range AllowedSegmentMethodEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
