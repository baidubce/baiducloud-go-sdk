package scssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func CreateParameterTemplate() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createParameterTemplateRequest := &scs.CreateParameterTemplateRequest{
		Name:          util.PtrString(""),
		Engine:        util.PtrString(""),
		EngineVersion: util.PtrString(""),
		ClusterType:   util.PtrString(""),
		TemplateType:  util.PtrInt32(int32(0)),
		Comment:       util.PtrString(""),
		Parameters:    []*scs.Parameters{},
	}
	result, err := client.CreateParameterTemplate(createParameterTemplateRequest)
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
