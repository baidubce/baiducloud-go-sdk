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
