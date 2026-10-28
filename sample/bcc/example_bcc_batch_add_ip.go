package bccsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func BatchAddIp() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := bcc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	batchAddIpRequest := &bcc.BatchAddIpRequest{
		InstanceId:                     util.PtrString(""),
		SecondaryPrivateIpAddressCount: util.PtrInt32(int32(0)),
		PrivateIps:                     []*string{},
		AllocateMultiIpv6Addr:          util.PtrBool(false),
	}
	err = client.BatchAddIp(batchAddIpRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
