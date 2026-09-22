package face

// FaceTypeEnum the model 'FaceTypeEnum'
type FaceTypeEnum string

// List of FaceTypeEnum
const (
	FaceTypeEnumLive      FaceTypeEnum = "LIVE"
	FaceTypeEnumIdcard    FaceTypeEnum = "IDCARD"
	FaceTypeEnumWatermark FaceTypeEnum = "WATERMARK"
	FaceTypeEnumCert      FaceTypeEnum = "CERT"
	FaceTypeEnumInfrared  FaceTypeEnum = "INFRARED"
	FaceTypeEnumHybrid    FaceTypeEnum = "HYBRID"
)

// All allowed values of FaceTypeEnum enum
var AllowedFaceTypeEnumEnumValues = []FaceTypeEnum{
	"LIVE",
	"IDCARD",
	"WATERMARK",
	"CERT",
	"INFRARED",
	"HYBRID",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FaceTypeEnum) IsValid() bool {
	for _, existing := range AllowedFaceTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
