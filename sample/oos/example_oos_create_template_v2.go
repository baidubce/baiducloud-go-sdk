package oossample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/oos"
)

func CreateTemplateV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := oos.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createTemplateV2Request := &oos.CreateTemplateV2Request{
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Tags:        []*oos.KeyValuePair{},
		Linear:      util.PtrBool(false),
		Parallelism: util.PtrInt32(int32(0)),
		Operators:   []*oos.Operator{},
		Links:       []*oos.LinkModel{},
		Properties:  []*oos.Property{},
	}
	result, err := client.CreateTemplateV2(createTemplateV2Request)
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
