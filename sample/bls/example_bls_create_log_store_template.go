package blssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bls"
)

func CreateLogStoreTemplate() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bls.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Template := &bls.Template{
		Retention:             util.PtrInt32(int32(0)),
		ShardCount:            util.PtrInt32(int32(0)),
		DisableShardAutoSplit: util.PtrBool(false),
		MaxShardCount:         util.PtrInt32(int32(0)),
		EnableHotRetention:    util.PtrBool(false),
		HotRetention:          util.PtrInt32(int32(0)),
		Index: &bls.Index{
			Fulltext:       util.PtrBool(false),
			CaseSensitive:  util.PtrBool(false),
			IncludeChinese: util.PtrBool(false),
			Separators:     util.PtrString(""),
			Fields:         &map[string]*bls.Field{},
		},
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		CreatedTimestamp: util.PtrString(""),
		UpdatedTimestamp: util.PtrString(""),
	}
	createLogStoreTemplateRequest := &bls.CreateLogStoreTemplateRequest{
		Name:             util.PtrString(""),
		ProjectPatterns:  []*string{},
		LogstorePatterns: []*string{},
		Priority:         util.PtrInt32(int32(0)),
		Template:         Template,
	}
	result, err := client.CreateLogStoreTemplate(createLogStoreTemplateRequest)
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
