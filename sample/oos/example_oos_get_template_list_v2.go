package oossample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/oos"
)

func GetTemplateListV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := oos.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getTemplateListV2Request := &oos.GetTemplateListV2Request{
		Locale:                util.PtrString(""),
		Namespace:             util.PtrString(""),
		Name:                  util.PtrString(""),
		Id:                    util.PtrString(""),
		OosType:               util.PtrString(""),
		Sort:                  util.PtrString(""),
		Ascending:             util.PtrBool(false),
		PageNo:                util.PtrInt32(int32(0)),
		PageSize:              util.PtrInt32(int32(0)),
		SupportedInstanceType: util.PtrString(""),
	}
	result, err := client.GetTemplateListV2(getTemplateListV2Request)
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
