package scssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func SetBackupPolicy() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	setBackupPolicyRequest := &scs.SetBackupPolicyRequest{
		InstanceId: util.PtrString(""),
		BackupTime: util.PtrString(""),
		BackupDays: util.PtrString(""),
		ExpireDay:  util.PtrInt32(int32(0)),
		IsEncrypt:  util.PtrString(""),
	}
	err = client.SetBackupPolicy(setBackupPolicyRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
