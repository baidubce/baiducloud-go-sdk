package axsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ax"
)

func DeleteSandbox() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	// ak, sk := "Your Ak", "Your Sk"
	// client, err := ax.NewClient(ak, sk, endpoint)

	// ==== API Key 鉴权 ====
	apiKey := "Your ApiKey"
	client, err := ax.NewClientWithApiKey(apiKey, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteSandboxRequest := &ax.DeleteSandboxRequest{
		SandboxID: util.PtrString(""),
	}
	err = client.DeleteSandbox(deleteSandboxRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
