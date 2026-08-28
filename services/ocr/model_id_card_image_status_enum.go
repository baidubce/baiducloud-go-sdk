package ocr

// IdCardImageStatusEnum the model 'IdCardImageStatusEnum'
type IdCardImageStatusEnum string

// List of IdCardImageStatusEnum
const (
	IdCardImageStatusEnumNormal        IdCardImageStatusEnum = "normal"
	IdCardImageStatusEnumReversedSide  IdCardImageStatusEnum = "reversed_side"
	IdCardImageStatusEnumNonIdcard     IdCardImageStatusEnum = "non_idcard"
	IdCardImageStatusEnumBlurred       IdCardImageStatusEnum = "blurred"
	IdCardImageStatusEnumOtherTypeCard IdCardImageStatusEnum = "other_type_card"
	IdCardImageStatusEnumOverExposure  IdCardImageStatusEnum = "over_exposure"
	IdCardImageStatusEnumOverDark      IdCardImageStatusEnum = "over_dark"
	IdCardImageStatusEnumUnknown       IdCardImageStatusEnum = "unknown"
)

// All allowed values of IdCardImageStatusEnum enum
var AllowedIdCardImageStatusEnumEnumValues = []IdCardImageStatusEnum{
	"normal",
	"reversed_side",
	"non_idcard",
	"blurred",
	"other_type_card",
	"over_exposure",
	"over_dark",
	"unknown",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IdCardImageStatusEnum) IsValid() bool {
	for _, existing := range AllowedIdCardImageStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
