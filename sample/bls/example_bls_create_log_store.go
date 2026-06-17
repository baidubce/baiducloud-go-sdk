package blssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bls"
)

func CreateLogStore() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bls.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Index := &bls.Index{
		Fulltext:       util.PtrBool(false),
		CaseSensitive:  util.PtrBool(false),
		IncludeChinese: util.PtrBool(false),
		Separators:     util.PtrString(""),
		Fields:         &map[string]*bls.Field{},
	}
	createLogStoreRequest := &bls.CreateLogStoreRequest{
		Project:               util.PtrString(""),
		LogStoreName:          util.PtrString(""),
		Retention:             util.PtrInt32(int32(0)),
		Tags:                  []*bls.Tag{},
		Index:                 Index,
		ShardCount:            util.PtrInt32(int32(0)),
		MaxShardCount:         util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		IndexEnabled:          util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
		EnableHotRetention:    util.PtrBool(false),
	}
	result, err := client.CreateLogStore(createLogStoreRequest)
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
