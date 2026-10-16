package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateBlb() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateBlbRequest := &blb.UpdateBlbRequest{
		BlbId:        util.PtrString(""),
		ClientToken:  util.PtrString(""),
		Name:         util.PtrString(""),
		Desc:         util.PtrString(""),
		AllowDelete:  util.PtrBool(false),
		AllocateIpv6: util.PtrBool(false),
	}
	err = client.UpdateBlb(updateBlbRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
