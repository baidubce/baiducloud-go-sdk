package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func CreateAppBlbHttpsListener() {
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
	createAppBlbHttpsListenerRequest := &blb.CreateAppBlbHttpsListenerRequest{
		BlbId:                 util.PtrString(""),
		ClientToken:           util.PtrString(""),
		ListenerPort:          util.PtrInt32(int32(0)),
		Scheduler:             util.PtrString(""),
		KeepSession:           util.PtrBool(false),
		KeepSessionType:       util.PtrString(""),
		KeepSessionTimeout:    util.PtrInt32(int32(0)),
		KeepSessionCookieName: util.PtrString(""),
		XForwardedFor:         util.PtrBool(false),
		XForwardedProto:       util.PtrBool(false),
		AdditionalAttributes:  AdditionalAttributes,
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
	err = client.CreateAppBlbHttpsListener(createAppBlbHttpsListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
