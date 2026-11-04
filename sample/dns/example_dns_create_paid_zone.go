package dnssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/dns"
)

func CreatePaidZone() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := dns.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &dns.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &dns.Reservation{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	createPaidZoneRequest := &dns.CreatePaidZoneRequest{
		ClientToken:    util.PtrString(""),
		Names:          []*string{},
		ProductVersion: util.PtrString(""),
		Billing:        Billing,
	}
	err = client.CreatePaidZone(createPaidZoneRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
