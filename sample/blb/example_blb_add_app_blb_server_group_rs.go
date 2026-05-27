package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func AddAppBlbServerGroupRs() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := blb.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	addAppBlbServerGroupRsRequest := &blb.AddAppBlbServerGroupRsRequest{
		BlbId:             util.PtrString(""),
		ClientToken:       util.PtrString(""),
		SgId:              util.PtrString(""),
		BackendServerList: []*blb.AppBackendServer{},
	}
	err = client.AddAppBlbServerGroupRs(addAppBlbServerGroupRsRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
