package vpc

import (
	"bytes"
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

func TestClient_CloseVpcRelay(t *testing.T) {
	closeVpcRelayRequest := &CloseVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.CloseVpcRelay(closeVpcRelayRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateAReservedNetworkSegment(t *testing.T) {
	createAReservedNetworkSegmentRequest := &CreateAReservedNetworkSegmentRequest{
		ClientToken: util.PtrString(""),
		SubnetId:    util.PtrString(""),
		IpCidr:      util.PtrString(""),
		IpVersion:   util.PtrInt32(int32(0)),
		Description: util.PtrString(""),
	}
	result := &CreateAReservedNetworkSegmentResponse{}
	result, err := VPC_CLIENT.CreateAReservedNetworkSegment(createAReservedNetworkSegmentRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_DeleteReservedNetworkSegment(t *testing.T) {
	deleteReservedNetworkSegmentRequest := &DeleteReservedNetworkSegmentRequest{
		IpReserveId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteReservedNetworkSegment(deleteReservedNetworkSegmentRequest)
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
func TestClient_DeleteVpc(t *testing.T) {
	deleteVpcRequest := &DeleteVpcRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.DeleteVpc(deleteVpcRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_EnableVpcRelay(t *testing.T) {
	enableVpcRelayRequest := &EnableVpcRelayRequest{
		VpcId:       util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err := VPC_CLIENT.EnableVpcRelay(enableVpcRelayRequest)
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheIpAddressesOccupiedByProductsWithinVpc(t *testing.T) {
	queryTheIpAddressesOccupiedByProductsWithinVpcRequest := &QueryTheIpAddressesOccupiedByProductsWithinVpcRequest{
		VpcId:        util.PtrString(""),
		SubnetId:     util.PtrString(""),
		ResourceType: util.PtrString(""),
		PageNo:       util.PtrInt32(int32(0)),
		PageSize:     util.PtrInt32(int32(0)),
	}
	result := &QueryTheIpAddressesOccupiedByProductsWithinVpcResponse{}
	result, err := VPC_CLIENT.QueryTheIpAddressesOccupiedByProductsWithinVpc(queryTheIpAddressesOccupiedByProductsWithinVpcRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_QueryTheReservedNetworkSegmentList(t *testing.T) {
	queryTheReservedNetworkSegmentListRequest := &QueryTheReservedNetworkSegmentListRequest{
		SubnetId: util.PtrString(""),
		Marker:   util.PtrString(""),
		MaxKeys:  util.PtrInt32(int32(0)),
	}
	result := &QueryTheReservedNetworkSegmentListResponse{}
	result, err := VPC_CLIENT.QueryTheReservedNetworkSegmentList(queryTheReservedNetworkSegmentListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
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
