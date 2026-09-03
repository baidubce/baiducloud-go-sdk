package image

// SegmentReturnFormEnum the model 'SegmentReturnFormEnum'
type SegmentReturnFormEnum string

// List of SegmentReturnFormEnum
const (
	SegmentReturnFormEnumRgba SegmentReturnFormEnum = "rgba"
	SegmentReturnFormEnumMask SegmentReturnFormEnum = "mask"
)

// All allowed values of SegmentReturnFormEnum enum
var AllowedSegmentReturnFormEnumEnumValues = []SegmentReturnFormEnum{
	"rgba",
	"mask",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SegmentReturnFormEnum) IsValid() bool {
	for _, existing := range AllowedSegmentReturnFormEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
