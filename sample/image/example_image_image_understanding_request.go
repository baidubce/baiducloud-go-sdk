package imagesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/image"
)

func ImageUnderstandingRequest() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	// ak, sk := "Your Ak", "Your Sk"
	// client, err := image.NewClient(ak, sk, endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// apiKey, secretKey := "Your ApiKey", "Your SecretKey"
	// client, err := image.NewClientWithAccessToken(apiKey, secretKey, endpoint)

	// ==== API Key 鉴权 ====
	apiKey := "Your ApiKey"
	client, err := image.NewClientWithApiKey(apiKey, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	imageUnderstandingRequestRequest := &image.ImageUnderstandingRequestRequest{
		Question: util.PtrString(""),
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
	}
	result, err := client.ImageUnderstandingRequest(imageUnderstandingRequestRequest)
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
