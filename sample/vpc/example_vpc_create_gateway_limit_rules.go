package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateGatewayLimitRules() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := vpc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createGatewayLimitRulesRequest := &vpc.CreateGatewayLimitRulesRequest{
		ClientToken:    util.PtrString(""),
		IpVersion:      util.PtrString(""),
		Name:           util.PtrString(""),
		Description:    util.PtrString(""),
		ServiceType:    util.PtrString(""),
		SubServiceType: util.PtrString(""),
		PeerRegion:     util.PtrString(""),
		ResourceId:     util.PtrString(""),
		Direction:      util.PtrString(""),
		Cidr:           util.PtrString(""),
		Bandwidth:      util.PtrInt32(int32(0)),
	}
	result := &vpc.CreateGatewayLimitRulesResponse{}
	result, err = client.CreateGatewayLimitRules(createGatewayLimitRulesRequest)
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
