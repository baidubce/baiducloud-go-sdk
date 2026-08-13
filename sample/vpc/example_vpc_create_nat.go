package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateNat() {
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
	SessionConfig := &vpc.SessionConfig{
		TcpTimeout:  util.PtrInt32(int32(0)),
		UdpTimeout:  util.PtrInt32(int32(0)),
		IcmpTimeout: util.PtrInt32(int32(0)),
	}
	createNatRequest := &vpc.CreateNatRequest{
		ClientToken:     util.PtrString(""),
		Name:            util.PtrString(""),
		VpcId:           util.PtrString(""),
		CuNum:           util.PtrInt32(int32(0)),
		IpVersion:       util.PtrString(""),
		BindEips:        []*string{},
		Billing:         Billing,
		SessionConfig:   SessionConfig,
		Tags:            []*vpc.TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result, err := client.CreateNat(createNatRequest)
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
