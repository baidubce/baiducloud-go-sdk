package pfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func QryL2PolExecLog() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	qryL2PolExecLogRequest := &pfs.QryL2PolExecLogRequest{
		Action:     util.PtrString(""),
		InstanceId: util.PtrString(""),
		PolicyId:   util.PtrString(""),
		StartTime:  util.PtrInt32(int32(0)),
		EndTime:    util.PtrInt32(int32(0)),
	}
	result, err := client.QryL2PolExecLog(qryL2PolExecLogRequest)
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
