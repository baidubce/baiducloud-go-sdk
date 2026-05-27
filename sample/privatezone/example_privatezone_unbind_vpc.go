package privatezonesample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/privatezone"
)

func UnbindVpc() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := privatezone.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	unbindVpcRequest := &privatezone.UnbindVpcRequest{
		ZoneId:      util.PtrString(""),
		Action:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		Region:      util.PtrString("bj"),
		VpcIds:      []*string{util.PtrString("vpc-8zn7k6fny75x")},
	}
	err = client.UnbindVpc(unbindVpcRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		if serviceErr, ok := err.(*bce.BceServiceError); ok {
			fmt.Printf("request failed, RequestId: %s, Code: %s, Message: %s\n",
				serviceErr.RequestId, serviceErr.Code, serviceErr.Message)
		} else {
			fmt.Println("request failed:", err)
		}
		return
	}
	fmt.Println("UnbindVpc success")
}
