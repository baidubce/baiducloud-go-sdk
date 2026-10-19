package rdssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func BatchConfigurationOfInstances() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	batchConfigurationOfInstancesRequest := &rds.BatchConfigurationOfInstancesRequest{
		InstanceId:     util.PtrString(""),
		CpuCount:       util.PtrInt32(int32(0)),
		MemoryCapacity: util.PtrInt32(int32(0)),
		VolumeCapacity: util.PtrInt32(int32(0)),
		ReadReplicas:   []*rds.ReplicaRequest{},
		IsBatchResize:  util.PtrBool(false),
		IsDirectPay:    util.PtrBool(false),
		IsEnhanced:     util.PtrBool(false),
		EffectiveTime:  util.PtrString(""),
		MasterAzone:    util.PtrString(""),
		BackupAzone:    util.PtrString(""),
		DiskIoType:     util.PtrString(""),
		SubnetId:       util.PtrString(""),
		EdgeSubnetId:   util.PtrString(""),
		Subnets:        []*rds.SubnetMap{},
		CouponId:       []*int64{},
	}
	result, err := client.BatchConfigurationOfInstances(batchConfigurationOfInstancesRequest)
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
