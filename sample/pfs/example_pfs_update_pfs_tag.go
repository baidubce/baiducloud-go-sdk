package pfssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/services/pfs"
)

func UpdatePFSTag() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := pfs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updatePFSTagRequest := &pfs.UpdatePFSTagRequest{
		InstanceId: []*string{},
		Tags:       []*pfs.Tag{},
	}
	err = client.UpdatePFSTag(updatePFSTagRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
