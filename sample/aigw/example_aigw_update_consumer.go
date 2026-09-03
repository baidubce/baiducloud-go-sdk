package aigwsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aigw"
)

func UpdateConsumer() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := aigw.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	CredentialOp := &aigw.CredentialOp{
		Operation:      util.PtrString(""),
		CredentialName: util.PtrString(""),
		Value:          util.PtrString(""),
	}
	CredentialLocation := &aigw.ConsumerCredentialLocation{
		InHeader: util.PtrBool(false),
		InQuery:  util.PtrBool(false),
		KeyNames: []*string{},
	}
	IamCredential := &aigw.IAMCredentialSpec{
		Name:             util.PtrString(""),
		IamApiKeyId:      util.PtrString(""),
		IamTokenIdMasked: util.PtrString(""),
		IamUserId:        util.PtrString(""),
		IamDomainId:      util.PtrString(""),
		ResourceIds:      []*string{},
		InHeader:         util.PtrBool(false),
		InQuery:          util.PtrBool(false),
		KeyNames:         []*string{},
		Status:           util.PtrString(""),
	}
	updateConsumerRequest := &aigw.UpdateConsumerRequest{
		InstanceId:         util.PtrString(""),
		ConsumerId:         util.PtrString(""),
		XRegion:            util.PtrString(""),
		KeyType:            util.PtrString(""),
		Description:        util.PtrString(""),
		RouteNames:         []*string{},
		Tags:               []*aigw.Tag{},
		CredentialOp:       CredentialOp,
		CredentialLocation: CredentialLocation,
		IamCredential:      IamCredential,
	}
	result, err := client.UpdateConsumer(updateConsumerRequest)
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
