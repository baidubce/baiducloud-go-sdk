package rdssample

import (
	"fmt"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/services/rds"
)

func QueryTaskList() {
	endpoint := "Your Endpoint"

	// ==== AK/SK 鉴权 ====
	ak, sk := "Your Ak", "Your Sk"
	client, err := rds.NewClient(ak, sk, endpoint)

	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	queryTaskListRequest := &rds.QueryTaskListRequest{
		PageSize:     util.PtrString(""),
		PageNo:       util.PtrString(""),
		InstanceId:   util.PtrString(""),
		InstanceName: util.PtrString(""),
		TaskId:       util.PtrInt32(int32(0)),
		TaskType:     util.PtrString(""),
		TaskStatus:   util.PtrString(""),
		StartTime:    util.PtrString(""),
		EndTime:      util.PtrString(""),
	}
	err = client.QueryTaskList(queryTaskListRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}
