package axsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ax"
)

func SetSandboxTimeout() {
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
	setSandboxTimeoutRequest := &ax.SetSandboxTimeoutRequest{
		SandboxID: util.PtrString(""),
		Timeout:   util.PtrInt32(int32(0)),
	}
	err = client.SetSandboxTimeout(setSandboxTimeoutRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
