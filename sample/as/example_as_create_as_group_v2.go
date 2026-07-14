package assample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/as"
)

func CreateAsGroupV2() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := as.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Config := &as.GroupConfig{
		MinNodeNum:    util.PtrInt32(int32(0)),
		MaxNodeNum:    util.PtrInt32(int32(0)),
		CooldownInSec: util.PtrInt32(int32(0)),
		ExpectNum:     util.PtrInt32(int32(0)),
		InitNum:       util.PtrInt32(int32(0)),
	}
	Eip := &as.EipInfo{
		IfBindEip:       util.PtrBool(false),
		BandwidthInMbps: util.PtrInt32(int32(0)),
		EipProductType:  util.PtrString(""),
		PurchaseType:    util.PtrString(""),
	}
	EipConfig := &as.EipConfig{
		EipGroupBindStrategy:   util.PtrString(""),
		EipGroupUnbindStrategy: util.PtrString(""),
		EipGroupIdList:         []*string{},
		Increase: &as.EipGroupIncrease{
			Enabled:  util.PtrBool(false),
			Strategy: util.PtrString(""),
		},
		Decrease: &as.EipGroupDecrease{
			Enabled: util.PtrBool(false),
		},
		Bandwidth: &as.EipGroupBandwidth{
			Max:      util.PtrInt32(int32(0)),
			Min:      util.PtrInt32(int32(0)),
			Standard: util.PtrInt32(int32(0)),
		},
	}
	Billing := &as.BillingInfo{
		PaymentTiming: util.PtrString(""),
		Reservation: &as.Reservation{
			ReservationLengthInMonth: util.PtrInt32(int32(0)),
		},
	}
	HealthCheck := &as.HealthCheckConfig{
		HealthCheckInterval: util.PtrInt32(int32(0)),
		GraceTime:           util.PtrInt32(int32(0)),
	}
	AssignTagInfo := &as.AssignTagInfo{
		ResourceId:  util.PtrString(""),
		RelationTag: util.PtrBool(false),
		Tags:        []*as.TagInfo{},
	}
	CmdConfig := &as.CmdConfig{
		HasDecreaseCmd: util.PtrBool(false),
		DecCmdStrategy: util.PtrString(""),
		DecCmdData:     util.PtrString(""),
		DecCmdTimeout:  util.PtrInt32(int32(0)),
		DecCmdManual:   util.PtrBool(false),
		HasIncreaseCmd: util.PtrBool(false),
		IncCmdStrategy: util.PtrString(""),
		IncCmdData:     util.PtrString(""),
		IncCmdTimeout:  util.PtrInt32(int32(0)),
		IncCmdManual:   util.PtrBool(false),
	}
	BccNameConfig := &as.BccNameConfig{
		BccName:            util.PtrString(""),
		BccHostname:        util.PtrString(""),
		AutoSeqSuffix:      util.PtrBool(false),
		OpenHostnameDomain: util.PtrBool(false),
	}
	createAsGroupV2Request := &as.CreateAsGroupV2Request{
		GroupName:         util.PtrString(""),
		ZoneInfo:          []*as.ZoneInfo{},
		Config:            Config,
		KeypairId:         util.PtrString(""),
		KeypairName:       util.PtrString(""),
		KeepImageLogin:    util.PtrInt32(int32(0)),
		Blb:               []*as.BlbInfo{},
		BlbUnbindWaitTime: util.PtrInt32(int32(0)),
		Nodes:             []*as.NodeInfo{},
		Eip:               Eip,
		EipConfig:         EipConfig,
		Billing:           Billing,
		Rds:               []*string{},
		Scs:               []*string{},
		HealthCheck:       HealthCheck,
		ExpansionStrategy: util.PtrString(""),
		ShrinkageStrategy: util.PtrString(""),
		AssignTagInfo:     AssignTagInfo,
		CmdConfig:         CmdConfig,
		BccNameConfig:     BccNameConfig,
	}
	result, err := client.CreateAsGroupV2(createAsGroupV2Request)
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
