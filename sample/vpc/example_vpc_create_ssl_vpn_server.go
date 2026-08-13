package vpcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/vpc"
)

func CreateSslVpnServer() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := vpc.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createSslVpnServerRequest := &vpc.CreateSslVpnServerRequest{
		VpnId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SslVpnServerName: util.PtrString(""),
		InterfaceType:    util.PtrString(""),
		LocalSubnets:     []*string{},
		RemoteSubnet:     util.PtrString(""),
		ClientDns:        util.PtrString(""),
	}
	result, err := client.CreateSslVpnServer(createSslVpnServerRequest)
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
