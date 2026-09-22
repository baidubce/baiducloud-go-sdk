package scssample

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/services/scs"
)

func CreateHotGroup() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := scs.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	Leader := &scs.Leader{
		GroupName:         util.PtrString(""),
		LeaderId:          util.PtrString(""),
		LeaderRegion:      util.PtrString(""),
		ClusterName:       util.PtrString(""),
		ClusterShowId:     util.PtrString(""),
		Region:            util.PtrString(""),
		Status:            util.PtrString(""),
		TotalCapacityInGB: util.PtrFloat32(float32(0)),
		UsedCapacityInGB:  util.PtrInt32(int32(0)),
		ShardNum:          util.PtrInt32(int32(0)),
		Flavor:            util.PtrInt32(int32(0)),
		QpsWrite:          util.PtrInt64(int64(0)),
		QpsRead:           util.PtrInt64(int64(0)),
		StaleReadable:     util.PtrBool(false),
		ForbidWrite:       util.PtrInt32(int32(0)),
		AvailabilityZone:  util.PtrString(""),
		ExpiredTime:       util.PtrString(""),
	}
	createHotGroupRequest := &scs.CreateHotGroupRequest{
		Leader: Leader,
	}
	result, err := client.CreateHotGroup(createHotGroupRequest)
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
