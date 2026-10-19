package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func UpdateAutomaticStorageExpansionConfiguration() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAutomaticStorageExpansionConfigurationRequest := &rds.UpdateAutomaticStorageExpansionConfigurationRequest{
		InstanceId:         util.PtrString(""),
		Action:             util.PtrString(""),
		FreeSpaceThreshold: util.PtrInt32(int32(0)),
		DiskMaxLimit:       util.PtrInt32(int32(0)),
	}
	err = client.UpdateAutomaticStorageExpansionConfiguration(updateAutomaticStorageExpansionConfigurationRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
