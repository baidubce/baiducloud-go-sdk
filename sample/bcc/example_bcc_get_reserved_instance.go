package bccsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func GetReservedInstance() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := bcc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getReservedInstanceRequest := &bcc.GetReservedInstanceRequest{
		Marker:                 util.PtrString(""),
		MaxKeys:                util.PtrInt32(int32(0)),
		ReservedInstanceIds:    []*string{},
		ReservedInstanceName:   util.PtrString(""),
		ZoneName:               util.PtrString(""),
		ReservedInstanceStatus: util.PtrString(""),
		Spec:                   util.PtrString(""),
		OfferingType:           util.PtrString(""),
		OsType:                 util.PtrString(""),
		InstanceId:             util.PtrString(""),
		InstanceName:           util.PtrString(""),
		IsDeduct:               util.PtrBool(false),
		EhcClusterId:           util.PtrString(""),
		SortKey:                util.PtrString(""),
		SortDir:                util.PtrString(""),
		ReservedInstanceSource: util.PtrString(""),
		Scope:                  util.PtrString(""),
	}
	result, err := client.GetReservedInstance(getReservedInstanceRequest)
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
