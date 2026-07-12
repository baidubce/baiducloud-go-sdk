package oos

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "oos." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_TEMPLATE = "template"

	CONSTANT_EXECUTION = "execution"

	CONSTANT_CHECK = "check"

	CONSTANT_LIST = "list"

	CONSTANT_TASK = "task"

	CONSTANT_OPERATOR = "operator"

	CONSTANT_CHILDREN = "children"
)

// Client of oos service is a kind of BceClient, so derived from BceClient
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

func getCheckTemplateV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_CHECK
}
func getCreateExecutionV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EXECUTION
}
func getCreateTemplateV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getDeleteTemplateV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getGetExecutionDetailV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EXECUTION
}
func getGetExecutionListV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_EXECUTION + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetOperatorListV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_OPERATOR + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetTaskChildrenListV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_CHILDREN + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetTaskDetailV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TASK
}
func getGetTemplateDetailV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getGetTemplateListV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE + bce.URI_PREFIX + CONSTANT_LIST
}
func getUpdateTemplateV2Uri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
