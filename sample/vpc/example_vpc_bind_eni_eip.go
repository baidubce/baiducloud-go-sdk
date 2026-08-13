package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func BindEniEip() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	bindEniEipRequest := &vpc.BindEniEipRequest{
		EniId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		PrivateIpAddress: util.PtrString(""),
		PublicIpAddress:  util.PtrString(""),
	}
	err = client.BindEniEip(bindEniEipRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
