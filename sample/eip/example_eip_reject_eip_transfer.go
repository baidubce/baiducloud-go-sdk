package eipsample

import (
    "fmt"
    "github.com/baidubce/baiducloud-go-sdk/services/eip"
    "github.com/baidubce/baiducloud-go-sdk/core/util"

)

func RejectEipTransfer() {
    // 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
    ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
    client, err := eip.NewClient(ak, sk, endpoint)
    if err != nil {
        fmt.Println("create client err:", err)
        return
    }
    rejectEipTransferRequest := &eip.RejectEipTransferRequest{
        Action : util.PtrString(""),
        ClientToken : util.PtrString(""),
        TransferIdList : []*string{},
    }
    err = client.RejectEipTransfer(rejectEipTransferRequest)
    if err != nil {
        // 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
        fmt.Println("request failed:", err)
    }
}
