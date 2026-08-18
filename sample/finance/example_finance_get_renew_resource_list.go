package financesample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/finance"
)

func GetRenewResourceList() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := finance.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	getRenewResourceListRequest := &finance.GetRenewResourceListRequest{
		QueryAccountId:     util.PtrString(""),
		ServiceType:        util.PtrString(""),
		Region:             util.PtrString(""),
		ExpiredDays:        util.PtrInt32(int32(0)),
		ShortOrInstanceIds: []*string{},
		PageNo:             util.PtrInt32(int32(0)),
		PageSize:           util.PtrInt32(int32(0)),
	}
	result, err := client.GetRenewResourceList(getRenewResourceListRequest)
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
