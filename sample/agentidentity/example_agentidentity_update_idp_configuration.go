package agentidentitysample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/agentidentity"
)

func UpdateIdpConfiguration() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := agentidentity.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	updateIdpConfigurationRequest := &agentidentity.UpdateIdpConfigurationRequest{
		UserPoolId:            util.PtrString(""),
		Id:                    util.PtrString(""),
		Name:                  util.PtrString(""),
		ClientId:              util.PtrString(""),
		ClientSecret:          util.PtrString(""),
		AuthorizationEndpoint: util.PtrString(""),
		TokenEndpoint:         util.PtrString(""),
		UserinfoEndpoint:      util.PtrString(""),
		Scopes:                []*string{},
		UserIdClaim:           util.PtrString(""),
		DisplayNameClaim:      util.PtrString(""),
		AutoCreateUser:        util.PtrBool(false),
	}
	err = client.UpdateIdpConfiguration(updateIdpConfigurationRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
