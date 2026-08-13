package pfssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func CreateFileset() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createFilesetRequest := &pfs.CreateFilesetRequest{
		Action:           util.PtrString(""),
		InstanceId:       util.PtrString(""),
		FilesetName:      util.PtrString(""),
		FilesetPath:      util.PtrString(""),
		BlockQuota:       util.PtrInt32(int32(0)),
		FilesQuota:       util.PtrInt64(int64(0)),
		QpsLimit:         util.PtrInt32(int32(0)),
		BandwidthLimitMb: util.PtrInt32(int32(0)),
	}
	result, err := client.CreateFileset(createFilesetRequest)
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
