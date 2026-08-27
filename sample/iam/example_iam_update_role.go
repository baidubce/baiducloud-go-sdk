package iamsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/iam"
)

func UpdateRole() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := iam.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateRoleRequest := &iam.UpdateRoleRequest{
		RoleName:                 util.PtrString(""),
		Name:                     util.PtrString(""),
		Description:              util.PtrString(""),
		AssumeRolePolicyDocument: util.PtrString(""),
	}
	err = client.UpdateRole(updateRoleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
