package dbscsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/dbsc"
)

func GetMongodbDatabaseSpaceTrend() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := dbsc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getMongodbDatabaseSpaceTrendRequest := &dbsc.GetMongodbDatabaseSpaceTrendRequest{
		AppId:      util.PtrString(""),
		Database:   util.PtrString(""),
		Period:     util.PtrInt32(int32(0)),
		NodeId:     util.PtrString(""),
		Start:      util.PtrString(""),
		End:        util.PtrString(""),
		Metrics:    util.PtrString(""),
		Statistics: util.PtrString(""),
	}
	result, err := client.GetMongodbDatabaseSpaceTrend(getMongodbDatabaseSpaceTrendRequest)
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
