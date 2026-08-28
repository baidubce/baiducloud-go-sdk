package ocr

// ExitentrypermitTypeEnum the model 'ExitentrypermitTypeEnum'
type ExitentrypermitTypeEnum string

// List of ExitentrypermitTypeEnum
const (
	ExitentrypermitTypeEnumHkMcPassportFront       ExitentrypermitTypeEnum = "hk_mc_passport_front"
	ExitentrypermitTypeEnumHkMcPassportBack        ExitentrypermitTypeEnum = "hk_mc_passport_back"
	ExitentrypermitTypeEnumTwPassportFront         ExitentrypermitTypeEnum = "tw_passport_front"
	ExitentrypermitTypeEnumTwPassportBack          ExitentrypermitTypeEnum = "tw_passport_back"
	ExitentrypermitTypeEnumTwReturnPassportFront   ExitentrypermitTypeEnum = "tw_return_passport_front"
	ExitentrypermitTypeEnumTwReturnPassportBack    ExitentrypermitTypeEnum = "tw_return_passport_back"
	ExitentrypermitTypeEnumHkMcReturnPassportFront ExitentrypermitTypeEnum = "hk_mc_return_passport_front"
	ExitentrypermitTypeEnumHkMcReturnPassportBack  ExitentrypermitTypeEnum = "hk_mc_return_passport_back"
)

// All allowed values of ExitentrypermitTypeEnum enum
var AllowedExitentrypermitTypeEnumEnumValues = []ExitentrypermitTypeEnum{
	"hk_mc_passport_front",
	"hk_mc_passport_back",
	"tw_passport_front",
	"tw_passport_back",
	"tw_return_passport_front",
	"tw_return_passport_back",
	"hk_mc_return_passport_front",
	"hk_mc_return_passport_back",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExitentrypermitTypeEnum) IsValid() bool {
	for _, existing := range AllowedExitentrypermitTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
