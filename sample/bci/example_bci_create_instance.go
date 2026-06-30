package bcisample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bci"
)

func CreateInstance() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bci.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	volume := &bci.Volume{
		Nfs:        []*bci.NfsVolume{},
		EmptyDir:   []*bci.EmptyDirVolume{},
		ConfigFile: []*bci.ConfigFileVolume{},
	}
	createInstanceRequest := &bci.CreateInstanceRequest{
		ClientToken:                   util.PtrString(""),
		Name:                          util.PtrString(""),
		ZoneName:                      util.PtrString(""),
		SecurityGroupIds:              []*string{},
		SubnetIds:                     []*string{},
		RestartPolicy:                 util.PtrString(""),
		EipIp:                         util.PtrString(""),
		AutoCreateEip:                 util.PtrBool(false),
		EipName:                       util.PtrString(""),
		EipRouteType:                  util.PtrString(""),
		EipBandwidthInMbps:            util.PtrInt32(int32(0)),
		EipBillingMethod:              util.PtrString(""),
		GpuType:                       util.PtrString(""),
		TerminationGracePeriodSeconds: util.PtrInt64(int64(0)),
		HostName:                      util.PtrString(""),
		Tags:                          []*bci.Tag{},
		ImageRegistryCredentials:      []*bci.ImageRegistryCredential{},
		Containers:                    []*bci.Container{},
		InitContainers:                []*bci.Container{},
		Volume:                        volume,
	}
	result, err := client.CreateInstance(createInstanceRequest)
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
