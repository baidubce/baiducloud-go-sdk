package ocr

// MultiIdcardImageStatusEnum the model 'MultiIdcardImageStatusEnum'
type MultiIdcardImageStatusEnum string

// List of MultiIdcardImageStatusEnum
const (
	MultiIdcardImageStatusEnumNormal        MultiIdcardImageStatusEnum = "normal"
	MultiIdcardImageStatusEnumReversedSide  MultiIdcardImageStatusEnum = "reversed_side"
	MultiIdcardImageStatusEnumNonIdcard     MultiIdcardImageStatusEnum = "non_idcard"
	MultiIdcardImageStatusEnumBlurred       MultiIdcardImageStatusEnum = "blurred"
	MultiIdcardImageStatusEnumOtherTypeCard MultiIdcardImageStatusEnum = "other_type_card"
	MultiIdcardImageStatusEnumOverExposure  MultiIdcardImageStatusEnum = "over_exposure"
	MultiIdcardImageStatusEnumOverDark      MultiIdcardImageStatusEnum = "over_dark"
	MultiIdcardImageStatusEnumUnknown       MultiIdcardImageStatusEnum = "unknown"
)

// All allowed values of MultiIdcardImageStatusEnum enum
var AllowedMultiIdcardImageStatusEnumEnumValues = []MultiIdcardImageStatusEnum{
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
func (v MultiIdcardImageStatusEnum) IsValid() bool {
	for _, existing := range AllowedMultiIdcardImageStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
