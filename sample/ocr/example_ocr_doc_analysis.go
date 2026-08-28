package ocrsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ocr"
)

func DocAnalysis() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	// ak, sk := "Your Ak", "Your Sk"
	// client, err := ocr.NewClient(ak, sk, endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// apiKey, secretKey := "Your ApiKey", "Your SecretKey"
	// client, err := ocr.NewClientWithAccessToken(apiKey, secretKey, endpoint)

	// ==== API Key 鉴权 ====
	apiKey := "Your ApiKey"
	client, err := ocr.NewClientWithApiKey(apiKey, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	docAnalysisRequest := &ocr.DocAnalysisRequest{
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
	result, err := client.DocAnalysis(docAnalysisRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
}
