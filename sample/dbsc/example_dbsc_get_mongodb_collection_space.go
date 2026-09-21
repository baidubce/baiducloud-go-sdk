package dbscsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/dbsc"
)

func GetMongodbCollectionSpace() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := dbsc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getMongodbCollectionSpaceRequest := &dbsc.GetMongodbCollectionSpaceRequest{
		AppId:      util.PtrString(""),
		NodeId:     util.PtrString(""),
		Database:   util.PtrString(""),
		Collection: util.PtrString(""),
		OrderBy:    util.PtrString(""),
		Order:      util.PtrString(""),
		Page:       util.PtrInt32(int32(0)),
		PageSize:   util.PtrInt32(int32(0)),
	}
	result, err := client.GetMongodbCollectionSpace(getMongodbCollectionSpaceRequest)
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
