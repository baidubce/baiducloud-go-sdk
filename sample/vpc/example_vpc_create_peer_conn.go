package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreatePeerConn() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
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
	createPeerConnRequest := &vpc.CreatePeerConnRequest{
		ClientToken:     util.PtrString(""),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		Description:     util.PtrString(""),
		LocalIfName:     util.PtrString(""),
		LocalVpcId:      util.PtrString(""),
		PeerAccountId:   util.PtrString(""),
		PeerVpcId:       util.PtrString(""),
		PeerRegion:      util.PtrString(""),
		PeerIfName:      util.PtrString(""),
		Billing:         Billing,
		Tags:            []*vpc.TagModel{},
		ResourceGroupId: util.PtrString(""),
		DeleteProtect:   util.PtrBool(false),
	}
	result, err := client.CreatePeerConn(createPeerConnRequest)
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
