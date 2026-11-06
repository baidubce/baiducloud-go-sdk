package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func CreateADatabase() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createADatabaseRequest := &rds.CreateADatabaseRequest{
		InstanceId:        util.PtrString(""),
		CharacterSetName:  util.PtrString(""),
		DbName:            util.PtrString(""),
		Remark:            util.PtrString(""),
		AccountPrivileges: []*rds.AccountPrivilege{},
		Ctype:             util.PtrString(""),
		Collate:           util.PtrString(""),
		Owner:             util.PtrString(""),
	}
	err = client.CreateADatabase(createADatabaseRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
