package ocr

// SceneTypeEnum the model 'SceneTypeEnum'
type SceneTypeEnum string

// List of SceneTypeEnum
const (
	SceneTypeEnumPaper       SceneTypeEnum = "paper"
	SceneTypeEnumAnswerSheet SceneTypeEnum = "answer_sheet"
)

// All allowed values of SceneTypeEnum enum
var AllowedSceneTypeEnumEnumValues = []SceneTypeEnum{
	"paper",
	"answer_sheet",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SceneTypeEnum) IsValid() bool {
	for _, existing := range AllowedSceneTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
