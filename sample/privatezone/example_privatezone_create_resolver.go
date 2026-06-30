package privatezonesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/privatezone"
)

func CreateResolver() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := privatezone.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	// ClientToken 用作幂等标识：同一次逻辑创建必须使用同一个 token。
	// SDK 默认在网络抖动/5xx 等情况会自动重试，相同 ClientToken 可让服务端识别为同一请求，
	// 避免重复创建解析器。
	createResolverRequest := &privatezone.CreateResolverRequest{
		ClientToken:     util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		VpcRegion:       util.PtrString(""),
		IpModels: []*privatezone.IpModel{
        		{
        			SubnetId:  util.PtrString(""),
        			IpAddress: util.PtrString(""),
        		},
        		{
        			SubnetId:  util.PtrString(""),
        			IpAddress: util.PtrString(""),
        		},
        	},
		PrivatezoneType: util.PtrString(""),
	}
	result, err := client.CreateResolver(createResolverRequest)
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
