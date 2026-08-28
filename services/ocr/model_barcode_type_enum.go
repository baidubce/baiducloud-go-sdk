package ocr

// BarcodeTypeEnum the model 'BarcodeTypeEnum'
type BarcodeTypeEnum string

// List of BarcodeTypeEnum
const (
	BarcodeTypeEnumUpcA       BarcodeTypeEnum = "UPC_A"
	BarcodeTypeEnumUpcE       BarcodeTypeEnum = "UPC_E"
	BarcodeTypeEnumEan13      BarcodeTypeEnum = "EAN_13"
	BarcodeTypeEnumEan8       BarcodeTypeEnum = "EAN_8"
	BarcodeTypeEnumCode39     BarcodeTypeEnum = "CODE_39"
	BarcodeTypeEnumCode93     BarcodeTypeEnum = "CODE_93"
	BarcodeTypeEnumCode128    BarcodeTypeEnum = "CODE_128"
	BarcodeTypeEnumItf        BarcodeTypeEnum = "ITF"
	BarcodeTypeEnumCodabar    BarcodeTypeEnum = "CODABAR"
	BarcodeTypeEnumQrCode     BarcodeTypeEnum = "QR_CODE"
	BarcodeTypeEnumDataMatrix BarcodeTypeEnum = "DATA_MATRIX"
	BarcodeTypeEnumAztec      BarcodeTypeEnum = "AZTEC"
	BarcodeTypeEnumPdf417     BarcodeTypeEnum = "PDF_417"
)

// All allowed values of BarcodeTypeEnum enum
var AllowedBarcodeTypeEnumEnumValues = []BarcodeTypeEnum{
	"UPC_A",
	"UPC_E",
	"EAN_13",
	"EAN_8",
	"CODE_39",
	"CODE_93",
	"CODE_128",
	"ITF",
	"CODABAR",
	"QR_CODE",
	"DATA_MATRIX",
	"AZTEC",
	"PDF_417",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BarcodeTypeEnum) IsValid() bool {
	for _, existing := range AllowedBarcodeTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
