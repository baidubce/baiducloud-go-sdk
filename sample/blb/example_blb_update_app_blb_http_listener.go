package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateAppBlbHttpListener() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := blb.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	AdditionalAttributes := &blb.AdditionalAttributesModel{
		GzipJson: util.PtrString(""),
	}
	updateAppBlbHttpListenerRequest := &blb.UpdateAppBlbHttpListenerRequest{
		BlbId:                 util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		ClientToken:           util.PtrString(""),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		AdditionalAttributes:  AdditionalAttributes,
		ServerTimeout:         util.PtrInt32(int32(0)),
		RedirectPort:          util.PtrInt32(int32(0)),
		Description:           util.PtrString(""),
	}
	err = client.UpdateAppBlbHttpListener(updateAppBlbHttpListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
