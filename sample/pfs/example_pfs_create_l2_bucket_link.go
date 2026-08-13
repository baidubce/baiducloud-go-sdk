package pfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func CreateL2BucketLink() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createL2BucketLinkRequest := &pfs.CreateL2BucketLinkRequest{
		Action:               util.PtrString(""),
		InstanceId:           util.PtrString(""),
		ConflictPolicy:       util.PtrString(""),
		BucketName:           util.PtrString(""),
		BucketPrefix:         util.PtrString(""),
		ThroughputLimitBytes: util.PtrString(""),
		ReportObjectName:     util.PtrString(""),
		BucketLinkName:       util.PtrString(""),
		TransferType:         util.PtrInt32(int32(0)),
		PfsPath:              util.PtrString(""),
		Cron:                 util.PtrString(""),
		BucketBelongUserId:   util.PtrString(""),
		LccId:                util.PtrString(""),
		Scope:                util.PtrInt32(int32(0)),
	}
	result, err := client.CreateL2BucketLink(createL2BucketLinkRequest)
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
