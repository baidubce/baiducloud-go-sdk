package cpromsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/cprom"
	"os"
)

func RemoteWrite() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := cprom.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	body, err := os.Open("path/to/your/file")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer body.Close()
	remoteWriteRequest := &cprom.RemoteWriteRequest{
		RemoteWriteUrl:  util.PtrString(""),
		ContentType:     util.PtrString(""),
		ContentEncoding: util.PtrString(""),
		InstanceId:      util.PtrString(""),
		Authorization:   util.PtrString(""),
		Body:            body,
	}
	err = client.RemoteWrite(remoteWriteRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
