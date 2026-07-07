package agentidentitysample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/agentidentity"
)

func CreateOauth2Client() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := agentidentity.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createOauth2ClientRequest := &agentidentity.CreateOauth2ClientRequest{
		UserPoolId:      util.PtrString(""),
		Name:            util.PtrString(""),
		Description:     util.PtrString(""),
		ClientType:      util.PtrString(""),
		RedirectUris:    []*string{},
		GrantTypes:      []*string{},
		Scopes:          []*string{},
		AccessTokenTtl:  util.PtrInt32(int32(0)),
		RefreshTokenTtl: util.PtrInt32(int32(0)),
	}
	err = client.CreateOauth2Client(createOauth2ClientRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
