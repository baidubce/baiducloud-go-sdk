package ccesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cce"
)

func DeleteNodesClusterScalingV2() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := cce.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	DeleteOption := &cce.DeleteOption{
		DeleteResource:    util.PtrBool(false),
		DeleteCDSSnapshot: util.PtrBool(false),
		MoveOut:           util.PtrBool(false),
	}
	deleteNodesClusterScalingV2Request := &cce.DeleteNodesClusterScalingV2Request{
		ClusterID:    util.PtrString(""),
		DeleteOption: DeleteOption,
		InstanceIDs:  []*string{},
		ScaleDown:    util.PtrBool(false),
	}
	result, err := client.DeleteNodesClusterScalingV2(deleteNodesClusterScalingV2Request)
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
