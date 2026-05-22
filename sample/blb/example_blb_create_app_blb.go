package blbsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/blb"
)

func CreateAppBlb() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := blb.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &blb.BillingForCreate{
		PaymentTiming: util.PtrString(""),
		BillingMethod: util.PtrString(""),
		Reservation: &blb.ReservationForCreate{
			ReservationLength: util.PtrInt32(int32(0)),
		},
	}
	createAppBlbRequest := &blb.CreateAppBlbRequest{
		ClientToken:                  util.PtrString(""),
		Name:                         util.PtrString(""),
		BlbType:                      util.PtrString(""),
		Desc:                         util.PtrString(""),
		SubnetId:                     util.PtrString(""),
		VpcId:                        util.PtrString(""),
		Address:                      util.PtrString(""),
		Eip:                          util.PtrString(""),
		Tags:                         []*blb.TagModel{},
		Billing:                      Billing,
		PerformanceLevel:             util.PtrString(""),
		AutoRenewLength:              util.PtrInt32(int32(0)),
		AutoRenewTimeUnit:            util.PtrString(""),
		ResourceGroupId:              util.PtrString(""),
		AllowDelete:                  util.PtrBool(false),
		AllowModify:                  util.PtrBool(false),
		ModificationProtectionReason: util.PtrString(""),
		AllocateIpv6:                 util.PtrBool(false),
	}
	result, err := client.CreateAppBlb(createAppBlbRequest)
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
