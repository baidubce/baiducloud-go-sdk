package ocr

// DocAnalysisLayoutTypeEnum the model 'DocAnalysisLayoutTypeEnum'
type DocAnalysisLayoutTypeEnum string

// List of DocAnalysisLayoutTypeEnum
const (
	DocAnalysisLayoutTypeEnumTable    DocAnalysisLayoutTypeEnum = "table"
	DocAnalysisLayoutTypeEnumFigure   DocAnalysisLayoutTypeEnum = "figure"
	DocAnalysisLayoutTypeEnumText     DocAnalysisLayoutTypeEnum = "text"
	DocAnalysisLayoutTypeEnumTitle    DocAnalysisLayoutTypeEnum = "title"
	DocAnalysisLayoutTypeEnumContents DocAnalysisLayoutTypeEnum = "contents"
)

// All allowed values of DocAnalysisLayoutTypeEnum enum
var AllowedDocAnalysisLayoutTypeEnumEnumValues = []DocAnalysisLayoutTypeEnum{
	"table",
	"figure",
	"text",
	"title",
	"contents",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocAnalysisLayoutTypeEnum) IsValid() bool {
	for _, existing := range AllowedDocAnalysisLayoutTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
