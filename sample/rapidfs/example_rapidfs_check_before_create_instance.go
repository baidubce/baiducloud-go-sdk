package rapidfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rapidfs"
)

func CheckBeforeCreateInstance() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := rapidfs.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	checkBeforeCreateInstanceRequest := &rapidfs.CheckBeforeCreateInstanceRequest{
		Action:             util.PtrString(""),
		Zone:               util.PtrString(""),
		VpcId:              util.PtrString(""),
		SubnetId:           util.PtrString(""),
		ManagedMode:        util.PtrString(""),
		MetaSpec:           util.PtrString(""),
		DataSpec:           util.PtrString(""),
		RapidfsType:        util.PtrString(""),
		CapacityTiB:        util.PtrInt32(int32(0)),
		CceClusterId:       util.PtrString(""),
		AihcResourcePoolId: util.PtrString(""),
		K8sControllerId:    util.PtrString(""),
		K8sControllerToken: util.PtrString(""),
	}
	result, err := client.CheckBeforeCreateInstance(checkBeforeCreateInstanceRequest)
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
