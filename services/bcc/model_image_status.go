package bcc

// ImageStatus the model 'ImageStatus'
type ImageStatus string

// List of ImageStatus
const (
	ImageStatusCreating      ImageStatus = "Creating"
	ImageStatusCreatedfailed ImageStatus = "CreatedFailed"
	ImageStatusAvailable     ImageStatus = "Available"
	ImageStatusNotavailable  ImageStatus = "NotAvailable"
	ImageStatusError         ImageStatus = "Error"
)

// All allowed values of ImageStatus enum
var AllowedImageStatusEnumValues = []ImageStatus{
	"Creating",
	"CreatedFailed",
	"Available",
	"NotAvailable",
	"Error",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageStatus) IsValid() bool {
	for _, existing := range AllowedImageStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
