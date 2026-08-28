package ocr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	OCR_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK        string
	SK        string
	Endpoint  string
	ApiKey    string
	SecretKey string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	// ==== AK/SK 鉴权 ====
	// OCR_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// OCR_CLIENT, _ = NewClientWithAccessToken(confObj.ApiKey, confObj.SecretKey, confObj.Endpoint)

	// ==== API Key 鉴权 ====
	OCR_CLIENT, _ = NewClientWithApiKey(confObj.ApiKey, confObj.Endpoint)

	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_AccountOpening(t *testing.T) {
	accountOpeningRequest := &AccountOpeningRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &AccountOpeningResponse{}
	result, err := OCR_CLIENT.AccountOpening(accountOpeningRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Accurate(t *testing.T) {
	accurateRequest := &AccurateRequest{
		Image:                     util.PtrString(""),
		Url:                       util.PtrString(""),
		PdfFile:                   util.PtrString(""),
		PdfFileNum:                util.PtrInt32(int32(0)),
		OfdFile:                   util.PtrString(""),
		OfdFileNum:                util.PtrInt32(int32(0)),
		LanguageType:              util.PtrString(""),
		EngGranularity:            util.PtrString(""),
		RecognizeGranularity:      util.PtrString(""),
		DetectDirection:           util.PtrBool(false),
		VertexesLocation:          util.PtrBool(false),
		Paragraph:                 util.PtrBool(false),
		Probability:               util.PtrBool(false),
		CharProbability:           util.PtrBool(false),
		MultidirectionalRecognize: util.PtrBool(false),
	}
	result := &AccurateResponse{}
	result, err := OCR_CLIENT.Accurate(accurateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AccurateBasic(t *testing.T) {
	accurateBasicRequest := &AccurateBasicRequest{
		Image:                     util.PtrString(""),
		Url:                       util.PtrString(""),
		PdfFile:                   util.PtrString(""),
		PdfFileNum:                util.PtrInt32(int32(0)),
		OfdFile:                   util.PtrString(""),
		OfdFileNum:                util.PtrInt32(int32(0)),
		LanguageType:              util.PtrString(""),
		DetectDirection:           util.PtrBool(false),
		Paragraph:                 util.PtrBool(false),
		Probability:               util.PtrBool(false),
		MultidirectionalRecognize: util.PtrBool(false),
	}
	result := &AccurateBasicResponse{}
	result, err := OCR_CLIENT.AccurateBasic(accurateBasicRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_AirTicket(t *testing.T) {
	airTicketRequest := &AirTicketRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		OfdFile:     util.PtrString(""),
		OfdFileNum:  util.PtrInt32(int32(0)),
		MultiDetect: util.PtrBool(false),
	}
	result := &AirTicketResponse{}
	result, err := OCR_CLIENT.AirTicket(airTicketRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BankReceiptNew(t *testing.T) {
	bankReceiptNewRequest := &BankReceiptNewRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		OfdFile:     util.PtrString(""),
		OfdFileNum:  util.PtrInt32(int32(0)),
		Probability: util.PtrBool(false),
		Location:    util.PtrBool(false),
	}
	result := &BankReceiptNewResponse{}
	result, err := OCR_CLIENT.BankReceiptNew(bankReceiptNewRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Bankcard(t *testing.T) {
	bankcardRequest := &BankcardRequest{
		Image:         util.PtrString(""),
		Url:           util.PtrString(""),
		Location:      util.PtrBool(false),
		DetectQuality: util.PtrBool(false),
	}
	result := &BankcardResponse{}
	result, err := OCR_CLIENT.Bankcard(bankcardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BirthCertificate(t *testing.T) {
	birthCertificateRequest := &BirthCertificateRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &BirthCertificateResponse{}
	result, err := OCR_CLIENT.BirthCertificate(birthCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusTicket(t *testing.T) {
	busTicketRequest := &BusTicketRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &BusTicketResponse{}
	result, err := OCR_CLIENT.BusTicket(busTicketRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusinessLicense(t *testing.T) {
	businessLicenseRequest := &BusinessLicenseRequest{
		Image:          util.PtrString(""),
		Url:            util.PtrString(""),
		Accuracy:       util.PtrString(""),
		RiskWarn:       util.PtrBool(false),
		DetectQuality:  util.PtrBool(false),
		FullwidthShift: util.PtrBool(false),
	}
	result := &BusinessLicenseResponse{}
	result, err := OCR_CLIENT.BusinessLicense(businessLicenseRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusinesslicenseDetailed(t *testing.T) {
	businesslicenseDetailedRequest := &BusinesslicenseDetailedRequest{
		Verifynum: util.PtrString(""),
	}
	result := &BusinesslicenseDetailedResponse{}
	result, err := OCR_CLIENT.BusinesslicenseDetailed(businesslicenseDetailedRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusinesslicenseStandard(t *testing.T) {
	businesslicenseStandardRequest := &BusinesslicenseStandardRequest{
		Verifynum: util.PtrString(""),
	}
	result := &BusinesslicenseStandardResponse{}
	result, err := OCR_CLIENT.BusinesslicenseStandard(businesslicenseStandardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusinesslicenseVerificationDetailed(t *testing.T) {
	businesslicenseVerificationDetailedRequest := &BusinesslicenseVerificationDetailedRequest{
		Verifynum: util.PtrString(""),
	}
	result := &BusinesslicenseVerificationDetailedResponse{}
	result, err := OCR_CLIENT.BusinesslicenseVerificationDetailed(businesslicenseVerificationDetailedRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BusinesslicenseVerificationStandard(t *testing.T) {
	businesslicenseVerificationStandardRequest := &BusinesslicenseVerificationStandardRequest{
		Verifynum: util.PtrString(""),
	}
	result := &BusinesslicenseVerificationStandardResponse{}
	result, err := OCR_CLIENT.BusinesslicenseVerificationStandard(businesslicenseVerificationStandardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CorrectEduCreateTask(t *testing.T) {
	correctEduCreateTaskRequest := &CorrectEduCreateTaskRequest{
		Image:             util.PtrString(""),
		Url:               util.PtrString(""),
		OnlySplit:         util.PtrBool(false),
		DisablePreprocess: util.PtrBool(false),
	}
	result := &CorrectEduCreateTaskResponse{}
	result, err := OCR_CLIENT.CorrectEduCreateTask(correctEduCreateTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CorrectEduGetResult(t *testing.T) {
	correctEduGetResultRequest := &CorrectEduGetResultRequest{
		TaskId: util.PtrString(""),
	}
	result := &CorrectEduGetResultResponse{}
	result, err := OCR_CLIENT.CorrectEduGetResult(correctEduGetResultRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DivorceCertificate(t *testing.T) {
	divorceCertificateRequest := &DivorceCertificateRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &DivorceCertificateResponse{}
	result, err := OCR_CLIENT.DivorceCertificate(divorceCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DocAnalysis(t *testing.T) {
	docAnalysisRequest := &DocAnalysisRequest{
		Image:                 util.PtrString(""),
		Url:                   util.PtrString(""),
		PdfFile:               util.PtrString(""),
		PdfFileNum:            util.PtrInt32(int32(0)),
		LanguageType:          util.PtrString(""),
		ResultType:            util.PtrString(""),
		DetectDirection:       util.PtrBool(false),
		LineProbability:       util.PtrBool(false),
		DispLinePoly:          util.PtrBool(false),
		WordsType:             util.PtrString(""),
		LayoutAnalysis:        util.PtrBool(false),
		RecgFormula:           util.PtrBool(false),
		RecgLongDivision:      util.PtrBool(false),
		DispUnderlineAnalysis: util.PtrBool(false),
		RecgAlter:             util.PtrBool(false),
	}
	result := &DocAnalysisResponse{}
	result, err := OCR_CLIENT.DocAnalysis(docAnalysisRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DocAnalysisOffice(t *testing.T) {
	docAnalysisOfficeRequest := &DocAnalysisOfficeRequest{
		Image:                 util.PtrString(""),
		Url:                   util.PtrString(""),
		PdfFile:               util.PtrString(""),
		PdfFileNum:            util.PtrInt32(int32(0)),
		OfdFile:               util.PtrString(""),
		OfdFileNum:            util.PtrInt32(int32(0)),
		LanguageType:          util.PtrString(""),
		ResultType:            util.PtrString(""),
		CharProbability:       util.PtrBool(false),
		DetectDirection:       util.PtrBool(false),
		LineProbability:       util.PtrBool(false),
		DispLinePoly:          util.PtrBool(false),
		WordsType:             util.PtrString(""),
		LayoutAnalysis:        util.PtrBool(false),
		RecgTables:            util.PtrBool(false),
		RecogSeal:             util.PtrBool(false),
		RecgFormula:           util.PtrBool(false),
		EraseSeal:             util.PtrBool(false),
		DispUnderlineAnalysis: util.PtrBool(false),
	}
	result := &DocAnalysisOfficeResponse{}
	result, err := OCR_CLIENT.DocAnalysisOffice(docAnalysisOfficeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DocClassify(t *testing.T) {
	docClassifyRequest := &DocClassifyRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &DocClassifyResponse{}
	result, err := OCR_CLIENT.DocClassify(docClassifyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DocCropEnhance(t *testing.T) {
	docCropEnhanceRequest := &DocCropEnhanceRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		ScanType:    util.PtrInt32(int32(0)),
		Points:      util.PtrString(""),
		EnhanceType: util.PtrInt32(int32(0)),
	}
	result := &DocCropEnhanceResponse{}
	result, err := OCR_CLIENT.DocCropEnhance(docCropEnhanceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DrivingLicense(t *testing.T) {
	drivingLicenseRequest := &DrivingLicenseRequest{
		Image:              util.PtrString(""),
		Url:                util.PtrString(""),
		DetectDirection:    util.PtrBool(false),
		DrivingLicenseSide: util.PtrString(""),
		UnifiedValidPeriod: util.PtrBool(false),
		QualityWarn:        util.PtrBool(false),
		RiskWarn:           util.PtrBool(false),
	}
	result := &DrivingLicenseResponse{}
	result, err := OCR_CLIENT.DrivingLicense(drivingLicenseRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Facade(t *testing.T) {
	facadeRequest := &FacadeRequest{
		Image: util.PtrString(""),
	}
	result := &FacadeResponse{}
	result, err := OCR_CLIENT.Facade(facadeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FerryTicket(t *testing.T) {
	ferryTicketRequest := &FerryTicketRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &FerryTicketResponse{}
	result, err := OCR_CLIENT.FerryTicket(ferryTicketRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ForeignResidentIdCard(t *testing.T) {
	foreignResidentIdCardRequest := &ForeignResidentIdCardRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &ForeignResidentIdCardResponse{}
	result, err := OCR_CLIENT.ForeignResidentIdCard(foreignResidentIdCardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ForgeryDetection(t *testing.T) {
	forgeryDetectionRequest := &ForgeryDetectionRequest{
		Image:               util.PtrString(""),
		Url:                 util.PtrString(""),
		DetectProportion:    util.PtrBool(false),
		DetectThreshold:     util.PtrFloat64(float64(0)),
		ReturnHeatmap:       util.PtrBool(false),
		RestrictProbability: util.PtrFloat64(float64(0)),
	}
	result := &ForgeryDetectionResponse{}
	result, err := OCR_CLIENT.ForgeryDetection(forgeryDetectionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FourFactorsVerification(t *testing.T) {
	fourFactorsVerificationRequest := &FourFactorsVerificationRequest{
		Name:    util.PtrString(""),
		Idcard:  util.PtrString(""),
		Company: util.PtrString(""),
		Regnum:  util.PtrString(""),
	}
	result := &FourFactorsVerificationResponse{}
	result, err := OCR_CLIENT.FourFactorsVerification(fourFactorsVerificationRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_General(t *testing.T) {
	generalRequest := &GeneralRequest{
		Image:                util.PtrString(""),
		Url:                  util.PtrString(""),
		PdfFile:              util.PtrString(""),
		PdfFileNum:           util.PtrInt32(int32(0)),
		OfdFile:              util.PtrString(""),
		OfdFileNum:           util.PtrInt32(int32(0)),
		RecognizeGranularity: util.PtrString(""),
		LanguageType:         util.PtrString(""),
		DetectDirection:      util.PtrBool(false),
		DetectLanguage:       util.PtrBool(false),
		Paragraph:            util.PtrBool(false),
		VertexesLocation:     util.PtrBool(false),
		Probability:          util.PtrBool(false),
	}
	result := &GeneralResponse{}
	result, err := OCR_CLIENT.General(generalRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GeneralBasic(t *testing.T) {
	generalBasicRequest := &GeneralBasicRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		PdfFile:         util.PtrString(""),
		PdfFileNum:      util.PtrInt32(int32(0)),
		OfdFile:         util.PtrString(""),
		OfdFileNum:      util.PtrInt32(int32(0)),
		LanguageType:    util.PtrString(""),
		DetectDirection: util.PtrBool(false),
		DetectLanguage:  util.PtrBool(false),
		Paragraph:       util.PtrBool(false),
		Probability:     util.PtrBool(false),
	}
	result := &GeneralBasicResponse{}
	result, err := OCR_CLIENT.GeneralBasic(generalBasicRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Handwriting(t *testing.T) {
	handwritingRequest := &HandwritingRequest{
		Image:                util.PtrString(""),
		Url:                  util.PtrString(""),
		PdfFile:              util.PtrString(""),
		PdfFileNum:           util.PtrInt32(int32(0)),
		OfdFile:              util.PtrString(""),
		OfdFileNum:           util.PtrInt32(int32(0)),
		RecognizeGranularity: util.PtrString(""),
		EngGranularity:       util.PtrString(""),
		Probability:          util.PtrBool(false),
		DetectDirection:      util.PtrBool(false),
		DetectAlteration:     util.PtrBool(false),
		LanguageType:         util.PtrString(""),
	}
	result := &HandwritingResponse{}
	result, err := OCR_CLIENT.Handwriting(handwritingRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HandwritingCompositionCreateTask(t *testing.T) {
	handwritingCompositionCreateTaskRequest := &HandwritingCompositionCreateTaskRequest{
		Image:                util.PtrString(""),
		Url:                  util.PtrString(""),
		PdfFile:              util.PtrString(""),
		RecognizeGranularity: util.PtrString(""),
		PdfFileNum:           util.PtrInt32(int32(0)),
	}
	result := &HandwritingCompositionCreateTaskResponse{}
	result, err := OCR_CLIENT.HandwritingCompositionCreateTask(handwritingCompositionCreateTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HandwritingCompositionGetResult(t *testing.T) {
	handwritingCompositionGetResultRequest := &HandwritingCompositionGetResultRequest{
		TaskId: util.PtrString(""),
	}
	result := &HandwritingCompositionGetResultResponse{}
	result, err := OCR_CLIENT.HandwritingCompositionGetResult(handwritingCompositionGetResultRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HealthReport(t *testing.T) {
	healthReportRequest := &HealthReportRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &HealthReportResponse{}
	result, err := OCR_CLIENT.HealthReport(healthReportRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HkMacauTaiwanExitentrypermit(t *testing.T) {
	hkMacauTaiwanExitentrypermitRequest := &HkMacauTaiwanExitentrypermitRequest{
		ExitentrypermitType: util.PtrString(""),
		Image:               util.PtrString(""),
		Url:                 util.PtrString(""),
		PdfFile:             util.PtrString(""),
		PdfFileNum:          util.PtrInt32(int32(0)),
		Probability:         util.PtrBool(false),
		Location:            util.PtrBool(false),
	}
	result := &HkMacauTaiwanExitentrypermitResponse{}
	result, err := OCR_CLIENT.HkMacauTaiwanExitentrypermit(hkMacauTaiwanExitentrypermitRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HkMacauTaiwanpermit(t *testing.T) {
	hkMacauTaiwanpermitRequest := &HkMacauTaiwanpermitRequest{
		ExitentrypermitType: util.PtrString(""),
		Image:               util.PtrString(""),
		Url:                 util.PtrString(""),
		PdfFile:             util.PtrString(""),
		PdfFileNum:          util.PtrInt32(int32(0)),
		Probability:         util.PtrBool(false),
		Location:            util.PtrBool(false),
	}
	result := &HkMacauTaiwanpermitResponse{}
	result, err := OCR_CLIENT.HkMacauTaiwanpermit(hkMacauTaiwanpermitRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_HouseholdRegister(t *testing.T) {
	householdRegisterRequest := &HouseholdRegisterRequest{
		Image:                 util.PtrString(""),
		Url:                   util.PtrString(""),
		HouseholdRegisterSide: util.PtrString(""),
	}
	result := &HouseholdRegisterResponse{}
	result, err := OCR_CLIENT.HouseholdRegister(householdRegisterRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Idcard(t *testing.T) {
	idcardRequest := &IdcardRequest{
		IdCardSide:       util.PtrString(""),
		Image:            util.PtrString(""),
		Url:              util.PtrString(""),
		DetectPs:         util.PtrBool(false),
		DetectRisk:       util.PtrBool(false),
		DetectQuality:    util.PtrBool(false),
		DetectPhoto:      util.PtrBool(false),
		DetectCard:       util.PtrBool(false),
		DetectDirection:  util.PtrBool(false),
		DetectScreenshot: util.PtrBool(false),
	}
	result := &IdcardResponse{}
	result, err := OCR_CLIENT.Idcard(idcardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Invoice(t *testing.T) {
	invoiceRequest := &InvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &InvoiceResponse{}
	result, err := OCR_CLIENT.Invoice(invoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_LicensePlate(t *testing.T) {
	licensePlateRequest := &LicensePlateRequest{
		Image:          util.PtrString(""),
		Url:            util.PtrString(""),
		MultiDetect:    util.PtrBool(false),
		MultiScale:     util.PtrBool(false),
		DetectComplete: util.PtrBool(false),
		DetectRisk:     util.PtrBool(false),
	}
	result := &LicensePlateResponse{}
	result, err := OCR_CLIENT.LicensePlate(licensePlateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MarriageCertificate(t *testing.T) {
	marriageCertificateRequest := &MarriageCertificateRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		Probability: util.PtrBool(false),
		Location:    util.PtrBool(false),
	}
	result := &MarriageCertificateResponse{}
	result, err := OCR_CLIENT.MarriageCertificate(marriageCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalDetail(t *testing.T) {
	medicalDetailRequest := &MedicalDetailRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalDetailResponse{}
	result, err := OCR_CLIENT.MedicalDetail(medicalDetailRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalInvoice(t *testing.T) {
	medicalInvoiceRequest := &MedicalInvoiceRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
		MediQuery:   util.PtrString(""),
	}
	result := &MedicalInvoiceResponse{}
	result, err := OCR_CLIENT.MedicalInvoice(medicalInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalPrescription(t *testing.T) {
	medicalPrescriptionRequest := &MedicalPrescriptionRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalPrescriptionResponse{}
	result, err := OCR_CLIENT.MedicalPrescription(medicalPrescriptionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalRecord(t *testing.T) {
	medicalRecordRequest := &MedicalRecordRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalRecordResponse{}
	result, err := OCR_CLIENT.MedicalRecord(medicalRecordRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalReportDetection(t *testing.T) {
	medicalReportDetectionRequest := &MedicalReportDetectionRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalReportDetectionResponse{}
	result, err := OCR_CLIENT.MedicalReportDetection(medicalReportDetectionRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalStatement(t *testing.T) {
	medicalStatementRequest := &MedicalStatementRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalStatementResponse{}
	result, err := OCR_CLIENT.MedicalStatement(medicalStatementRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MedicalSummary(t *testing.T) {
	medicalSummaryRequest := &MedicalSummaryRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		Location:    util.PtrBool(false),
		Probability: util.PtrBool(false),
	}
	result := &MedicalSummaryResponse{}
	result, err := OCR_CLIENT.MedicalSummary(medicalSummaryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Meter(t *testing.T) {
	meterRequest := &MeterRequest{
		Image:        util.PtrString(""),
		Url:          util.PtrString(""),
		Probability:  util.PtrBool(false),
		PolyLocation: util.PtrBool(false),
	}
	result := &MeterResponse{}
	result, err := OCR_CLIENT.Meter(meterRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MixedMultiVehicle(t *testing.T) {
	mixedMultiVehicleRequest := &MixedMultiVehicleRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		DetectDirection: util.PtrBool(false),
		Unified:         util.PtrBool(false),
	}
	result := &MixedMultiVehicleResponse{}
	result, err := OCR_CLIENT.MixedMultiVehicle(mixedMultiVehicleRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MultiIdcard(t *testing.T) {
	multiIdcardRequest := &MultiIdcardRequest{
		Image:            util.PtrString(""),
		Url:              util.PtrString(""),
		DetectRisk:       util.PtrBool(false),
		DetectQuality:    util.PtrBool(false),
		DetectPhoto:      util.PtrBool(false),
		DetectCard:       util.PtrBool(false),
		DetectScreenshot: util.PtrBool(false),
	}
	result := &MultiIdcardResponse{}
	result, err := OCR_CLIENT.MultiIdcard(multiIdcardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_MultipleInvoice(t *testing.T) {
	multipleInvoiceRequest := &MultipleInvoiceRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		PdfFile:         util.PtrString(""),
		PdfFileNum:      util.PtrInt32(int32(0)),
		OfdFile:         util.PtrString(""),
		OfdFileNum:      util.PtrInt32(int32(0)),
		VerifyParameter: util.PtrBool(false),
		Probability:     util.PtrBool(false),
		Location:        util.PtrBool(false),
	}
	result := &MultipleInvoiceResponse{}
	result, err := OCR_CLIENT.MultipleInvoice(multipleInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Numbers(t *testing.T) {
	numbersRequest := &NumbersRequest{
		Image:                util.PtrString(""),
		Url:                  util.PtrString(""),
		PdfFile:              util.PtrString(""),
		PdfFileNum:           util.PtrInt32(int32(0)),
		OfdFile:              util.PtrString(""),
		OfdFileNum:           util.PtrInt32(int32(0)),
		RecognizeGranularity: util.PtrString(""),
		DetectDirection:      util.PtrBool(false),
	}
	result := &NumbersResponse{}
	result, err := OCR_CLIENT.Numbers(numbersRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_OnlineTaxiItinerary(t *testing.T) {
	onlineTaxiItineraryRequest := &OnlineTaxiItineraryRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &OnlineTaxiItineraryResponse{}
	result, err := OCR_CLIENT.OnlineTaxiItinerary(onlineTaxiItineraryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_OverseasPassport(t *testing.T) {
	overseasPassportRequest := &OverseasPassportRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &OverseasPassportResponse{}
	result, err := OCR_CLIENT.OverseasPassport(overseasPassportRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PaddleVlParserTask(t *testing.T) {
	paddleVlParserTaskRequest := &PaddleVlParserTaskRequest{
		FileName:        util.PtrString(""),
		FileData:        util.PtrString(""),
		FileUrl:         util.PtrString(""),
		AnalysisChart:   util.PtrBool(false),
		MergeTables:     util.PtrBool(false),
		RelevelTitles:   util.PtrBool(false),
		RecognizeSeal:   util.PtrBool(false),
		ReturnSpanBoxes: util.PtrBool(false),
	}
	result := &PaddleVlParserTaskResponse{}
	result, err := OCR_CLIENT.PaddleVlParserTask(paddleVlParserTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PaddleVlParserTaskQuery(t *testing.T) {
	paddleVlParserTaskQueryRequest := &PaddleVlParserTaskQueryRequest{
		TaskId: util.PtrString(""),
	}
	result := &PaddleVlParserTaskQueryResponse{}
	result, err := OCR_CLIENT.PaddleVlParserTaskQuery(paddleVlParserTaskQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PaperCutEdu(t *testing.T) {
	paperCutEduRequest := &PaperCutEduRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		PdfFile:         util.PtrString(""),
		PdfFileNum:      util.PtrInt32(int32(0)),
		LanguageType:    util.PtrString(""),
		DetectDirection: util.PtrBool(false),
		WordsType:       util.PtrString(""),
		SpliceText:      util.PtrBool(false),
		Enhance:         util.PtrBool(false),
		OnlySplit:       util.PtrBool(false),
	}
	result := &PaperCutEduResponse{}
	result, err := OCR_CLIENT.PaperCutEdu(paperCutEduRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PaperCutEduVlmCreateTask(t *testing.T) {
	paperCutEduVlmCreateTaskRequest := &PaperCutEduVlmCreateTaskRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OnlySplit:  util.PtrBool(false),
		SceneType:  util.PtrString(""),
		Enhance:    util.PtrBool(false),
	}
	result := &PaperCutEduVlmCreateTaskResponse{}
	result, err := OCR_CLIENT.PaperCutEduVlmCreateTask(paperCutEduVlmCreateTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PaperCutEduVlmGetResult(t *testing.T) {
	paperCutEduVlmGetResultRequest := &PaperCutEduVlmGetResultRequest{
		TaskId: util.PtrString(""),
	}
	result := &PaperCutEduVlmGetResultResponse{}
	result, err := OCR_CLIENT.PaperCutEduVlmGetResult(paperCutEduVlmGetResultRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ParserTask(t *testing.T) {
	parserTaskRequest := &ParserTaskRequest{
		FileName:           util.PtrString(""),
		FileData:           util.PtrString(""),
		FileUrl:            util.PtrString(""),
		RecognizeFormula:   util.PtrBool(false),
		AnalysisChart:      util.PtrBool(false),
		AngleAdjust:        util.PtrBool(false),
		ParseImageLayout:   util.PtrBool(false),
		LanguageType:       util.PtrString(""),
		SwitchDigitalWidth: util.PtrString(""),
		HtmlTableFormat:    util.PtrBool(false),
		ReturnDocChunks:    util.PtrString(""),
	}
	result := &ParserTaskResponse{}
	result, err := OCR_CLIENT.ParserTask(parserTaskRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ParserTaskQuery(t *testing.T) {
	parserTaskQueryRequest := &ParserTaskQueryRequest{
		TaskId: util.PtrString(""),
	}
	result := &ParserTaskQueryResponse{}
	result, err := OCR_CLIENT.ParserTaskQuery(parserTaskQueryRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Passport(t *testing.T) {
	passportRequest := &PassportRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &PassportResponse{}
	result, err := OCR_CLIENT.Passport(passportRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Qrcode(t *testing.T) {
	qrcodeRequest := &QrcodeRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
		Location:   util.PtrBool(false),
	}
	result := &QrcodeResponse{}
	result, err := OCR_CLIENT.Qrcode(qrcodeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuotaInvoice(t *testing.T) {
	quotaInvoiceRequest := &QuotaInvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &QuotaInvoiceResponse{}
	result, err := OCR_CLIENT.QuotaInvoice(quotaInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RealEstateCertificate(t *testing.T) {
	realEstateCertificateRequest := &RealEstateCertificateRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		Probability: util.PtrBool(false),
		Location:    util.PtrBool(false),
	}
	result := &RealEstateCertificateResponse{}
	result, err := OCR_CLIENT.RealEstateCertificate(realEstateCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RemoveHandwriting(t *testing.T) {
	removeHandwritingRequest := &RemoveHandwritingRequest{
		Image:        util.PtrString(""),
		Url:          util.PtrString(""),
		PdfFile:      util.PtrString(""),
		PdfFileNum:   util.PtrInt32(int32(0)),
		EnableDetect: util.PtrBool(false),
	}
	result := &RemoveHandwritingResponse{}
	result, err := OCR_CLIENT.RemoveHandwriting(removeHandwritingRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RoadTransportCertificate(t *testing.T) {
	roadTransportCertificateRequest := &RoadTransportCertificateRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &RoadTransportCertificateResponse{}
	result, err := OCR_CLIENT.RoadTransportCertificate(roadTransportCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Seal(t *testing.T) {
	sealRequest := &SealRequest{
		Image:        util.PtrString(""),
		Url:          util.PtrString(""),
		PdfFile:      util.PtrString(""),
		PdfFileNum:   util.PtrInt32(int32(0)),
		OfdFile:      util.PtrString(""),
		OfdFileNum:   util.PtrInt32(int32(0)),
		ReturnImage:  util.PtrBool(false),
		FlattenImage: util.PtrBool(false),
	}
	result := &SealResponse{}
	result, err := OCR_CLIENT.Seal(sealRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ShoppingReceipt(t *testing.T) {
	shoppingReceiptRequest := &ShoppingReceiptRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		OfdFile:     util.PtrString(""),
		OfdFileNum:  util.PtrInt32(int32(0)),
		Probability: util.PtrBool(false),
		Location:    util.PtrBool(false),
	}
	result := &ShoppingReceiptResponse{}
	result, err := OCR_CLIENT.ShoppingReceipt(shoppingReceiptRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SmartStruct(t *testing.T) {
	smartStructRequest := &SmartStructRequest{
		Image:          util.PtrString(""),
		Url:            util.PtrString(""),
		PdfFile:        util.PtrString(""),
		PdfFileNum:     util.PtrInt32(int32(0)),
		ReturnRelation: util.PtrBool(false),
	}
	result := &SmartStructResponse{}
	result, err := OCR_CLIENT.SmartStruct(smartStructRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SocialSecurityCard(t *testing.T) {
	socialSecurityCardRequest := &SocialSecurityCardRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &SocialSecurityCardResponse{}
	result, err := OCR_CLIENT.SocialSecurityCard(socialSecurityCardRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Table(t *testing.T) {
	tableRequest := &TableRequest{
		Image:        util.PtrString(""),
		Url:          util.PtrString(""),
		PdfFile:      util.PtrString(""),
		PdfFileNum:   util.PtrInt32(int32(0)),
		OfdFile:      util.PtrString(""),
		OfdFileNum:   util.PtrInt32(int32(0)),
		ReturnExcel:  util.PtrBool(false),
		CellContents: util.PtrBool(false),
	}
	result := &TableResponse{}
	result, err := OCR_CLIENT.Table(tableRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TaxiReceipt(t *testing.T) {
	taxiReceiptRequest := &TaxiReceiptRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &TaxiReceiptResponse{}
	result, err := OCR_CLIENT.TaxiReceipt(taxiReceiptRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ThreeFactorsVerification(t *testing.T) {
	threeFactorsVerificationRequest := &ThreeFactorsVerificationRequest{
		Name:    util.PtrString(""),
		Company: util.PtrString(""),
		Regnum:  util.PtrString(""),
	}
	result := &ThreeFactorsVerificationResponse{}
	result, err := OCR_CLIENT.ThreeFactorsVerification(threeFactorsVerificationRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TollInvoice(t *testing.T) {
	tollInvoiceRequest := &TollInvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &TollInvoiceResponse{}
	result, err := OCR_CLIENT.TollInvoice(tollInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TrainTicket(t *testing.T) {
	trainTicketRequest := &TrainTicketRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
	}
	result := &TrainTicketResponse{}
	result, err := OCR_CLIENT.TrainTicket(trainTicketRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_TwoFactorsVerification(t *testing.T) {
	twoFactorsVerificationRequest := &TwoFactorsVerificationRequest{
		Company: util.PtrString(""),
		Regnum:  util.PtrString(""),
	}
	result := &TwoFactorsVerificationResponse{}
	result, err := OCR_CLIENT.TwoFactorsVerification(twoFactorsVerificationRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UsedVehicleInvoice(t *testing.T) {
	usedVehicleInvoiceRequest := &UsedVehicleInvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &UsedVehicleInvoiceResponse{}
	result, err := OCR_CLIENT.UsedVehicleInvoice(usedVehicleInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VatInvoice(t *testing.T) {
	vatInvoiceRequest := &VatInvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
		OfdFile:    util.PtrString(""),
		OfdFileNum: util.PtrInt32(int32(0)),
		OcrType:    util.PtrString(""),
		SealTag:    util.PtrBool(false),
	}
	result := &VatInvoiceResponse{}
	result, err := OCR_CLIENT.VatInvoice(vatInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VehicleCertificate(t *testing.T) {
	vehicleCertificateRequest := &VehicleCertificateRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &VehicleCertificateResponse{}
	result, err := OCR_CLIENT.VehicleCertificate(vehicleCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VehicleInvoice(t *testing.T) {
	vehicleInvoiceRequest := &VehicleInvoiceRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &VehicleInvoiceResponse{}
	result, err := OCR_CLIENT.VehicleInvoice(vehicleInvoiceRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VehicleLicense(t *testing.T) {
	vehicleLicenseRequest := &VehicleLicenseRequest{
		Image:              util.PtrString(""),
		Url:                util.PtrString(""),
		DetectDirection:    util.PtrBool(false),
		VehicleLicenseSide: util.PtrString(""),
		Unified:            util.PtrBool(false),
		QualityWarn:        util.PtrBool(false),
		RiskWarn:           util.PtrBool(false),
	}
	result := &VehicleLicenseResponse{}
	result, err := OCR_CLIENT.VehicleLicense(vehicleLicenseRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VehicleRegCertificate(t *testing.T) {
	vehicleRegCertificateRequest := &VehicleRegCertificateRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &VehicleRegCertificateResponse{}
	result, err := OCR_CLIENT.VehicleRegCertificate(vehicleRegCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VehicleRegistrationCertificate(t *testing.T) {
	vehicleRegistrationCertificateRequest := &VehicleRegistrationCertificateRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &VehicleRegistrationCertificateResponse{}
	result, err := OCR_CLIENT.VehicleRegistrationCertificate(vehicleRegistrationCertificateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_VinCode(t *testing.T) {
	vinCodeRequest := &VinCodeRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &VinCodeResponse{}
	result, err := OCR_CLIENT.VinCode(vinCodeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Waybill(t *testing.T) {
	waybillRequest := &WaybillRequest{
		Image:                    util.PtrString(""),
		Url:                      util.PtrString(""),
		IsIdentifyVirtualWaybill: util.PtrBool(false),
	}
	result := &WaybillResponse{}
	result, err := OCR_CLIENT.Waybill(waybillRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_WebImage(t *testing.T) {
	webImageRequest := &WebImageRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		PdfFile:         util.PtrString(""),
		PdfFileNum:      util.PtrInt32(int32(0)),
		OfdFile:         util.PtrString(""),
		OfdFileNum:      util.PtrInt32(int32(0)),
		DetectDirection: util.PtrBool(false),
		DetectLanguage:  util.PtrBool(false),
	}
	result := &WebImageResponse{}
	result, err := OCR_CLIENT.WebImage(webImageRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_WebImageLoc(t *testing.T) {
	webImageLocRequest := &WebImageLocRequest{
		Image:                util.PtrString(""),
		Url:                  util.PtrString(""),
		PdfFile:              util.PtrString(""),
		PdfFileNum:           util.PtrInt32(int32(0)),
		OfdFile:              util.PtrString(""),
		OfdFileNum:           util.PtrInt32(int32(0)),
		DetectDirection:      util.PtrBool(false),
		Probability:          util.PtrBool(false),
		PolyLocation:         util.PtrBool(false),
		RecognizeGranularity: util.PtrString(""),
	}
	result := &WebImageLocResponse{}
	result, err := OCR_CLIENT.WebImageLoc(webImageLocRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_WeightNote(t *testing.T) {
	weightNoteRequest := &WeightNoteRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		PdfFile:     util.PtrString(""),
		PdfFileNum:  util.PtrInt32(int32(0)),
		Probability: util.PtrBool(false),
	}
	result := &WeightNoteResponse{}
	result, err := OCR_CLIENT.WeightNote(weightNoteRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
