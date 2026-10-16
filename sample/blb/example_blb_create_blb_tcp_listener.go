package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func CreateBlbTcpListener() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createBlbTcpListenerRequest := &blb.CreateBlbTcpListenerRequest{
		BlbId:                      util.PtrString(""),
		ClientToken:                util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		TcpSessionTimeout:          util.PtrInt32(int32(0)),
		HealthCheckType:            util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
	}
	err = client.CreateBlbTcpListener(createBlbTcpListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
