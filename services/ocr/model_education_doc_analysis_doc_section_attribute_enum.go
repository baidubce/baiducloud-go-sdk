package ocr

// EducationDocAnalysisDocSectionAttributeEnum the model 'EducationDocAnalysisDocSectionAttributeEnum'
type EducationDocAnalysisDocSectionAttributeEnum string

// List of EducationDocAnalysisDocSectionAttributeEnum
const (
	EducationDocAnalysisDocSectionAttributeEnumSection  EducationDocAnalysisDocSectionAttributeEnum = "section"
	EducationDocAnalysisDocSectionAttributeEnumHeader   EducationDocAnalysisDocSectionAttributeEnum = "header"
	EducationDocAnalysisDocSectionAttributeEnumFooter   EducationDocAnalysisDocSectionAttributeEnum = "footer"
	EducationDocAnalysisDocSectionAttributeEnumNumber   EducationDocAnalysisDocSectionAttributeEnum = "number"
	EducationDocAnalysisDocSectionAttributeEnumFootnote EducationDocAnalysisDocSectionAttributeEnum = "footnote"
)

// All allowed values of EducationDocAnalysisDocSectionAttributeEnum enum
var AllowedEducationDocAnalysisDocSectionAttributeEnumEnumValues = []EducationDocAnalysisDocSectionAttributeEnum{
	"section",
	"header",
	"footer",
	"number",
	"footnote",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EducationDocAnalysisDocSectionAttributeEnum) IsValid() bool {
	for _, existing := range AllowedEducationDocAnalysisDocSectionAttributeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
