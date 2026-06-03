package rapidfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "rapidfs." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_CACHENODE = "cachenode"

	CONSTANT_ORDER = "order"

	CONSTANT_AUTHGROUP = "authgroup"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_METASYNCRULE = "metasyncrule"

	CONSTANT_DATASRC = "datasrc"

	CONSTANT_CACHEDEPLOYGROUP = "cachedeploygroup"

	CONSTANT_CACHERULE = "cacherule"

	CONSTANT_AIHCRESOURCEPOOL = "aihcresourcepool"

	CONSTANT_TAG = "tag"

	CONSTANT_CCECLUSTER = "ccecluster"

	CONSTANT_ZONE = "zone"

	CONSTANT_PRICE = "price"

	CONSTANT_SPECIFICATION = "specification"
)

// Client of rapidfs service is a kind of BceClient, so derived from BceClient
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

func getAddCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getCancelCacheRuleJobUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getCancelMetaSyncJobUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getCheckBeforeAddCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getCheckBeforeCreateInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateAndAssignTagUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_TAG
}
func getCreateAuthGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AUTHGROUP
}
func getCreateCacheRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getCreateInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getDeleteAuthGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AUTHGROUP
}
func getDeleteCacheRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getDeleteInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDeleteMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getDescribeAihcResourcePoolsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AIHCRESOURCEPOOL
}
func getDescribeAllocatableDataSrcQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getDescribeAuthGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AUTHGROUP
}
func getDescribeAuthGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AUTHGROUP
}
func getDescribeCacheDeployGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHEDEPLOYGROUP
}
func getDescribeCacheDeployGroupsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHEDEPLOYGROUP
}
func getDescribeCacheNodeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getDescribeCacheNodeBccCandidatesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getDescribeCacheNodeQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getDescribeCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getDescribeCacheRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getDescribeCacheRuleJobsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getDescribeCacheRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getDescribeCceClustersUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CCECLUSTER
}
func getDescribeDataSrcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getDescribeDataSrcsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getDescribeInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDescribeInstancesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDescribeMetaSyncJobsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getDescribeMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getDescribeMetaSyncRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getDescribeOrderUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ORDER
}
func getDescribePriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PRICE
}
func getDescribeSpecsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SPECIFICATION
}
func getDescribeTokenUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDescribeZonesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ZONE
}
func getDisableMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getEnableMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getExecuteCacheRuleJobUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHERULE
}
func getExecuteMetaSyncJobUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getImportDataSrcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getModifyAuthGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_AUTHGROUP
}
func getModifyDataSrcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getModifyMetaSyncRuleUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_METASYNCRULE
}
func getModifyTokenUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getRemoveCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getRemoveDataSrcUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DATASRC
}
func getResizeInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getRestartCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getStartCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
func getStopCacheNodesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CACHENODE
}
