package nlp

// CommentTagTypeEnum the model 'CommentTagTypeEnum'
type CommentTagTypeEnum int32

// List of CommentTagTypeEnum
const (
	CommentTagTypeEnumValue1  CommentTagTypeEnum = 1
	CommentTagTypeEnumValue2  CommentTagTypeEnum = 2
	CommentTagTypeEnumValue3  CommentTagTypeEnum = 3
	CommentTagTypeEnumValue4  CommentTagTypeEnum = 4
	CommentTagTypeEnumValue5  CommentTagTypeEnum = 5
	CommentTagTypeEnumValue6  CommentTagTypeEnum = 6
	CommentTagTypeEnumValue7  CommentTagTypeEnum = 7
	CommentTagTypeEnumValue8  CommentTagTypeEnum = 8
	CommentTagTypeEnumValue9  CommentTagTypeEnum = 9
	CommentTagTypeEnumValue10 CommentTagTypeEnum = 10
	CommentTagTypeEnumValue11 CommentTagTypeEnum = 11
	CommentTagTypeEnumValue12 CommentTagTypeEnum = 12
	CommentTagTypeEnumValue13 CommentTagTypeEnum = 13
)

// All allowed values of CommentTagTypeEnum enum
var AllowedCommentTagTypeEnumEnumValues = []CommentTagTypeEnum{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommentTagTypeEnum) IsValid() bool {
	for _, existing := range AllowedCommentTagTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
