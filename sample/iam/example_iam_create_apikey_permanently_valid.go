package iamsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/iam"
)

func CreateApikeyPermanentlyValid() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := iam.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Acl := &iam.ACL{
		Id:                util.PtrString(""),
		Version:           util.PtrString(""),
		AccessControlList: []*iam.ACLEntry{},
	}
	createApikeyPermanentlyValidRequest := &iam.CreateApikeyPermanentlyValidRequest{
		UserId: util.PtrString(""),
		Acl:    Acl,
		Name:   util.PtrString(""),
	}
	result, err := client.CreateApikeyPermanentlyValid(createApikeyPermanentlyValidRequest)
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
