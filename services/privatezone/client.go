package privatezone

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "privatezone." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_PRIVATEZONE = "privatezone"

	CONSTANT_RESOLVER = "resolver"

	CONSTANT_RECORD = "record"

	CONSTANT_RULE = "rule"

	CONSTANT_UNBIND = "unbind"

	CONSTANT_BIND = "bind"
)

// Client of privatezone service is a kind of BceClient, so derived from BceClient
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

func getAddRecordUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getBindVpcUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getBindVpcToRuleUri(version string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId + bce.URI_PREFIX + CONSTANT_BIND
}
func getCreatePrivateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getCreateResolverUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER
}
func getCreateResolverRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE
}
func getDeletePrivateZoneUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getDeleteRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getDeleteResolverUri(version string, ResolverId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + ResolverId
}
func getDeleteResolverRuleUri(version string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId
}
func getDisableRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getEnableRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getGetDnsResolverDetailUri(version string, ResolverId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + ResolverId
}
func getGetDnsResolverListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER
}
func getGetDnsResolverRuleDetailUri(version string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId
}
func getGetDnsResolverRuleListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE
}
func getGetPrivateZoneUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getListPrivateZoneUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE
}
func getListRecordUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId + bce.URI_PREFIX + CONSTANT_RECORD
}
func getUnbindVpcUri(version string, ZoneId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + ZoneId
}
func getUnbindVpcToRuleUri(version string, RuleId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + RuleId + bce.URI_PREFIX + CONSTANT_UNBIND
}
func getUpdateDnsParserUri(version string, ResolverId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + ResolverId
}
func getUpdateRecordUri(version string, RecordId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + RecordId
}
func getUpdateResolverRuleUri(version string, ReluId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRIVATEZONE + bce.URI_PREFIX + CONSTANT_RESOLVER + bce.URI_PREFIX + CONSTANT_RULE + bce.URI_PREFIX + ReluId
}
