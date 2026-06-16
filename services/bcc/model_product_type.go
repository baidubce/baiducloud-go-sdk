package bcc

// ProductType the model 'ProductType'
type ProductType string

// List of ProductType
const (
	ProductTypePrepaid  ProductType = "Prepaid"
	ProductTypePostpaid ProductType = "Postpaid"
)

// All allowed values of ProductType enum
var AllowedProductTypeEnumValues = []ProductType{
	"Prepaid",
	"Postpaid",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductType) IsValid() bool {
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
