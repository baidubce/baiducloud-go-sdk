package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateProbe() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createProbeRequest := &vpc.CreateProbeRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		VpcId:       util.PtrString(""),
		SubnetId:    util.PtrString(""),
		Protocol:    util.PtrString(""),
		Frequency:   util.PtrInt32(int32(0)),
		SourceIps:   []*string{},
		SourceIpNum: util.PtrInt32(int32(0)),
		DestIp:      util.PtrString(""),
		DestPort:    util.PtrInt32(int32(0)),
		Payload:     util.PtrString(""),
	}
	result, err := client.CreateProbe(createProbeRequest)
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
