package bci

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
	BCI_CLIENT *Client
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

	BCI_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_BatchDeleteImageCaches(t *testing.T) {
	batch - delete - image - cachesRequest := &BatchDeleteImageCachesRequest{
		ImageCacheIds: []*string{},
	}
	err := BCI_CLIENT.BatchDeleteImageCaches(batch - delete - image - cachesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_BatchDeleteInstances(t *testing.T) {
	batch - delete - instancesRequest := &BatchDeleteInstancesRequest{
		InstanceIds:        []*string{},
		RelatedReleaseFlag: util.PtrBool(false),
	}
	err := BCI_CLIENT.BatchDeleteInstances(batch - delete - instancesRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateImageCache(t *testing.T) {
	create - image - cacheRequest := &CreateImageCacheRequest{
		ImageCacheName:       util.PtrString(""),
		OriginImages:         []*OriginImage{},
		SubnetId:             util.PtrString(""),
		SecurityGroupId:      util.PtrString(""),
		ZoneName:             util.PtrString(""),
		TemporaryStorageSize: util.PtrInt32(int32(0)),
		NeedEip:              util.PtrBool(false),
		EipIp:                util.PtrString(""),
		AutoMatchImageCache:  util.PtrBool(false),
		ImageRegistrySecrets: []*ImageRegistryCredential{},
	}
	result := &CreateImageCacheResponse{}
	result, err := BCI_CLIENT.CreateImageCache(create - image - cacheRequest)
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
func TestClient_CreateInstance(t *testing.T) {
	Volume := &Volume{
		Nfs:        []*NfsVolume{},
		EmptyDir:   []*EmptyDirVolume{},
		ConfigFile: []*ConfigFileVolume{},
	}
	create - instanceRequest := &CreateInstanceRequest{
		ClientToken:                   util.PtrString(""),
		Name:                          util.PtrString(""),
		ZoneName:                      util.PtrString(""),
		SecurityGroupIds:              []*string{},
		SubnetIds:                     []*string{},
		RestartPolicy:                 util.PtrString(""),
		EipIp:                         util.PtrString(""),
		AutoCreateEip:                 util.PtrBool(false),
		EipName:                       util.PtrString(""),
		EipRouteType:                  util.PtrString(""),
		EipBandwidthInMbps:            util.PtrInt32(int32(0)),
		EipBillingMethod:              util.PtrString(""),
		GpuType:                       util.PtrString(""),
		TerminationGracePeriodSeconds: util.PtrInt64(int64(0)),
		HostName:                      util.PtrString(""),
		Tags:                          []*Tag{},
		ImageRegistryCredentials:      []*ImageRegistryCredential{},
		Containers:                    []*Container{},
		InitContainers:                []*Container{},
		Volume:                        Volume,
	}
	result := &CreateInstanceResponse{}
	result, err := BCI_CLIENT.CreateInstance(create - instanceRequest)
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
func TestClient_DeleteInstance(t *testing.T) {
	delete - instanceRequest := &DeleteInstanceRequest{
		InstanceId:         util.PtrString(""),
		RelatedReleaseFlag: util.PtrBool(false),
	}
	err := BCI_CLIENT.DeleteInstance(delete - instanceRequest)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetInstance(t *testing.T) {
	get - instanceRequest := &GetInstanceRequest{
		InstanceId: util.PtrString(""),
	}
	result := &GetInstanceResponse{}
	result, err := BCI_CLIENT.GetInstance(get - instanceRequest)
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
func TestClient_ListImageCaches(t *testing.T) {
	list - image - cachesRequest := &ListImageCachesRequest{
		PageSize: util.PtrInt32(int32(0)),
		PageNo:   util.PtrInt32(int32(0)),
	}
	result := &ListImageCachesResponse{}
	result, err := BCI_CLIENT.ListImageCaches(list - image - cachesRequest)
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
func TestClient_ListInstances(t *testing.T) {
	list - instancesRequest := &ListInstancesRequest{
		KeywordType: util.PtrString(""),
		Keyword:     util.PtrString(""),
		Marker:      util.PtrString(""),
		MaxKeys:     util.PtrInt32(int32(0)),
	}
	result := &ListInstancesResponse{}
	result, err := BCI_CLIENT.ListInstances(list - instancesRequest)
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
