package scssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func ChangeConfiguration() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &scs.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &scs.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	changeConfigurationRequest := &scs.ChangeConfigurationRequest{
		InstanceId:    util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Billing:       Billing,
		EngineVersion: util.PtrString(""),
		NodeType:      util.PtrString(""),
		ShardNum:      util.PtrInt32(int32(0)),
		DiskFlavor:    util.PtrInt32(int32(0)),
	}
	result, err := client.ChangeConfiguration(changeConfigurationRequest)
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
