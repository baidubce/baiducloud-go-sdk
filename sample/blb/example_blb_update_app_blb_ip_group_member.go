package blbsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func UpdateAppBlbIpGroupMember() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := blb.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAppBlbIpGroupMemberRequest := &blb.UpdateAppBlbIpGroupMemberRequest{
		BlbId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		IpGroupId:   util.PtrString(""),
		MemberList:  []*blb.AppIpGroupMemberForUpdate{},
	}
	err = client.UpdateAppBlbIpGroupMember(updateAppBlbIpGroupMemberRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
