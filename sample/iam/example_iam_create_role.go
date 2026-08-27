package iamsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/iam"
)

func CreateRole() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := iam.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createRoleRequest := &iam.CreateRoleRequest{
		Name:                     util.PtrString(""),
		Description:              util.PtrString(""),
		GrantType:                util.PtrString(""),
		AssumeRolePolicyDocument: util.PtrString(""),
	}
	err = client.CreateRole(createRoleRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
