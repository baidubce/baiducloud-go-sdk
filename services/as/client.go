package as

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "as." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V2 = "v2"

	CONSTANT_GROUP = "group"

	CONSTANT_V1 = "v1"

	CONSTANT_RULE = "rule"

	CONSTANT_NODE = "node"

	CONSTANT_DELETE = "delete"

	CONSTANT_TASK = "task"
)

// Client of as service is a kind of BceClient, so derived from BceClient
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

func getAdjustNumV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getAttachNodeV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getCreateAsGroupV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP
}
func getCreateRuleV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_RULE
}
func getDeleteAsGroupV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRuleV2Uri(RuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId
}
func getDetachNodeV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getExecRuleV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getGetAsGroupV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getGetRuleV2Uri(RuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId
}
func getListAsGroupV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_GROUP
}
func getListAsNodeV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_NODE
}
func getListRuleV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_RULE
}
func getListTaskV2Uri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TASK
}
func getScalingDownV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getScalingUpV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
func getUpdateIsManagedV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_NODE + bce.URI_PREFIX + GroupId
}
func getUpdateProtectV2Uri(GroupId string) string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + GroupId
}
