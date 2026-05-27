package csn

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "csn." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_CSN = "csn"

	CONSTANT_ROUTE_TABLE = "routeTable"

	CONSTANT_RULE = "rule"

	CONSTANT_BP = "bp"

	CONSTANT_LIMIT = "limit"

	CONSTANT_PROPAGATION = "propagation"

	CONSTANT_TGW = "tgw"

	CONSTANT_ASSOCIATION = "association"

	CONSTANT_PRICE = "price"

	CONSTANT_DELETE = "delete"

	CONSTANT_INSTANCE = "instance"
)

// Client of csn service is a kind of BceClient, so derived from BceClient
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

func getAddRouteRuleUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_RULE
}
func getAttachCsnInstanceUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId
}
func getBindCsnBpUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getCreateAssociationRelationUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_ASSOCIATION
}
func getCreateCsnUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN
}
func getCreateCsnBpUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP
}
func getCreateRegionBandwidthUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId + bce.URI_PREFIX + CONSTANT_LIMIT
}
func getCreateStudyRelationUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_PROPAGATION
}
func getDeleteAssociationRelationUri(version string, CsnRtId string, AttachId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_ASSOCIATION + bce.URI_PREFIX + AttachId
}
func getDeleteCsnUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId
}
func getDeleteCsnBpUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getDeleteRegionBandwidthUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId + bce.URI_PREFIX + CONSTANT_LIMIT + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRouteRuleUri(version string, CsnRtId string, CsnRtRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + CsnRtRuleId
}
func getDeleteStudyRelationUri(version string, CsnRtId string, AttachId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_PROPAGATION + bce.URI_PREFIX + AttachId
}
func getDetachCsnInstanceUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId
}
func getQueryAssociationRelationUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_ASSOCIATION
}
func getQueryCsnBpDetailUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getQueryCsnBpListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP
}
func getQueryCsnBpPriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CONSTANT_PRICE
}
func getQueryCsnDetailUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId
}
func getQueryCsnInstanceUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getQueryCsnListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN
}
func getQueryRegionBandwidthUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId + bce.URI_PREFIX + CONSTANT_LIMIT
}
func getQueryRegionBandwidthByCsnUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CONSTANT_LIMIT
}
func getQueryRouteRuleUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_RULE
}
func getQueryRouteTableListUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE
}
func getQueryStudyRelationUri(version string, CsnRtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_ROUTE_TABLE + bce.URI_PREFIX + CsnRtId + bce.URI_PREFIX + CONSTANT_PROPAGATION
}
func getQueryTgwListUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_TGW
}
func getQueryTgwRouteUri(version string, CsnId string, TgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_TGW + bce.URI_PREFIX + TgwId + bce.URI_PREFIX + CONSTANT_RULE
}
func getResizeCsnBpUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getUnbindCsnBpUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getUpdateCsnUri(version string, CsnId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId
}
func getUpdateCsnBpUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId
}
func getUpdateRegionBandwidthUri(version string, CsnBpId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CONSTANT_BP + bce.URI_PREFIX + CsnBpId + bce.URI_PREFIX + CONSTANT_LIMIT
}
func getUpdateTgwUri(version string, CsnId string, TgwId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CSN + bce.URI_PREFIX + CsnId + bce.URI_PREFIX + CONSTANT_TGW + bce.URI_PREFIX + TgwId
}
