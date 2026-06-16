package cfw

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "cfw." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_CFW = "cfw"

	CONSTANT_STATELESS = "stateless"

	CONSTANT_RULE = "rule"

	CONSTANT_DELETE = "delete"

	CONSTANT_INSTANCE = "instance"
)

// Client of cfw service is a kind of BceClient, so derived from BceClient
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

func getBindCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getCreateCfwUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW
}
func getCreateCfwRuleUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId + bce.URI_PREFIX + CONSTANT_RULE
}
func getCreateStatelessCfwUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CONSTANT_STATELESS
}
func getDeleteCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getDeleteCfwRuleUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId + bce.URI_PREFIX + CONSTANT_DELETE + bce.URI_PREFIX + CONSTANT_RULE
}
func getDisableCfwProtectUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getEnableCfwProtectUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getGetCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getGetStatelessCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CONSTANT_STATELESS + bce.URI_PREFIX + CfwId
}
func getListCfwUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW
}
func getListProtectInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListStatelessCfwUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CONSTANT_STATELESS
}
func getUnbindCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getUpdateCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId
}
func getUpdateCfwRuleUri(version string, CfwId string, CfwRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CfwId + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CfwRuleId
}
func getUpdateStatelessCfwUri(version string, CfwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFW + bce.URI_PREFIX + CONSTANT_STATELESS + bce.URI_PREFIX + CfwId
}
