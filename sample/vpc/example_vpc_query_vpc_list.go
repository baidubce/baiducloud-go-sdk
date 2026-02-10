package vpcsample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func QueryVpcList() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := vpc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	queryVpcListRequest := &vpc.QueryVpcListRequest{
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		IsDefault: util.PtrBool(false),
		VpcIds:    util.PtrString(""),
	}
	result := &vpc.QueryVpcListResponse{}
	result, err = client.QueryVpcList(queryVpcListRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
}
