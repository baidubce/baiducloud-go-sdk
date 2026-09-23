package nlp

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
	NLP_CLIENT *Client
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
	// NLP_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// NLP_CLIENT, _ = NewClientWithAccessToken(confObj.ApiKey, confObj.SecretKey, confObj.Endpoint)

	// ==== API Key 鉴权 ====
	NLP_CLIENT, _ = NewClientWithApiKey(confObj.ApiKey, confObj.Endpoint)

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

func TestClient_Address(t *testing.T) {
	addressRequest := &AddressRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
	}
	result := &AddressResponse{}
	result, err := NLP_CLIENT.Address(addressRequest)
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
func TestClient_CommentTag(t *testing.T) {
	commentTagRequest := &CommentTagRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
		NlpType: util.PtrInt32(int32(0)),
	}
	result := &CommentTagResponse{}
	result, err := NLP_CLIENT.CommentTag(commentTagRequest)
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
func TestClient_Ecnet(t *testing.T) {
	ecnetRequest := &EcnetRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
	}
	result := &EcnetResponse{}
	result, err := NLP_CLIENT.Ecnet(ecnetRequest)
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
func TestClient_Emotion(t *testing.T) {
	emotionRequest := &EmotionRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
		Scene:   util.PtrString(""),
	}
	result := &EmotionResponse{}
	result, err := NLP_CLIENT.Emotion(emotionRequest)
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
func TestClient_EntityAnalysis(t *testing.T) {
	entityAnalysisRequest := &EntityAnalysisRequest{
		Text:    util.PtrString(""),
		Mention: util.PtrString(""),
	}
	result := &EntityAnalysisResponse{}
	result, err := NLP_CLIENT.EntityAnalysis(entityAnalysisRequest)
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
func TestClient_Keyword(t *testing.T) {
	keywordRequest := &KeywordRequest{
		Charset: util.PtrString(""),
		Title:   util.PtrString(""),
		Content: util.PtrString(""),
	}
	result := &KeywordResponse{}
	result, err := NLP_CLIENT.Keyword(keywordRequest)
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
func TestClient_Lexer(t *testing.T) {
	lexerRequest := &LexerRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
	}
	result := &LexerResponse{}
	result, err := NLP_CLIENT.Lexer(lexerRequest)
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
func TestClient_NewsSummary(t *testing.T) {
	newsSummaryRequest := &NewsSummaryRequest{
		Charset:       util.PtrString(""),
		Title:         util.PtrString(""),
		Content:       util.PtrString(""),
		MaxSummaryLen: util.PtrInt32(int32(0)),
	}
	result := &NewsSummaryResponse{}
	result, err := NLP_CLIENT.NewsSummary(newsSummaryRequest)
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
func TestClient_SentimentClassify(t *testing.T) {
	sentimentClassifyRequest := &SentimentClassifyRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
	}
	result := &SentimentClassifyResponse{}
	result, err := NLP_CLIENT.SentimentClassify(sentimentClassifyRequest)
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
func TestClient_Simnet(t *testing.T) {
	simnetRequest := &SimnetRequest{
		Charset: util.PtrString(""),
		Text1:   util.PtrString(""),
		Text2:   util.PtrString(""),
		Model:   util.PtrString(""),
	}
	result := &SimnetResponse{}
	result, err := NLP_CLIENT.Simnet(simnetRequest)
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
func TestClient_TextCorrection(t *testing.T) {
	textCorrectionRequest := &TextCorrectionRequest{
		Charset: util.PtrString(""),
		Text:    util.PtrString(""),
	}
	result := &TextCorrectionResponse{}
	result, err := NLP_CLIENT.TextCorrection(textCorrectionRequest)
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
func TestClient_Topic(t *testing.T) {
	topicRequest := &TopicRequest{
		Charset: util.PtrString(""),
		Title:   util.PtrString(""),
		Content: util.PtrString(""),
	}
	result := &TopicResponse{}
	result, err := NLP_CLIENT.Topic(topicRequest)
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
func TestClient_TxtKeywordsExtraction(t *testing.T) {
	txtKeywordsExtractionRequest := &TxtKeywordsExtractionRequest{
		Charset: util.PtrString(""),
		Text:    []*string{},
		Num:     util.PtrInt32(int32(0)),
	}
	result := &TxtKeywordsExtractionResponse{}
	result, err := NLP_CLIENT.TxtKeywordsExtraction(txtKeywordsExtractionRequest)
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
func TestClient_TxtMonet(t *testing.T) {
	txtMonetRequest := &TxtMonetRequest{
		Charset:     util.PtrString(""),
		ContentList: []*ContentItem{},
	}
	result := &TxtMonetResponse{}
	result, err := NLP_CLIENT.TxtMonet(txtMonetRequest)
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
