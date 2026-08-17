package ocr

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "ocr." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_REST = "rest"

	CONSTANT_2_0 = "2.0"

	CONSTANT_OCR = "ocr"

	CONSTANT_V1 = "v1"

	CONSTANT_MEDICAL_SUMMARY = "medical_summary"

	CONSTANT_MEDICAL_STATEMENT = "medical_statement"

	CONSTANT_MEDICAL_PRESCRIPTION = "medical_prescription"

	CONSTANT_MEDICAL_INVOICE = "medical_invoice"

	CONSTANT_MEDICAL_RECORD = "medical_record"

	CONSTANT_HEALTH_REPORT = "health_report"

	CONSTANT_MEDICAL_DETAIL = "medical_detail"

	CONSTANT_MEDICAL_REPORT_DETECTION = "medical_report_detection"
)

// Client of ocr service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewClientWithApiKey(apiKey, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithApiKey(apiKey, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewClientWithAccessToken(apiKey, secretKey, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAccessToken(apiKey, secretKey, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getHealthReportUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HEALTH_REPORT
}
func getMedicalDetailUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_DETAIL
}
func getMedicalInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_INVOICE
}
func getMedicalPrescriptionUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_PRESCRIPTION
}
func getMedicalRecordUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_RECORD
}
func getMedicalReportDetectionUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_REPORT_DETECTION
}
func getMedicalStatementUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_STATEMENT
}
func getMedicalSummaryUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MEDICAL_SUMMARY
}
