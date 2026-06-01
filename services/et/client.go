package et

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "et." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_ET = "et"

	CONSTANT_CHANNEL = "channel"

	CONSTANT_ROUTE = "route"

	CONSTANT_RULE = "rule"

	CONSTANT_INIT = "init"

	CONSTANT_BFD = "bfd"
)

// Client of et service is a kind of BceClient, so derived from BceClient
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

func getApplyPhysicalDedicatedLineUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + CONSTANT_INIT
}
func getAssociatedDedicatedChannelUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getCreateDedicatedChannelUri(version string, EtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL
}
func getCreateDedicatedChannelBfdUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_BFD
}
func getCreateDedicatedChannelRouteParametersUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getCreateDedicatedChannelRouteRulesUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getCreateDedicatedChannelUserObjectUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getDeleteDedicatedChannelUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getDeleteDedicatedChannelBfdUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_BFD
}
func getDeleteDedicatedChannelRouteRulesUri(version string, EtId string, EtChannelId string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getDisableDedicatedChannelIpv6Uri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getEnableDedicatedChannelIpv6Uri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getQueryDedicatedChannelUri(version string, EtId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL
}
func getQueryDedicatedChannelRouteRulesUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE
}
func getQueryDedicatedLineDetailUri(version string, DcphyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + DcphyId
}
func getQueryDedicatedLinesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET
}
func getRemoveDedicatedChannelRouteParametersUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getRemoveDedicatedChannelUserObjectUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getResubmitDedicatedChannelUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getUnrelatedDedicatedLineChannelUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getUpdateDedicatedChannelUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId
}
func getUpdateDedicatedChannelBfdUri(version string, EtId string, EtChannelId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_BFD
}
func getUpdateDedicatedChannelRouteRulesUri(version string, EtId string, EtChannelId string, RouteRuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + EtId + bce.URI_PREFIX + CONSTANT_CHANNEL + bce.URI_PREFIX + EtChannelId + bce.URI_PREFIX + CONSTANT_ROUTE + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RouteRuleId
}
func getUpdatePhysicalDedicatedLineUri(version string, DcphyId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ET + bce.URI_PREFIX + DcphyId
}
