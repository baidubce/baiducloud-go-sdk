package image

// SelfieAnimeTypeEnum the model 'SelfieAnimeTypeEnum'
type SelfieAnimeTypeEnum string

// List of SelfieAnimeTypeEnum
const (
	SelfieAnimeTypeEnumAnime     SelfieAnimeTypeEnum = "anime"
	SelfieAnimeTypeEnumAnimeMask SelfieAnimeTypeEnum = "anime_mask"
)

// All allowed values of SelfieAnimeTypeEnum enum
var AllowedSelfieAnimeTypeEnumEnumValues = []SelfieAnimeTypeEnum{
	"anime",
	"anime_mask",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SelfieAnimeTypeEnum) IsValid() bool {
	for _, existing := range AllowedSelfieAnimeTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
