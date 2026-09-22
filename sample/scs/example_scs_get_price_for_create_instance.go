package scssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func GetPriceForCreateInstance() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getPriceForCreateInstanceRequest := &scs.GetPriceForCreateInstanceRequest{
		Engine:            util.PtrInt32(int32(0)),
		ClusterType:       util.PtrString(""),
		NodeType:          util.PtrString(""),
		CacheInstanceType: util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		ReplicationNum:    util.PtrInt32(int32(0)),
		InstanceNum:       util.PtrInt32(int32(0)),
		DiskType:          util.PtrString(""),
		DiskFlavor:        util.PtrInt32(int32(0)),
		ChargeType:        util.PtrString(""),
		Period:            util.PtrInt32(int32(0)),
		TimeUnit:          util.PtrString(""),
	}
	result, err := client.GetPriceForCreateInstance(getPriceForCreateInstanceRequest)
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
