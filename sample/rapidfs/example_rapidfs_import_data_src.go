package rapidfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rapidfs"
)

func ImportDataSrc() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := rapidfs.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	importDataSrcRequest := &rapidfs.ImportDataSrcRequest{
		Action:         util.PtrString(""),
		ClientToken:    util.PtrString(""),
		InstanceId:     util.PtrString(""),
		DataSrcName:    util.PtrString(""),
		Bucket:         util.PtrString(""),
		OtherAccount:   util.PtrBool(false),
		BucketAK:       util.PtrString(""),
		BucketSK:       util.PtrString(""),
		BucketStsToken: util.PtrString(""),
		DirPrefix:      util.PtrString(""),
		KeepSymlink:    util.PtrBool(false),
		AuthGroupId:    util.PtrString(""),
		Description:    util.PtrString(""),
		QuotaMiB:       util.PtrInt32(int32(0)),
	}
	result, err := client.ImportDataSrc(importDataSrcRequest)
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
