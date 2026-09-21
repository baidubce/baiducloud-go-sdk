package dbscsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/dbsc"
)

func UpdateMysqlRateLimitTask() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := dbsc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateMysqlRateLimitTaskRequest := &dbsc.UpdateMysqlRateLimitTaskRequest{
		FilterId:    util.PtrInt32(int32(0)),
		AppId:       util.PtrString(""),
		NodeId:      util.PtrString(""),
		FilterKey:   util.PtrString(""),
		FilterLimit: util.PtrInt32(int32(0)),
		FilterType:  util.PtrString(""),
	}
	err = client.UpdateMysqlRateLimitTask(updateMysqlRateLimitTaskRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
