package ccrsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/ccr"
)

func AddVpcLink() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := ccr.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	addVpcLinkRequest := &ccr.AddVpcLinkRequest{
		InstanceId:     util.PtrString(""),
		VpcID:          util.PtrString(""),
		SubnetID:       util.PtrString(""),
		IpType:         util.PtrString(""),
		IpAddress:      util.PtrString(""),
		AutoDnsResolve: util.PtrBool(false),
	}
	err = client.AddVpcLink(addVpcLinkRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
