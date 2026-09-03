package image

// ImageRetMsgEnum the model 'ImageRetMsgEnum'
type ImageRetMsgEnum string

// List of ImageRetMsgEnum
const (
	ImageRetMsgEnumSuccess    ImageRetMsgEnum = "success"
	ImageRetMsgEnumProcessing ImageRetMsgEnum = "processing"
)

// All allowed values of ImageRetMsgEnum enum
var AllowedImageRetMsgEnumEnumValues = []ImageRetMsgEnum{
	"success",
	"processing",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageRetMsgEnum) IsValid() bool {
	for _, existing := range AllowedImageRetMsgEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
