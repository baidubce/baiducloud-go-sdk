package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateAppBlbUdpListener() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAppBlbUdpListenerRequest := &blb.UpdateAppBlbUdpListenerRequest{
		BlbId:             util.PtrString(""),
		ListenerPort:      util.PtrInt32(int32(0)),
		ClientToken:       util.PtrString(""),
		Scheduler:         util.PtrString(""),
		UdpSessionTimeout: util.PtrInt32(int32(0)),
		Description:       util.PtrString(""),
	}
	err = client.UpdateAppBlbUdpListener(updateAppBlbUdpListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
