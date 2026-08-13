package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func DeleteEnterpriseSecurityGroup() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteEnterpriseSecurityGroupRequest := &vpc.DeleteEnterpriseSecurityGroupRequest{
		EnterpriseSecurityGroupId: util.PtrString(""),
		ClientToken:               util.PtrString(""),
	}
	err = client.DeleteEnterpriseSecurityGroup(deleteEnterpriseSecurityGroupRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
