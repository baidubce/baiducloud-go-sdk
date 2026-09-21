package dbscsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/dbsc"
)

func GetPostgresqlSlowLogTemplate() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := dbsc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getPostgresqlSlowLogTemplateRequest := &dbsc.GetPostgresqlSlowLogTemplateRequest{
		AppId:          util.PtrString(""),
		NodeId:         util.PtrString(""),
		Start:          util.PtrString(""),
		End:            util.PtrString(""),
		Users:          util.PtrString(""),
		DbNames:        util.PtrString(""),
		ClientIps:      util.PtrString(""),
		FingerprintMd5: util.PtrString(""),
		Page:           util.PtrInt32(int32(0)),
		PageSize:       util.PtrInt32(int32(0)),
		OrderBy:        util.PtrString(""),
		Order:          util.PtrString(""),
	}
	result, err := client.GetPostgresqlSlowLogTemplate(getPostgresqlSlowLogTemplateRequest)
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
