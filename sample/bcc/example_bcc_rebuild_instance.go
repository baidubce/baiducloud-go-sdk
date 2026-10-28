package bccsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func RebuildInstance() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := bcc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	rebuildInstanceRequest := &bcc.RebuildInstanceRequest{
		InstanceId:        util.PtrString(""),
		ImageId:           util.PtrString(""),
		KeepImageLogin:    util.PtrBool(false),
		IsPreserveData:    util.PtrBool(false),
		AdminPass:         util.PtrString(""),
		IsOpenHostEye:     util.PtrBool(false),
		SysRootSize:       util.PtrInt32(int32(0)),
		KeypairId:         util.PtrString(""),
		DataPartitionType: util.PtrString(""),
		RootPartitionType: util.PtrString(""),
		RaidId:            util.PtrString(""),
		UserData:          util.PtrString(""),
		UseLastUserData:   util.PtrBool(false),
		CleanLastUserData: util.PtrBool(false),
	}
	err = client.RebuildInstance(rebuildInstanceRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
