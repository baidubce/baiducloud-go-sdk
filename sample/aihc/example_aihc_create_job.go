package aihcsample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/aihc"
)

func CreateJob() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := aihc.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	JobSpec := &aihc.JobSpec{
		Image: util.PtrString(""),
		ImageConfig: &aihc.ImageConfig{
			Username: util.PtrString(""),
			Password: util.PtrString(""),
		},
		Replicas:    util.PtrInt32(int32(0)),
		Resources:   []*aihc.Resource{},
		Envs:        []*aihc.Env{},
		EnableRDMA:  util.PtrBool(false),
		HostNetwork: util.PtrBool(false),
	}
	TensorboardConfig := &aihc.TensorboardConfig{
		Enable:  util.PtrBool(false),
		LogPath: util.PtrString(""),
	}
	AlertConfig := &aihc.AlertConfig{
		InstanceId:   util.PtrString(""),
		AlertItems:   []*string{},
		AihcFor:      util.PtrString(""),
		NotifyRuleId: util.PtrString(""),
	}
	AdvancedSettings := &aihc.AdvancedSettings{
		RuntimeEnv:            util.PtrString(""),
		SubmitterBackoffLimit: util.PtrInt32(int32(0)),
	}
	createJobRequest := &aihc.CreateJobRequest{
		ResourcePoolId:     util.PtrString(""),
		QueueID:            util.PtrString(""),
		Name:               util.PtrString(""),
		Queue:              util.PtrString(""),
		JobType:            util.PtrString(""),
		JobSpec:            JobSpec,
		Command:            util.PtrString(""),
		Labels:             []*aihc.Label{},
		Priority:           util.PtrString(""),
		Datasources:        []*aihc.DataSource{},
		EnableBccl:         util.PtrBool(false),
		FaultTolerance:     util.PtrBool(false),
		FaultToleranceArgs: util.PtrString(""),
		TensorboardConfig:  TensorboardConfig,
		AlertConfig:        AlertConfig,
		RetentionPeriod:    util.PtrString(""),
		AdvancedSettings:   AdvancedSettings,
	}
	result, err := client.CreateJob(createJobRequest)
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
