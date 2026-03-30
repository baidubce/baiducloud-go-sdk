package vpc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	VPC_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	VPC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_BatchCreateSslVpnUsers(t *testing.T) {
	batchCreateSslVpnUsersRequest := &BatchCreateSslVpnUsersRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		SslVpnUsers: []*SslVpnUser{},
	}
	result := &BatchCreateSslVpnUsersResponse{}
	result, err := VPC_CLIENT.BatchCreateSslVpnUsers(batchCreateSslVpnUsersRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BindEip(t *testing.T) {
	bindEipRequest := &BindEipRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Eip:         util.PtrString(""),
	}
	err := VPC_CLIENT.BindEip(bindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CloseVpcRelay(t *testing.T) {
	closeVpcRelayRequest := &CloseVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.CloseVpcRelay(closeVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateGatewayLimitRules(t *testing.T) {
	createGatewayLimitRulesRequest := &CreateGatewayLimitRulesRequest{
		ClientToken:    util.PtrString(""),
		IpVersion:      util.PtrString(""),
		Name:           util.PtrString(""),
		Description:    util.PtrString(""),
		ServiceType:    util.PtrString(""),
		SubServiceType: util.PtrString(""),
		PeerRegion:     util.PtrString(""),
		ResourceId:     util.PtrString(""),
		Direction:      util.PtrString(""),
		Cidr:           util.PtrString(""),
		Bandwidth:      util.PtrInt32(int32(0)),
	}
	result := &CreateGatewayLimitRulesResponse{}
	result, err := VPC_CLIENT.CreateGatewayLimitRules(createGatewayLimitRulesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateIpReserved(t *testing.T) {
	createIpReservedRequest := &CreateIpReservedRequest{
		ClientToken: util.PtrString(""),
		SubnetId:    util.PtrString(""),
		IpCidr:      util.PtrString(""),
		IpVersion:   util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	result := &CreateIpReservedResponse{}
	result, err := VPC_CLIENT.CreateIpReserved(createIpReservedRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateSslVpnServer(t *testing.T) {
	createSslVpnServerRequest := &CreateSslVpnServerRequest{
		VpnId:            util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SslVpnServerName: util.PtrString(""),
		InterfaceType:    util.PtrString(""),
		LocalSubnets:     []*string{},
		RemoteSubnet:     util.PtrString(""),
		ClientDns:        util.PtrString(""),
	}
	result := &CreateSslVpnServerResponse{}
	result, err := VPC_CLIENT.CreateSslVpnServer(createSslVpnServerRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateSubnet(t *testing.T) {
	createSubnetRequest := &CreateSubnetRequest{
		ClientToken:      util.PtrString(""),
		Name:             util.PtrString(""),
		EnableIpv6:       util.PtrBool(false),
		ZoneName:         util.PtrString(""),
		Cidr:             util.PtrString(""),
		VpcId:            util.PtrString(""),
		VpcSecondaryCidr: util.PtrString(""),
		SubnetType:       util.PtrString(""),
		Description:      util.PtrString(""),
		Tags:             []*TagModel{},
	}
	result := &CreateSubnetResponse{}
	result, err := VPC_CLIENT.CreateSubnet(createSubnetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateUserGateway(t *testing.T) {
	createUserGatewayRequest := &CreateUserGatewayRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Ip:          util.PtrString(""),
		Description: util.PtrString(""),
	}
	result := &CreateUserGatewayResponse{}
	result, err := VPC_CLIENT.CreateUserGateway(createUserGatewayRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateVpc(t *testing.T) {
	createVpcRequest := &CreateVpcRequest{
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Cidr:        util.PtrString(""),
		EnableIpv6:  util.PtrBool(false),
		Tags:        []*TagModel{},
	}
	result := &CreateVpcResponse{}
	result, err := VPC_CLIENT.CreateVpc(createVpcRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateVpn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	createVpnRequest := &CreateVpnRequest{
		ClientToken:     util.PtrString(""),
		VpcId:           util.PtrString(""),
		SubnetId:        util.PtrString(""),
		VpnName:         util.PtrString(""),
		Type:            util.PtrString(""),
		Description:     util.PtrString(""),
		Eip:             util.PtrString(""),
		Tags:            []*TagModel{},
		ResourceGroupId: util.PtrString(""),
		Billing:         Billing,
		MaxConnection:   util.PtrInt32(int32(0)),
		DeleteProtect:   util.PtrBool(false),
	}
	result := &CreateVpnResponse{}
	result, err := VPC_CLIENT.CreateVpn(createVpnRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateVpnTunnel(t *testing.T) {
	IkeConfig := &IkeConfig{
		IkeVersion:  util.PtrString(""),
		IkeMode:     util.PtrString(""),
		IkeEncAlg:   util.PtrString(""),
		IkeAuthAlg:  util.PtrString(""),
		IkePfs:      util.PtrString(""),
		IkeLifeTime: util.PtrString(""),
	}
	IpsecConfig := &IpsecConfig{
		IpsecEncAlg:   util.PtrString(""),
		IpsecAuthAlg:  util.PtrString(""),
		IpsecPfs:      util.PtrString(""),
		IpsecLifetime: util.PtrString(""),
	}
	createVpnTunnelRequest := &CreateVpnTunnelRequest{
		VpnId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		SecretKey:     util.PtrString(""),
		LocalSubnets:  []*string{},
		CgwId:         util.PtrString(""),
		RemoteSubnets: []*string{},
		Description:   util.PtrString(""),
		VpnConnName:   util.PtrString(""),
		IkeConfig:     IkeConfig,
		IpsecConfig:   IpsecConfig,
	}
	result := &CreateVpnTunnelResponse{}
	result, err := VPC_CLIENT.CreateVpnTunnel(createVpnTunnelRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteGatewayLimitRule(t *testing.T) {
	deleteGatewayLimitRuleRequest := &DeleteGatewayLimitRuleRequest{
		GlrId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteGatewayLimitRule(deleteGatewayLimitRuleRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteIpReserve(t *testing.T) {
	deleteIpReserveRequest := &DeleteIpReserveRequest{
		IpReserveId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteIpReserve(deleteIpReserveRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSslVpnServer(t *testing.T) {
	deleteSslVpnServerRequest := &DeleteSslVpnServerRequest{
		VpnId:          util.PtrString(""),
		SslVpnServerId: util.PtrString(""),
		ClientToken:    util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSslVpnServer(deleteSslVpnServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSslVpnUser(t *testing.T) {
	deleteSslVpnUserRequest := &DeleteSslVpnUserRequest{
		VpnId:       util.PtrString(""),
		UserId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSslVpnUser(deleteSslVpnUserRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteSubnet(t *testing.T) {
	deleteSubnetRequest := &DeleteSubnetRequest{
		SubnetId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteSubnet(deleteSubnetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteUserGateway(t *testing.T) {
	deleteUserGatewayRequest := &DeleteUserGatewayRequest{
		CgwId: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteUserGateway(deleteUserGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteVpc(t *testing.T) {
	deleteVpcRequest := &DeleteVpcRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteVpc(deleteVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteVpnTunnel(t *testing.T) {
	deleteVpnTunnelRequest := &DeleteVpnTunnelRequest{
		VpnConnId:   util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteVpnTunnel(deleteVpnTunnelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetVpcResourceIpInfo(t *testing.T) {
	getVpcResourceIpInfoRequest := &GetVpcResourceIpInfoRequest{
		VpcId:        util.PtrString(""),
		SubnetId:     util.PtrString(""),
		ResourceType: util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &GetVpcResourceIpInfoResponse{}
	result, err := VPC_CLIENT.GetVpcResourceIpInfo(getVpcResourceIpInfoRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListIpReserve(t *testing.T) {
	listIpReserveRequest := &ListIpReserveRequest{
		SubnetId: util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &ListIpReserveResponse{}
	result, err := VPC_CLIENT.ListIpReserve(listIpReserveRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ModifyGatewayLimitRules(t *testing.T) {
	modifyGatewayLimitRulesRequest := &ModifyGatewayLimitRulesRequest{
		GlrId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		Bandwidth:   util.PtrInt32(int32(0)),
	}
	err := VPC_CLIENT.ModifyGatewayLimitRules(modifyGatewayLimitRulesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_OpenVpcRelay(t *testing.T) {
	openVpcRelayRequest := &OpenVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.OpenVpcRelay(openVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySpecifiedSubnet(t *testing.T) {
	querySpecifiedSubnetRequest := &QuerySpecifiedSubnetRequest{
		SubnetId: util.PtrString(""),
	}
	result := &QuerySpecifiedSubnetResponse{}
	result, err := VPC_CLIENT.QuerySpecifiedSubnet(querySpecifiedSubnetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySpecifiedVpc(t *testing.T) {
	querySpecifiedVpcRequest := &QuerySpecifiedVpcRequest{
		VpcId: util.PtrString(""),
	}
	result := &QuerySpecifiedVpcResponse{}
	result, err := VPC_CLIENT.QuerySpecifiedVpc(querySpecifiedVpcRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySslVpnServer(t *testing.T) {
	querySslVpnServerRequest := &QuerySslVpnServerRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result := &QuerySslVpnServerResponse{}
	result, err := VPC_CLIENT.QuerySslVpnServer(querySslVpnServerRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySslVpnUsers(t *testing.T) {
	querySslVpnUsersRequest := &QuerySslVpnUsersRequest{
		VpnId:    util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
		UserName: util.PtrString(""),
	}
	result := &QuerySslVpnUsersResponse{}
	result, err := VPC_CLIENT.QuerySslVpnUsers(querySslVpnUsersRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QuerySubnetList(t *testing.T) {
	querySubnetListRequest := &QuerySubnetListRequest{
		Marker:     util.PtrString(""),
		MaxKeys:    util.PtrInt32(int32(0)),
		VpcId:      util.PtrString(""),
		ZoneName:   util.PtrString(""),
		SubnetType: util.PtrString(""),
		SubnetIds:  util.PtrString(""),
	}
	result := &QuerySubnetListResponse{}
	result, err := VPC_CLIENT.QuerySubnetList(querySubnetListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryVpcIntranetIp(t *testing.T) {
	queryVpcIntranetIpRequest := &QueryVpcIntranetIpRequest{
		VpcId:              util.PtrString(""),
		PrivateIpAddresses: []*string{},
		PrivateIpRange:     util.PtrString(""),
	}
	result := &QueryVpcIntranetIpResponse{}
	result, err := VPC_CLIENT.QueryVpcIntranetIp(queryVpcIntranetIpRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryVpcList(t *testing.T) {
	queryVpcListRequest := &QueryVpcListRequest{
		Marker:    util.PtrString(""),
		MaxKeys:   util.PtrInt32(int32(0)),
		IsDefault: util.PtrBool(false),
		VpcIds:    util.PtrString(""),
	}
	result := &QueryVpcListResponse{}
	result, err := VPC_CLIENT.QueryVpcList(queryVpcListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryVpnList(t *testing.T) {
	queryVpnListRequest := &QueryVpnListRequest{
		VpcId:   util.PtrString(""),
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
		Eip:     util.PtrString(""),
		Type:    util.PtrString(""),
	}
	result := &QueryVpnListResponse{}
	result, err := VPC_CLIENT.QueryVpnList(queryVpnListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ReleaseVpn(t *testing.T) {
	releaseVpnRequest := &ReleaseVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.ReleaseVpn(releaseVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_RenewVpn(t *testing.T) {
	Billing := &Billing{
		PaymentTiming: util.PtrString(""),
		Reservation: &Reservation{
			ReservationLength:   util.PtrInt32(int32(0)),
			ReservationTimeUnit: util.PtrString(""),
		},
	}
	renewVpnRequest := &RenewVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Billing:     Billing,
	}
	err := VPC_CLIENT.RenewVpn(renewVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SearchForVpnDetails(t *testing.T) {
	searchForVpnDetailsRequest := &SearchForVpnDetailsRequest{
		VpnId: util.PtrString(""),
	}
	result := &SearchForVpnDetailsResponse{}
	result, err := VPC_CLIENT.SearchForVpnDetails(searchForVpnDetailsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SearchVpnTunnel(t *testing.T) {
	searchVpnTunnelRequest := &SearchVpnTunnelRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	result := &SearchVpnTunnelResponse{}
	result, err := VPC_CLIENT.SearchVpnTunnel(searchVpnTunnelRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UnbindEip(t *testing.T) {
	unbindEipRequest := &UnbindEipRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.UnbindEip(unbindEipRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSslVpnServer(t *testing.T) {
	updateSslVpnServerRequest := &UpdateSslVpnServerRequest{
		VpnId:            util.PtrString(""),
		SslVpnServerId:   util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SslVpnServerName: util.PtrString(""),
		LocalSubnets:     []*string{},
		RemoteSubnet:     util.PtrString(""),
		ClientDns:        util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateSslVpnServer(updateSslVpnServerRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSslVpnUsers(t *testing.T) {
	updateSslVpnUsersRequest := &UpdateSslVpnUsersRequest{
		VpnId:       util.PtrString(""),
		UserId:      util.PtrString(""),
		ClientToken: util.PtrString(""),
		ClientName:  util.PtrString(""),
		Password:    util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateSslVpnUsers(updateSslVpnUsersRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateSubnet(t *testing.T) {
	updateSubnetRequest := &UpdateSubnetRequest{
		SubnetId:    util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
		EnableIpv6:  util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdateSubnet(updateSubnetRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateUserGateway(t *testing.T) {
	updateUserGatewayRequest := &UpdateUserGatewayRequest{
		CgwId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		Name:        util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateUserGateway(updateUserGatewayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpc(t *testing.T) {
	updateVpcRequest := &UpdateVpcRequest{
		VpcId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		Name:          util.PtrString(""),
		Description:   util.PtrString(""),
		EnableIpv6:    util.PtrBool(false),
		SecondaryCidr: []*string{},
	}
	err := VPC_CLIENT.UpdateVpc(updateVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpn(t *testing.T) {
	updateVpnRequest := &UpdateVpnRequest{
		VpnId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
		VpnName:     util.PtrString(""),
		Description: util.PtrString(""),
	}
	err := VPC_CLIENT.UpdateVpn(updateVpnRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpnReleaseProtection(t *testing.T) {
	updateVpnReleaseProtectionRequest := &UpdateVpnReleaseProtectionRequest{
		VpnId:         util.PtrString(""),
		ClientToken:   util.PtrString(""),
		DeleteProtect: util.PtrBool(false),
	}
	err := VPC_CLIENT.UpdateVpnReleaseProtection(updateVpnReleaseProtectionRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateVpnTunnel(t *testing.T) {
	IkeConfig := &IkeConfig{
		IkeVersion:  util.PtrString(""),
		IkeMode:     util.PtrString(""),
		IkeEncAlg:   util.PtrString(""),
		IkeAuthAlg:  util.PtrString(""),
		IkePfs:      util.PtrString(""),
		IkeLifeTime: util.PtrString(""),
	}
	IpsecConfig := &IpsecConfig{
		IpsecEncAlg:   util.PtrString(""),
		IpsecAuthAlg:  util.PtrString(""),
		IpsecPfs:      util.PtrString(""),
		IpsecLifetime: util.PtrString(""),
	}
	updateVpnTunnelRequest := &UpdateVpnTunnelRequest{
		VpnConnId:     util.PtrString(""),
		ClientToken:   util.PtrString(""),
		VpnId:         util.PtrString(""),
		SecretKey:     util.PtrString(""),
		LocalSubnets:  []*string{},
		CgwId:         util.PtrString(""),
		RemoteSubnets: []*string{},
		Description:   util.PtrString(""),
		VpnConnName:   util.PtrString(""),
		IkeConfig:     IkeConfig,
		IpsecConfig:   IpsecConfig,
	}
	err := VPC_CLIENT.UpdateVpnTunnel(updateVpnTunnelRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserGatewayDetails(t *testing.T) {
	userGatewayDetailsRequest := &UserGatewayDetailsRequest{
		CgwId: util.PtrString(""),
	}
	result := &UserGatewayDetailsResponse{}
	result, err := VPC_CLIENT.UserGatewayDetails(userGatewayDetailsRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserGatewayList(t *testing.T) {
	userGatewayListRequest := &UserGatewayListRequest{
		Marker:  util.PtrString(""),
		MaxKeys: util.PtrInt32(int32(0)),
	}
	result := &UserGatewayListResponse{}
	result, err := VPC_CLIENT.UserGatewayList(userGatewayListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ViewGatewayLimitRules(t *testing.T) {
	viewGatewayLimitRulesRequest := &ViewGatewayLimitRulesRequest{
		ServiceType: util.PtrString(""),
		Name:        util.PtrString(""),
		GlrId:       util.PtrString(""),
		ResourceId:  util.PtrString(""),
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrString(""),
	}
	result := &ViewGatewayLimitRulesResponse{}
	result, err := VPC_CLIENT.ViewGatewayLimitRules(viewGatewayLimitRulesRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
