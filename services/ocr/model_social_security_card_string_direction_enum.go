package ocr

// SocialSecurityCardStringDirectionEnum the model 'SocialSecurityCardStringDirectionEnum'
type SocialSecurityCardStringDirectionEnum string

// List of SocialSecurityCardStringDirectionEnum
const (
	SocialSecurityCardStringDirectionEnumValueMinus1 SocialSecurityCardStringDirectionEnum = "-1"
	SocialSecurityCardStringDirectionEnumValue0      SocialSecurityCardStringDirectionEnum = "0"
	SocialSecurityCardStringDirectionEnumValue1      SocialSecurityCardStringDirectionEnum = "1"
	SocialSecurityCardStringDirectionEnumValue2      SocialSecurityCardStringDirectionEnum = "2"
	SocialSecurityCardStringDirectionEnumValue3      SocialSecurityCardStringDirectionEnum = "3"
)

// All allowed values of SocialSecurityCardStringDirectionEnum enum
var AllowedSocialSecurityCardStringDirectionEnumEnumValues = []SocialSecurityCardStringDirectionEnum{
	"-1",
	"0",
	"1",
	"2",
	"3",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SocialSecurityCardStringDirectionEnum) IsValid() bool {
	for _, existing := range AllowedSocialSecurityCardStringDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
