package etsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/et"
)

func ApplyPhysicalDedicatedLine() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := et.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Billing := &et.Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &et.Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	AutoRenew := &et.Reservation{
		ReservationLength:   util.PtrInt32(int32(0)),
		ReservationTimeUnit: util.PtrString(""),
	}
	applyPhysicalDedicatedLineRequest := &et.ApplyPhysicalDedicatedLineRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Isp:         util.PtrString(""),
		IntfType:    util.PtrString(""),
		ApType:      util.PtrString(""),
		ApAddr:      util.PtrString(""),
		LinkDelay:   util.PtrInt32(int32(0)),
		UserName:    util.PtrString(""),
		UserPhone:   util.PtrString(""),
		UserEmail:   util.PtrString(""),
		UserIdc:     util.PtrString(""),
		Billing:     Billing,
		AutoRenew:   AutoRenew,
		Tags:        []*et.TagModel{},
	}
	result, err := client.ApplyPhysicalDedicatedLine(applyPhysicalDedicatedLineRequest)
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
