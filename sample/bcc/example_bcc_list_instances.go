package bccsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func ListInstances() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bcc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	listInstancesRequest := &bcc.ListInstancesRequest{
		Marker:            util.PtrString(""),
		MaxKeys:           util.PtrInt32(int32(0)),
		InternalIp:        util.PtrString(""),
		DedicatedHostId:   util.PtrString(""),
		ZoneName:          util.PtrString(""),
		ShowRdmaTopo:      util.PtrString(""),
		InstanceIds:       util.PtrString(""),
		InstanceNames:     util.PtrString(""),
		FuzzyInstanceName: util.PtrString(""),
		VolumeIds:         util.PtrString(""),
		DeploySetIds:      util.PtrString(""),
		SecurityGroupIds:  util.PtrString(""),
		PaymentTiming:     util.PtrString(""),
		Status:            util.PtrString(""),
		Tags:              util.PtrString(""),
		VpcId:             util.PtrString(""),
		PrivateIps:        util.PtrString(""),
		EhcClusterId:      util.PtrString(""),
	}
	result, err := client.ListInstances(listInstancesRequest)
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
