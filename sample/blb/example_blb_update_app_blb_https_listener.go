package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateAppBlbHttpsListener() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAppBlbHttpsListenerRequest := &blb.UpdateAppBlbHttpsListenerRequest{
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
		ServerTimeout:         util.PtrInt32(int32(0)),
		CertIds:               []*string{},
		EncryptionType:        util.PtrString(""),
		EncryptionProtocols:   []*string{},
		AppliedCiphers:        util.PtrString(""),
		DualAuth:              util.PtrBool(false),
		ClientCertIds:         []*string{},
		AdditionalCertDomains: []*blb.AdditionalCertDomain{},
		Description:           util.PtrString(""),
	}
	err = client.UpdateAppBlbHttpsListener(updateAppBlbHttpsListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
