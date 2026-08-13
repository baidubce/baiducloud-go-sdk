package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateVpn() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &vpc.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &vpc.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createVpnRequest := &vpc.CreateVpnRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		SubnetId:        util.PtrString(""),
		VpnName:         util.PtrString(""),
		VpcType:         util.PtrString(""),
		Description:     util.PtrString(""),
		Eip:             util.PtrString(""),
		Tags:            []*vpc.TagModel{},
		ResourceGroupId: util.PtrString(""),
		Billing:         Billing,
		MaxConnection:   util.PtrInt32(int32(0)),
		DeleteProtect:   util.PtrBool(false),
	}
	result, err := client.CreateVpn(createVpnRequest)
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
