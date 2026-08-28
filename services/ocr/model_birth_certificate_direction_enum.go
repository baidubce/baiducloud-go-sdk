package ocr

// BirthCertificateDirectionEnum the model 'BirthCertificateDirectionEnum'
type BirthCertificateDirectionEnum int32

// List of BirthCertificateDirectionEnum
const (
	BirthCertificateDirectionEnumValueMinus1 BirthCertificateDirectionEnum = -1
	BirthCertificateDirectionEnumValue0      BirthCertificateDirectionEnum = 0
	BirthCertificateDirectionEnumValue1      BirthCertificateDirectionEnum = 1
	BirthCertificateDirectionEnumValue2      BirthCertificateDirectionEnum = 2
	BirthCertificateDirectionEnumValue3      BirthCertificateDirectionEnum = 3
)

// All allowed values of BirthCertificateDirectionEnum enum
var AllowedBirthCertificateDirectionEnumEnumValues = []BirthCertificateDirectionEnum{
	-1,
	0,
	1,
	2,
	3,
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BirthCertificateDirectionEnum) IsValid() bool {
	for _, existing := range AllowedBirthCertificateDirectionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
