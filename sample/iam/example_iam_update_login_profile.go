package iamsample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/iam"
)

func UpdateLoginProfile() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := iam.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateLoginProfileRequest := &iam.UpdateLoginProfileRequest{
		UserName:          util.PtrString(""),
		Password:          util.PtrString(""),
		NeedResetPassword: util.PtrBool(false),
		EnabledLogin:      util.PtrBool(false),
		EnabledLoginMfa:   util.PtrBool(false),
		LoginMfaType:      util.PtrString(""),
		ThirdPartyType:    util.PtrString(""),
		ThirdPartyAccount: util.PtrString(""),
	}
	err = client.UpdateLoginProfile(updateLoginProfileRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
