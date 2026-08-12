package axsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ax"
)

func CreateSandbox() {
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
	Metadata := make(map[string]string)
	EnvVars := make(map[string]string)
	AutoResume := make(map[string]interface{})
	Mcp := make(map[string]interface{})
	createSandboxRequest := &ax.CreateSandboxRequest{
		TemplateID:          util.PtrString(""),
		Timeout:             util.PtrInt32(int32(0)),
		Metadata:            nil,
		EnvVars:             nil,
		Secure:              util.PtrBool(false),
		AllowInternetAccess: util.PtrBool(false),
		AutoPause:           util.PtrBool(false),
		AutoResume:          nil,
		RuntimeType:         util.PtrString(""),
		Mcp:                 nil,
		VolumeMounts:        []*map[string]interface{}{},
	}
	result, err := client.CreateSandbox(createSandboxRequest)
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
