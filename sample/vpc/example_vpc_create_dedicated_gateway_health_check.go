package vpcsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateDedicatedGatewayHealthCheck() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createDedicatedGatewayHealthCheckRequest := &vpc.CreateDedicatedGatewayHealthCheckRequest{
		EtGatewayId:           util.PtrString(""),
		ClientToken:           util.PtrString(""),
		DcphyId:               util.PtrString(""),
		ChannelId:             util.PtrString(""),
		SubnetId:              util.PtrString(""),
		HealthCheckSourceIp:   util.PtrString(""),
		HealthCheckType:       util.PtrString(""),
		HealthCheckInterval:   util.PtrInt32(int32(0)),
		HealthThreshold:       util.PtrInt32(int32(0)),
		UnhealthThreshold:     util.PtrInt32(int32(0)),
		AutoGenerateRouteRule: util.PtrBool(false),
	}
	err = client.CreateDedicatedGatewayHealthCheck(createDedicatedGatewayHealthCheckRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
