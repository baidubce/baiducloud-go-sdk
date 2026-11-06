package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func CreateAccount() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createAccountRequest := &rds.CreateAccountRequest{
		InstanceId:         util.PtrString(""),
		AccountName:        util.PtrString(""),
		Password:           util.PtrString(""),
		AccountType:        util.PtrString(""),
		DatabasePrivileges: []*rds.DatabasePrivilege{},
		Desc:               util.PtrString(""),
		RdsType:            util.PtrString(""),
	}
	err = client.CreateAccount(createAccountRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
