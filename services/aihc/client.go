package aihc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "aihc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of aihc service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getBatchStopTrainingTasksV2Uri() string {
	return bce.URI_PREFIX
}
func getCreateDatasetUri() string {
	return bce.URI_PREFIX
}
func getCreateDatasetVersionUri() string {
	return bce.URI_PREFIX
}
func getCreateJobUri() string {
	return bce.URI_PREFIX
}
func getCreateModelUri() string {
	return bce.URI_PREFIX
}
func getCreateModelVersionUri() string {
	return bce.URI_PREFIX
}
func getDeleteDatasetUri() string {
	return bce.URI_PREFIX
}
func getDeleteDatasetVersionUri() string {
	return bce.URI_PREFIX
}
func getDeleteJobUri() string {
	return bce.URI_PREFIX
}
func getDeleteModelUri() string {
	return bce.URI_PREFIX
}
func getDeleteModelVersionUri() string {
	return bce.URI_PREFIX
}
func getDescribeDatasetUri() string {
	return bce.URI_PREFIX
}
func getDescribeDatasetVersionUri() string {
	return bce.URI_PREFIX
}
func getDescribeDatasetVersionsUri() string {
	return bce.URI_PREFIX
}
func getDescribeDatasetsUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobEventsUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobLogsUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobMetricsUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobNodesUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobWebterminalUri() string {
	return bce.URI_PREFIX
}
func getDescribeJobsUri() string {
	return bce.URI_PREFIX
}
func getDescribeModelUri() string {
	return bce.URI_PREFIX
}
func getDescribeModelVersionUri() string {
	return bce.URI_PREFIX
}
func getDescribeModelVersionsUri() string {
	return bce.URI_PREFIX
}
func getDescribeModelsUri() string {
	return bce.URI_PREFIX
}
func getDescribePodEventsUri() string {
	return bce.URI_PREFIX
}
func getModifyDatasetUri() string {
	return bce.URI_PREFIX
}
func getModifyJobUri() string {
	return bce.URI_PREFIX
}
func getModifyModelUri() string {
	return bce.URI_PREFIX
}
func getStopJobUri() string {
	return bce.URI_PREFIX
}
