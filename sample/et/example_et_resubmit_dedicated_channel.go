package etsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/et"
)

func ResubmitDedicatedChannel() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := et.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	resubmitDedicatedChannelRequest := &et.ResubmitDedicatedChannelRequest{
		EtId:                util.PtrString(""),
		EtChannelId:         util.PtrString(""),
		ClientToken:         util.PtrString(""),
		AuthorizedUsers:     []*string{},
		Description:         util.PtrString(""),
		BaiduAddress:        util.PtrString(""),
		Name:                util.PtrString(""),
		Networks:            []*string{},
		CustomerAddress:     util.PtrString(""),
		RouteType:           util.PtrString(""),
		VlanId:              util.PtrInt32(int32(0)),
		EnableIpv6:          util.PtrInt32(int32(0)),
		BaiduIpv6Address:    util.PtrString(""),
		CustomerIpv6Address: util.PtrString(""),
		Ipv6Networks:        []*string{},
	}
	err = client.ResubmitDedicatedChannel(resubmitDedicatedChannelRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
