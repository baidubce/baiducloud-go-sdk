package blssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bls"
)

func CreateLogShipper() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bls.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	DestConfig := &bls.DestConfig{
		BOSPath:                  util.PtrString(""),
		PartitionFormatTS:        util.PtrString(""),
		PartitionFormatLogStream: util.PtrBool(false),
		MaxObjectSize:            util.PtrInt32(int32(0)),
		CompressType:             util.PtrString(""),
		DeliverInterval:          util.PtrInt32(int32(0)),
		StorageFormat:            util.PtrString(""),
		CsvHeadline:              util.PtrBool(false),
		CsvDelimiter:             util.PtrString(""),
		CsvQuote:                 util.PtrString(""),
		NullIdentifier:           util.PtrString(""),
		SelectedColumnName:       util.PtrString(""),
		SelectedColumnType:       util.PtrString(""),
		FieldsName:               []*string{},
		FieldsType:               []*string{},
		ShipperType:              util.PtrString(""),
		KafkaConfig:              util.PtrString(""),
		DestType:                 util.PtrString(""),
		LogStore:                 util.PtrString(""),
		RateLimit:                util.PtrInt32(int32(0)),
		ClientCount:              util.PtrInt32(int32(0)),
	}
	createLogShipperRequest := &bls.CreateLogShipperRequest{
		Project:        util.PtrString(""),
		LogStoreName:   util.PtrString(""),
		LogShipperName: util.PtrString(""),
		StartTime:      util.PtrString(""),
		DestType:       util.PtrString(""),
		DestConfig:     DestConfig,
	}
	result, err := client.CreateLogShipper(createLogShipperRequest)
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
