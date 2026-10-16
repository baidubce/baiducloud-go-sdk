package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func CreateAppBlbSslListener() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createAppBlbSslListenerRequest := &blb.CreateAppBlbSslListenerRequest{
		BlbId:               util.PtrString(""),
		ClientToken:         util.PtrString(""),
		ListenerPort:        util.PtrInt32(int32(0)),
		Scheduler:           util.PtrString(""),
		CertIds:             []*string{},
		EncryptionType:      util.PtrString(""),
		EncryptionProtocols: []*string{},
		AppliedCiphers:      util.PtrString(""),
		DualAuth:            util.PtrBool(false),
		ClientCertIds:       []*string{},
		Description:         util.PtrString(""),
	}
	err = client.CreateAppBlbSslListener(createAppBlbSslListenerRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
