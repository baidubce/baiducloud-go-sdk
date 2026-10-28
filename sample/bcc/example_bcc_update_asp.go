package bccsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bcc"
)

func UpdateAsp() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := bcc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateAspRequest := &bcc.UpdateAspRequest{
		AspId:          util.PtrString(""),
		Name:           util.PtrString(""),
		TimePoints:     []*string{},
		RepeatWeekdays: []*string{},
		RetentionDays:  util.PtrString(""),
	}
	err = client.UpdateAsp(updateAspRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
