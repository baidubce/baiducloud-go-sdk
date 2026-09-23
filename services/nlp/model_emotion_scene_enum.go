package nlp

// EmotionSceneEnum the model 'EmotionSceneEnum'
type EmotionSceneEnum string

// List of EmotionSceneEnum
const (
	EmotionSceneEnumDefault         EmotionSceneEnum = "default"
	EmotionSceneEnumTalk            EmotionSceneEnum = "talk"
	EmotionSceneEnumTask            EmotionSceneEnum = "task"
	EmotionSceneEnumCustomerService EmotionSceneEnum = "customer_service"
)

// All allowed values of EmotionSceneEnum enum
var AllowedEmotionSceneEnumEnumValues = []EmotionSceneEnum{
	"default",
	"talk",
	"task",
	"customer_service",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmotionSceneEnum) IsValid() bool {
	for _, existing := range AllowedEmotionSceneEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
