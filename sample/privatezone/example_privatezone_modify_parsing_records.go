package privatezonesample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/privatezone"
)

func ModifyParsingRecords() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := privatezone.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	modifyParsingRecordsRequest := &privatezone.ModifyParsingRecordsRequest{
		RecordId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Rr:          util.PtrString(""),
		Value:       util.PtrString(""),
		Type:        util.PtrString(""),
		Ttl:         util.PtrInt32(int32(0)),
		Priority:    util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	err = client.ModifyParsingRecords(modifyParsingRecordsRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
