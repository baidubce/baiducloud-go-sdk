package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateBlbHttpListener() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	AdditionalAttributes := &blb.AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	updateBlbHttpListenerRequest := &blb.UpdateBlbHttpListenerRequest{
		BlbId:                      util.PtrString(""),
		ListenerPort:               util.PtrInt32(int32(0)),
		ClientToken:                util.PtrString(""),
		BackendPort:                util.PtrInt32(int32(0)),
		Scheduler:                  util.PtrString(""),
		KeepSession:                util.PtrBool(false),
		KeepSessionType:            util.PtrString(""),
		KeepSessionDuration:        util.PtrInt32(int32(0)),
		KeepSessionCookieName:      util.PtrString(""),
		XForwardFor:                util.PtrBool(false),
		XForwardedProto:            util.PtrBool(false),
		AdditionalAttributes:       AdditionalAttributes,
		HealthCheckType:            util.PtrString(""),
		HealthCheckHttpVersion:     util.PtrString(""),
		HealthCheckMethod:          util.PtrString(""),
		HealthCheckPort:            util.PtrInt32(int32(0)),
		HealthCheckURI:             util.PtrString(""),
		HealthCheckTimeoutInSecond: util.PtrInt32(int32(0)),
		HealthCheckInterval:        util.PtrInt32(int32(0)),
		UnhealthyThreshold:         util.PtrInt32(int32(0)),
		HealthyThreshold:           util.PtrInt32(int32(0)),
		HealthCheckNormalStatus:    util.PtrString(""),
		HealthCheckHost:            util.PtrString(""),
		ServerTimeout:              util.PtrInt32(int32(0)),
		RedirectPort:               util.PtrInt32(int32(0)),
	}
	err = client.UpdateBlbHttpListener(updateBlbHttpListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
