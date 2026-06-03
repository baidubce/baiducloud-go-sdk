package pfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func UpdPerL2BktLnkInfo() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := pfs.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updPerL2BktLnkInfoRequest := &pfs.UpdPerL2BktLnkInfoRequest{
		Action:                  util.PtrString(""),
		InstanceId:              util.PtrString(""),
		BucketLinkId:            util.PtrString(""),
		NewCron:                 util.PtrString(""),
		NewBucketLinkName:       util.PtrString(""),
		NewConflictPolicy:       util.PtrInt32(int32(0)),
		NewThroughputLimitBytes: util.PtrInt32(int32(0)),
		NewScope:                util.PtrInt32(int32(0)),
	}
	result, err := client.UpdPerL2BktLnkInfo(updPerL2BktLnkInfoRequest)
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
