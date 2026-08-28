package ocrsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ocr"
)

func FourFactorsVerification() {
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
	fourFactorsVerificationRequest := &ocr.FourFactorsVerificationRequest{
		Name:    util.PtrString(""),
		Idcard:  util.PtrString(""),
		Company: util.PtrString(""),
		Regnum:  util.PtrString(""),
	}
	result, err := client.FourFactorsVerification(fourFactorsVerificationRequest)
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
