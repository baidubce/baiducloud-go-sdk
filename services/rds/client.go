package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "rds." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_TAGS = "tags"

	CONSTANT_STATUS = "status"

	CONSTANT_DISK_AUTO_RESIZE = "diskAutoResize"

	CONSTANT_CONFIG = "config"

	CONSTANT_PROBE = "probe"

	CONSTANT_V1 = "v1"

	CONSTANT_TASK = "task"

	CONSTANT_RENEW = "renew"

	CONSTANT_AUTO_RESIZE_CONFIG = "autoResizeConfig"

	CONSTANT_ZONE = "zone"

	CONSTANT_SUBNET = "subnet"

	CONSTANT_MAINTAINTIME = "maintaintime"

	CONSTANT_AUTO_EXPANSION = "autoExpansion"
)

// Client of rds service is a kind of BceClient, so derived from BceClient
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

func getBatchConfigurationOfInstancesUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getBindingTagsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TAGS
}
func getClusterStatusCheckUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_STATUS
}
func getCreateAPrimaryInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateAProxyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateReadOnlyInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getDialTestInterfaceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_PROBE
}
func getDoesTheInstanceSupportEnablingAutomaticExpansionUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_EXPANSION
}
func getInstanceEnablesAutomaticRenewalUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getQueryInstanceListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getQuerySubnetListUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_SUBNET
}
func getQueryTaskListUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_TASK
}
func getQueryTheDetailsOfTheSpecifiedInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getQueryTheListOfAvailableZonesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ZONE
}
func getReleaseInstanceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getRenewalExampleUri() string {
	return bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + CONSTANT_RENEW
}
func getRestartTheInstanceUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getRetrieveInstanceAutoScalingConfigUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_AUTO_RESIZE_CONFIG
}
func getTransformationExampleUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateAutomaticStorageExpansionConfigurationUri(version string, InstanceId string, Action string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_DISK_AUTO_RESIZE + bce.URI_PREFIX + CONSTANT_CONFIG + bce.URI_PREFIX + Action
}
func getUpdateInstanceConnectionInformationUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstanceNameUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstancePublicNetworkAccessStatusUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateInstanceSynchronizationModeUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getUpdateTimeWindowUri(version string, InstanceId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_MAINTAINTIME
}
