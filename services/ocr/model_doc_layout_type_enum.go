package ocr

// DocLayoutTypeEnum the model 'DocLayoutTypeEnum'
type DocLayoutTypeEnum string

// List of DocLayoutTypeEnum
const (
	DocLayoutTypeEnumTable       DocLayoutTypeEnum = "table"
	DocLayoutTypeEnumFigure      DocLayoutTypeEnum = "figure"
	DocLayoutTypeEnumText        DocLayoutTypeEnum = "text"
	DocLayoutTypeEnumTitle       DocLayoutTypeEnum = "title"
	DocLayoutTypeEnumContents    DocLayoutTypeEnum = "contents"
	DocLayoutTypeEnumSeal        DocLayoutTypeEnum = "seal"
	DocLayoutTypeEnumTableTitle  DocLayoutTypeEnum = "table_title"
	DocLayoutTypeEnumFigureTitle DocLayoutTypeEnum = "figure_title"
	DocLayoutTypeEnumDocTitle    DocLayoutTypeEnum = "doc_title"
)

// All allowed values of DocLayoutTypeEnum enum
var AllowedDocLayoutTypeEnumEnumValues = []DocLayoutTypeEnum{
	"table",
	"figure",
	"text",
	"title",
	"contents",
	"seal",
	"table_title",
	"figure_title",
	"doc_title",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocLayoutTypeEnum) IsValid() bool {
	for _, existing := range AllowedDocLayoutTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
