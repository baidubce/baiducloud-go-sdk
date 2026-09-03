package aigwsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aigw"
)

func GetConsumerList() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := aigw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getConsumerListRequest := &aigw.GetConsumerListRequest{
		InstanceId: util.PtrString(""),
		XRegion:    util.PtrString(""),
		PageNo:     util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
		TagKey:     util.PtrString(""),
		TagValue:   util.PtrString(""),
	}
	result, err := client.GetConsumerList(getConsumerListRequest)
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
