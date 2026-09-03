package aigwsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aigw"
)

func CreateService() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := aigw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createServiceRequest := &aigw.CreateServiceRequest{
		InstanceId:       util.PtrString(""),
		XRegion:          util.PtrString(""),
		ServiceSource:    util.PtrString(""),
		Namespace:        util.PtrString(""),
		ServiceName:      util.PtrString(""),
		ClusterId:        util.PtrString(""),
		ClusterIds:       []*string{},
		ServiceList:      []*aigw.ServiceItem{},
		RegistryId:       util.PtrString(""),
		ServiceAddresses: []*string{},
		ServiceProtocol:  util.PtrString(""),
		Provider:         util.PtrString(""),
		Endpoint:         util.PtrString(""),
		ApiKeys:          []*string{},
		CredentialSource: util.PtrString(""),
		CredentialNames:  []*string{},
		FailoverEnabled:  util.PtrBool(false),
		FailoverModel:    util.PtrString(""),
	}
	result, err := client.CreateService(createServiceRequest)
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
