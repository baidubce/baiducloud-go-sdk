package ocr

// PaddleVlParserTaskStatusEnum the model 'PaddleVlParserTaskStatusEnum'
type PaddleVlParserTaskStatusEnum string

// List of PaddleVlParserTaskStatusEnum
const (
	PaddleVlParserTaskStatusEnumPending    PaddleVlParserTaskStatusEnum = "pending"
	PaddleVlParserTaskStatusEnumProcessing PaddleVlParserTaskStatusEnum = "processing"
	PaddleVlParserTaskStatusEnumSuccess    PaddleVlParserTaskStatusEnum = "success"
	PaddleVlParserTaskStatusEnumFailed     PaddleVlParserTaskStatusEnum = "failed"
)

// All allowed values of PaddleVlParserTaskStatusEnum enum
var AllowedPaddleVlParserTaskStatusEnumEnumValues = []PaddleVlParserTaskStatusEnum{
	"pending",
	"processing",
	"success",
	"failed",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaddleVlParserTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedPaddleVlParserTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
