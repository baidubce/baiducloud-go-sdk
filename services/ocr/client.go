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

	CONSTANT_BUSINESS_LICENSE = "business_license"

	CONSTANT_BRAIN = "brain"

	CONSTANT_ONLINE = "online"

	CONSTANT_V2 = "v2"

	CONSTANT_PADDLE_VL_PARSER = "paddle-vl-parser"

	CONSTANT_TASK = "task"

	CONSTANT_QUERY = "query"

	CONSTANT_ACCURATE = "accurate"

	CONSTANT_NUMBERS = "numbers"

	CONSTANT_FACADE = "facade"

	CONSTANT_FORGERY_DETECTION = "forgery_detection"

	CONSTANT_CORRECT_EDU = "correct_edu"

	CONSTANT_CREATE_TASK = "create_task"

	CONSTANT_ONLINE_TAXI_ITINERARY = "online_taxi_itinerary"

	CONSTANT_VEHICLE_REGISTRATION_CERTIFICATE = "vehicle_registration_certificate"

	CONSTANT_LICENSE_PLATE = "license_plate"

	CONSTANT_THREE_FACTORS_VERIFICATION = "three_factors_verification"

	CONSTANT_WAYBILL = "waybill"

	CONSTANT_PASSPORT = "passport"

	CONSTANT_WEBIMAGE = "webimage"

	CONSTANT_GET_RESULT = "get_result"

	CONSTANT_TOLL_INVOICE = "toll_invoice"

	CONSTANT_DOC_CLASSIFY = "doc_classify"

	CONSTANT_VAT_INVOICE = "vat_invoice"

	CONSTANT_DOC_ANALYSIS_OFFICE = "doc_analysis_office"

	CONSTANT_TAXI_RECEIPT = "taxi_receipt"

	CONSTANT_VEHICLE_CERTIFICATE = "vehicle_certificate"

	CONSTANT_METER = "meter"

	CONSTANT_HANDWRITING_COMPOSITION = "handwriting_composition"

	CONSTANT_HANDWRITING = "handwriting"

	CONSTANT_DOC_CROP_ENHANCE = "doc_crop_enhance"

	CONSTANT_MEDICAL_RECORD = "medical_record"

	CONSTANT_BUSINESSLICENSE_VERIFICATION_DETAILED = "businesslicense_verification_detailed"

	CONSTANT_WEIGHT_NOTE = "weight_note"

	CONSTANT_SOCIAL_SECURITY_CARD = "social_security_card"

	CONSTANT_SHOPPING_RECEIPT = "shopping_receipt"

	CONSTANT_MEDICAL_REPORT_DETECTION = "medical_report_detection"

	CONSTANT_FERRY_TICKET = "ferry_ticket"

	CONSTANT_BUSINESSLICENSE_VERIFICATION_STANDARD = "businesslicense_verification_standard"

	CONSTANT_BIRTH_CERTIFICATE = "birth_certificate"

	CONSTANT_REAL_ESTATE_CERTIFICATE = "real_estate_certificate"

	CONSTANT_REMOVE_HANDWRITING = "remove_handwriting"

	CONSTANT_ACCURATE_BASIC = "accurate_basic"

	CONSTANT_HOUSEHOLD_REGISTER = "household_register"

	CONSTANT_OVERSEAS_PASSPORT = "overseas_passport"

	CONSTANT_SMART_STRUCT = "smart_struct"

	CONSTANT_QUOTA_INVOICE = "quota_invoice"

	CONSTANT_MULTIPLE_INVOICE = "multiple_invoice"

	CONSTANT_WEBIMAGE_LOC = "webimage_loc"

	CONSTANT_ROAD_TRANSPORT_CERTIFICATE = "road_transport_certificate"

	CONSTANT_MULTI_IDCARD = "multi_idcard"

	CONSTANT_QRCODE = "qrcode"

	CONSTANT_GENERAL_BASIC = "general_basic"

	CONSTANT_SEAL = "seal"

	CONSTANT_TWO_FACTORS_VERIFICATION = "two_factors_verification"

	CONSTANT_FOUR_FACTORS_VERIFICATION = "four_factors_verification"

	CONSTANT_MEDICAL_DETAIL = "medical_detail"

	CONSTANT_PARSER = "parser"

	CONSTANT_PAPER_CUT_EDU_VLM = "paper_cut_edu_vlm"

	CONSTANT_GENERAL = "general"

	CONSTANT_BANK_RECEIPT_NEW = "bank_receipt_new"

	CONSTANT_VIN_CODE = "vin_code"

	CONSTANT_MEDICAL_STATEMENT = "medical_statement"

	CONSTANT_FOREIGN_RESIDENT_ID_CARD = "foreign_resident_id_card"

	CONSTANT_MIXED_MULTI_VEHICLE = "mixed_multi_vehicle"

	CONSTANT_HK_MACAU_TAIWAN_EXITENTRYPERMIT = "hk_macau_taiwan_exitentrypermit"

	CONSTANT_DOC_ANALYSIS = "doc_analysis"

	CONSTANT_TABLE = "table"

	CONSTANT_VEHICLE_LICENSE = "vehicle_license"

	CONSTANT_MEDICAL_PRESCRIPTION = "medical_prescription"

	CONSTANT_MARRIAGE_CERTIFICATE = "marriage_certificate"

	CONSTANT_MEDICAL_INVOICE = "medical_invoice"

	CONSTANT_USED_VEHICLE_INVOICE = "used_vehicle_invoice"

	CONSTANT_HEALTH_REPORT = "health_report"

	CONSTANT_PAPER_CUT_EDU = "paper_cut_edu"

	CONSTANT_ACCOUNT_OPENING = "account_opening"

	CONSTANT_AIR_TICKET = "air_ticket"

	CONSTANT_TRAIN_TICKET = "train_ticket"

	CONSTANT_MEDICAL_SUMMARY = "medical_summary"

	CONSTANT_VEHICLE_INVOICE = "vehicle_invoice"

	CONSTANT_DRIVING_LICENSE = "driving_license"

	CONSTANT_DIVORCE_CERTIFICATE = "divorce_certificate"

	CONSTANT_IDCARD = "idcard"

	CONSTANT_BUS_TICKET = "bus_ticket"

	CONSTANT_INVOICE = "invoice"

	CONSTANT_BANKCARD = "bankcard"
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

func getAccountOpeningUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ACCOUNT_OPENING
}
func getAccurateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ACCURATE
}
func getAccurateBasicUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ACCURATE_BASIC
}
func getAirTicketUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_AIR_TICKET
}
func getBankReceiptNewUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BANK_RECEIPT_NEW
}
func getBankcardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BANKCARD
}
func getBirthCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BIRTH_CERTIFICATE
}
func getBusTicketUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUS_TICKET
}
func getBusinessLicenseUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUSINESS_LICENSE
}
func getBusinesslicenseDetailedUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUSINESSLICENSE_VERIFICATION_DETAILED
}
func getBusinesslicenseStandardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUSINESSLICENSE_VERIFICATION_STANDARD
}
func getBusinesslicenseVerificationDetailedUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUSINESSLICENSE_VERIFICATION_DETAILED
}
func getBusinesslicenseVerificationStandardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_BUSINESSLICENSE_VERIFICATION_STANDARD
}
func getCorrectEduCreateTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_CORRECT_EDU + bce.URI_PREFIX + CONSTANT_CREATE_TASK
}
func getCorrectEduGetResultUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_CORRECT_EDU + bce.URI_PREFIX + CONSTANT_GET_RESULT
}
func getDivorceCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DIVORCE_CERTIFICATE
}
func getDocAnalysisUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DOC_ANALYSIS
}
func getDocAnalysisOfficeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DOC_ANALYSIS_OFFICE
}
func getDocClassifyUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DOC_CLASSIFY
}
func getDocCropEnhanceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DOC_CROP_ENHANCE
}
func getDrivingLicenseUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DRIVING_LICENSE
}
func getFacadeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FACADE
}
func getFerryTicketUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FERRY_TICKET
}
func getForeignResidentIdCardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FOREIGN_RESIDENT_ID_CARD
}
func getForgeryDetectionUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FORGERY_DETECTION
}
func getFourFactorsVerificationUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FOUR_FACTORS_VERIFICATION
}
func getGeneralUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_GENERAL
}
func getGeneralBasicUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_GENERAL_BASIC
}
func getHandwritingUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HANDWRITING
}
func getHandwritingCompositionCreateTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HANDWRITING_COMPOSITION + bce.URI_PREFIX + CONSTANT_CREATE_TASK
}
func getHandwritingCompositionGetResultUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HANDWRITING_COMPOSITION + bce.URI_PREFIX + CONSTANT_GET_RESULT
}
func getHealthReportUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HEALTH_REPORT
}
func getHkMacauTaiwanExitentrypermitUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HK_MACAU_TAIWAN_EXITENTRYPERMIT
}
func getHkMacauTaiwanpermitUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HK_MACAU_TAIWAN_EXITENTRYPERMIT
}
func getHouseholdRegisterUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_HOUSEHOLD_REGISTER
}
func getIdcardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_IDCARD
}
func getInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INVOICE
}
func getLicensePlateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_LICENSE_PLATE
}
func getMarriageCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MARRIAGE_CERTIFICATE
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
func getMeterUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_METER
}
func getMixedMultiVehicleUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MIXED_MULTI_VEHICLE
}
func getMultiIdcardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MULTI_IDCARD
}
func getMultipleInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MULTIPLE_INVOICE
}
func getNumbersUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_NUMBERS
}
func getOnlineTaxiItineraryUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ONLINE_TAXI_ITINERARY
}
func getOverseasPassportUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_OVERSEAS_PASSPORT
}
func getPaddleVlParserTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_BRAIN + bce.URI_PREFIX + CONSTANT_ONLINE + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_PADDLE_VL_PARSER + bce.URI_PREFIX + CONSTANT_TASK
}
func getPaddleVlParserTaskQueryUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_BRAIN + bce.URI_PREFIX + CONSTANT_ONLINE + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_PADDLE_VL_PARSER + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_QUERY
}
func getPaperCutEduUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PAPER_CUT_EDU
}
func getPaperCutEduVlmCreateTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PAPER_CUT_EDU_VLM + bce.URI_PREFIX + CONSTANT_CREATE_TASK
}
func getPaperCutEduVlmGetResultUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PAPER_CUT_EDU_VLM + bce.URI_PREFIX + CONSTANT_GET_RESULT
}
func getParserTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_BRAIN + bce.URI_PREFIX + CONSTANT_ONLINE + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_PARSER + bce.URI_PREFIX + CONSTANT_TASK
}
func getParserTaskQueryUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_BRAIN + bce.URI_PREFIX + CONSTANT_ONLINE + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_PARSER + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_QUERY
}
func getPassportUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_PASSPORT
}
func getQrcodeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_QRCODE
}
func getQuotaInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_QUOTA_INVOICE
}
func getRealEstateCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_REAL_ESTATE_CERTIFICATE
}
func getRemoveHandwritingUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_REMOVE_HANDWRITING
}
func getRoadTransportCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ROAD_TRANSPORT_CERTIFICATE
}
func getSealUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SEAL
}
func getShoppingReceiptUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SHOPPING_RECEIPT
}
func getSmartStructUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SMART_STRUCT
}
func getSocialSecurityCardUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SOCIAL_SECURITY_CARD
}
func getTableUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TABLE
}
func getTaxiReceiptUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TAXI_RECEIPT
}
func getThreeFactorsVerificationUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_THREE_FACTORS_VERIFICATION
}
func getTollInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TOLL_INVOICE
}
func getTrainTicketUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TRAIN_TICKET
}
func getTwoFactorsVerificationUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TWO_FACTORS_VERIFICATION
}
func getUsedVehicleInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_USED_VEHICLE_INVOICE
}
func getVatInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VAT_INVOICE
}
func getVehicleCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VEHICLE_CERTIFICATE
}
func getVehicleInvoiceUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VEHICLE_INVOICE
}
func getVehicleLicenseUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VEHICLE_LICENSE
}
func getVehicleRegCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VEHICLE_REGISTRATION_CERTIFICATE
}
func getVehicleRegistrationCertificateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VEHICLE_REGISTRATION_CERTIFICATE
}
func getVinCodeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_VIN_CODE
}
func getWaybillUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_WAYBILL
}
func getWebImageUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_WEBIMAGE
}
func getWebImageLocUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_WEBIMAGE_LOC
}
func getWeightNoteUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_OCR + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_WEIGHT_NOTE
}
