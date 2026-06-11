package bcc

// ImageType the model 'ImageType'
type ImageType string

// List of ImageType
const (
	ImageTypeAll           ImageType = "All"
	ImageTypeSystem        ImageType = "System"
	ImageTypeCustom        ImageType = "Custom"
	ImageTypeIntegration   ImageType = "Integration"
	ImageTypeSharing       ImageType = "Sharing"
	ImageTypeBbcsystem     ImageType = "BbcSystem"
	ImageTypeBbccustom     ImageType = "BbcCustom"
	ImageTypeGpubccsystem  ImageType = "GpuBccSystem"
	ImageTypeGpubcccustom  ImageType = "GpuBccCustom"
	ImageTypeGpubbcsystem  ImageType = "GpuBbcSystem"
	ImageTypeGpubbccustom  ImageType = "GpuBbcCustom"
	ImageTypeEbctotal      ImageType = "EbcTotal"
	ImageTypeEbcsystem     ImageType = "EbcSystem"
	ImageTypeEbccustom     ImageType = "EbcCustom"
	ImageTypeFpgabccsystem ImageType = "FpgaBccSystem"
	ImageTypeFpgabcccustom ImageType = "FpgaBccCustom"
)

// All allowed values of ImageType enum
var AllowedImageTypeEnumValues = []ImageType{
	"All",
	"System",
	"Custom",
	"Integration",
	"Sharing",
	"BbcSystem",
	"BbcCustom",
	"GpuBccSystem",
	"GpuBccCustom",
	"GpuBbcSystem",
	"GpuBbcCustom",
	"EbcTotal",
	"EbcSystem",
	"EbcCustom",
	"FpgaBccSystem",
	"FpgaBccCustom",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageType) IsValid() bool {
	for _, existing := range AllowedImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
