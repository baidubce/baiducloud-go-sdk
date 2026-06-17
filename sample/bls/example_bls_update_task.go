package blssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/bls"
)

func UpdateTask() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := bls.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Config := &bls.TaskConfig{
		SrcConfig: &bls.SrcConfig{
			SrcType:        util.PtrString(""),
			LogType:        util.PtrString(""),
			SrcDir:         util.PtrString(""),
			MatchedPattern: util.PtrString(""),
			IgnorePattern:  util.PtrString(""),
			TimeFormat:     util.PtrString(""),
			Ttl:            util.PtrInt32(int32(0)),
			UseMultiline:   util.PtrBool(false),
			MultilineRegex: util.PtrString(""),
			RecursiveDir:   util.PtrBool(false),
			ProcessType:    util.PtrString(""),
			ProcessConfig: &bls.ProcessConfig{
				Regex:            util.PtrString(""),
				Separator:        util.PtrString(""),
				CustomSeparator:  util.PtrString(""),
				Quote:            util.PtrString(""),
				KvKeyIndex:       util.PtrInt32(int32(0)),
				KvValueIndex:     util.PtrInt32(int32(0)),
				SampleLog:        util.PtrString(""),
				Keys:             util.PtrString(""),
				DataType:         util.PtrString(""),
				DiscardOnFailure: util.PtrBool(false),
				KeepOriginal:     util.PtrBool(false),
			},
			LogTime:        util.PtrString(""),
			TimestampKey:   util.PtrString(""),
			DateFormat:     util.PtrString(""),
			FilterExpr:     util.PtrString(""),
			AdditionConfig: nil,
			MetaEnv:        []*string{},
			MetaLabel:      []*string{},
			MetaContainer:  []*string{},
			MetaToFields:   util.PtrBool(false),
			HarvesterLimit: util.PtrInt32(int32(0)),
		},
		DestConfig: &bls.DestConfig{
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
		},
	}
	updateTaskRequest := &bls.UpdateTaskRequest{
		TaskId: util.PtrString(""),
		Name:   util.PtrString(""),
		Config: Config,
		Hosts:  []*bls.Host{},
		Tags:   []*bls.Tag{},
	}
	err = client.UpdateTask(updateTaskRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
