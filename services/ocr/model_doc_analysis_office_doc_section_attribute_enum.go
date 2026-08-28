package ocr

// DocAnalysisOfficeDocSectionAttributeEnum the model 'DocAnalysisOfficeDocSectionAttributeEnum'
type DocAnalysisOfficeDocSectionAttributeEnum string

// List of DocAnalysisOfficeDocSectionAttributeEnum
const (
	DocAnalysisOfficeDocSectionAttributeEnumSection  DocAnalysisOfficeDocSectionAttributeEnum = "section"
	DocAnalysisOfficeDocSectionAttributeEnumHeader   DocAnalysisOfficeDocSectionAttributeEnum = "header"
	DocAnalysisOfficeDocSectionAttributeEnumFooter   DocAnalysisOfficeDocSectionAttributeEnum = "footer"
	DocAnalysisOfficeDocSectionAttributeEnumNumber   DocAnalysisOfficeDocSectionAttributeEnum = "number"
	DocAnalysisOfficeDocSectionAttributeEnumFootnote DocAnalysisOfficeDocSectionAttributeEnum = "footnote"
)

// All allowed values of DocAnalysisOfficeDocSectionAttributeEnum enum
var AllowedDocAnalysisOfficeDocSectionAttributeEnumEnumValues = []DocAnalysisOfficeDocSectionAttributeEnum{
	"section",
	"header",
	"footer",
	"number",
	"footnote",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocAnalysisOfficeDocSectionAttributeEnum) IsValid() bool {
	for _, existing := range AllowedDocAnalysisOfficeDocSectionAttributeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
