package eipsample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/eip"
)

func SharedDataPackageInquiry() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := eip.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	sharedDataPackageInquiryRequest := &eip.SharedDataPackageInquiryRequest{
		ClientToken:       util.PtrString(""),
		ReservationLength: util.PtrInt32(int32(0)),
		Capacity:          util.PtrString(""),
		DeductPolicy:      util.PtrString(""),
		PackageType:       util.PtrString(""),
	}
	result := &eip.SharedDataPackageInquiryResponse{}
	result, err = client.SharedDataPackageInquiry(sharedDataPackageInquiryRequest)
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
