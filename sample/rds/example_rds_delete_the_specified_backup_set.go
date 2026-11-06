package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func DeleteTheSpecifiedBackupSet() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteTheSpecifiedBackupSetRequest := &rds.DeleteTheSpecifiedBackupSetRequest{
		InstanceId: util.PtrString(""),
		SnapshotId: util.PtrString(""),
	}
	err = client.DeleteTheSpecifiedBackupSet(deleteTheSpecifiedBackupSetRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
