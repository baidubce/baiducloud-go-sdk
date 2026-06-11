package aihcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aihc"
)

func CreateADatasetV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := aihc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	InitVersionEntry := &aihc.DatasetVersionEntry{
		Id:          util.PtrString(""),
		Version:     util.PtrString(""),
		Description: util.PtrString(""),
		StoragePath: util.PtrString(""),
		MountPath:   util.PtrString(""),
		CreateUser:  util.PtrString(""),
	}
	createADatasetV2Request := &aihc.CreateADatasetV2Request{
		Name:             util.PtrString(""),
		StorageType:      util.PtrString(""),
		StorageInstance:  util.PtrString(""),
		ImportFormat:     util.PtrString(""),
		Description:      util.PtrString(""),
		Owner:            util.PtrString(""),
		VisibilityScope:  util.PtrString(""),
		VisibilityUser:   []*aihc.PermissionEntry{},
		VisibilityGroup:  []*aihc.PermissionEntry{},
		InitVersionEntry: InitVersionEntry,
	}
	result, err := client.CreateADatasetV2(createADatasetV2Request)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
}
