package ocr

// ParserTaskStatusEnum the model 'ParserTaskStatusEnum'
type ParserTaskStatusEnum string

// List of ParserTaskStatusEnum
const (
	ParserTaskStatusEnumPending    ParserTaskStatusEnum = "pending"
	ParserTaskStatusEnumProcessing ParserTaskStatusEnum = "processing"
	ParserTaskStatusEnumSuccess    ParserTaskStatusEnum = "success"
	ParserTaskStatusEnumFailed     ParserTaskStatusEnum = "failed"
)

// All allowed values of ParserTaskStatusEnum enum
var AllowedParserTaskStatusEnumEnumValues = []ParserTaskStatusEnum{
	"pending",
	"processing",
	"success",
	"failed",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ParserTaskStatusEnum) IsValid() bool {
	for _, existing := range AllowedParserTaskStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
