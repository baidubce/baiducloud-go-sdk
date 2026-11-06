package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func ModifyBackupStrategy() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	modifyBackupStrategyRequest := &rds.ModifyBackupStrategyRequest{
		InstanceId:                    util.PtrString(""),
		BackupDays:                    util.PtrString(""),
		BackupTime:                    util.PtrString(""),
		Persistent:                    util.PtrBool(false),
		ExpireInDays:                  util.PtrInt32(int32(0)),
		LogBackupRetainDays:           util.PtrInt32(int32(0)),
		DataBackupType:                util.PtrString(""),
		IncrementalDataBackupEnable:   util.PtrBool(false),
		IncrementalDataBackupInterval: util.PtrInt32(int32(0)),
		LatestDataBackupRetainDays:    util.PtrInt32(int32(0)),
		RetentionPeriod:               util.PtrInt32(int32(0)),
		TimeUnit:                      util.PtrString(""),
		DataBackupCopyEnable:          util.PtrBool(false),
		DataBackupCopyStoragesRegions: []*string{},
		DataBackupCopyRetainDays:      util.PtrInt32(int32(0)),
	}
	err = client.ModifyBackupStrategy(modifyBackupStrategyRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
