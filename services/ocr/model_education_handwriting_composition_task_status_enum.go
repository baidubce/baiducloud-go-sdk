package ocr

// EducationHandwritingCompositionTaskStatusEnum the model 'EducationHandwritingCompositionTaskStatusEnum'
type EducationHandwritingCompositionTaskStatusEnum string

// List of EducationHandwritingCompositionTaskStatusEnum
const (
	EducationHandwritingCompositionTaskStatusEnumPending    EducationHandwritingCompositionTaskStatusEnum = "pending"
	EducationHandwritingCompositionTaskStatusEnumProcessing EducationHandwritingCompositionTaskStatusEnum = "processing"
	EducationHandwritingCompositionTaskStatusEnumSuccess    EducationHandwritingCompositionTaskStatusEnum = "success"
	EducationHandwritingCompositionTaskStatusEnumFailed     EducationHandwritingCompositionTaskStatusEnum = "failed"
)

// All allowed values of EducationHandwritingCompositionTaskStatusEnum enum
var AllowedEducationHandwritingCompositionTaskStatusEnumEnumValues = []EducationHandwritingCompositionTaskStatusEnum{
	"pending",
	"processing",
	"success",
	"failed",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationHandwritingCompositionTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedEducationHandwritingCompositionTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
