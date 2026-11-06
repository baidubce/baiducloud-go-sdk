package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func RestoreLibraryTablesBasedOnBackupSets() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	restoreLibraryTablesBasedOnBackupSetsRequest := &rds.RestoreLibraryTablesBasedOnBackupSetsRequest{
		InstanceId:       util.PtrString(""),
		SourceInstanceId: util.PtrString(""),
		TargetInstanceId: util.PtrString(""),
		SnapshotId:       util.PtrString(""),
		Data:             []*rds.RecoveryToSourceInstanceModel{},
		DbName:           util.PtrString(""),
		NewDbname:        util.PtrString(""),
		RestoreMode:      util.PtrString(""),
		Tables:           []*rds.RecoveryToSourceInstanceTableModel{},
		TableName:        util.PtrString(""),
		NewTablename:     util.PtrString(""),
	}
	err = client.RestoreLibraryTablesBasedOnBackupSets(restoreLibraryTablesBasedOnBackupSetsRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
