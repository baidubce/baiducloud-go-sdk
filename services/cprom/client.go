package cprom

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "cprom." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_V2 = "v2"

	CONSTANT_NOTIFY_RULE = "notify_rule"

	CONSTANT_BCM_SCOPES = "bcm_scopes"

	CONSTANT_INSTANCE = "instance"

	CONSTANT_SERVICE_MONITOR = "service_monitor"

	CONSTANT_ALERTING_RULE = "alerting_rule"

	CONSTANT_POD_MONITOR = "pod_monitor"

	CONSTANT_SCRAPE_JOB = "scrape_job"

	CONSTANT_BCM_JOB = "bcm_job"

	CONSTANT_SCOPES = "scopes"

	CONSTANT_POD_MONITOR_SERVICE = "pod_monitor_service"

	CONSTANT_EVENT = "event"

	CONSTANT_CLAIM = "claim"

	CONSTANT_PROMETHEUS = "prometheus"

	CONSTANT_API = "api"

	CONSTANT_V1 = "v1"

	CONSTANT_QUERY_RANGE = "query_range"

	CONSTANT_TOKEN = "token"

	CONSTANT_CLUSTER_BINDING = "cluster_binding"

	CONSTANT_WRITE = "write"

	CONSTANT_ALERTING_RULE_TEMPLATE = "alerting_rule_template"
)

// Client of cprom service is a kind of BceClient, so derived from BceClient
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

func getBindClusterUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getClaimAlertEventUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_EVENT + bce.URI_PREFIX + CONSTANT_CLAIM
}
func getCreateAlertUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE
}
func getCreateCustomScrapeTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SCRAPE_JOB
}
func getCreateInstanceUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getCreateNotificationPolicyUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NOTIFY_RULE
}
func getCreatePodmonitorUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_POD_MONITOR
}
func getCreateServiceMonitorUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SERVICE_MONITOR
}
func getDeleteAlertUri(AlertingRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE + bce.URI_PREFIX + AlertingRuleId
}
func getDeleteCustomScrapeTaskUri(ScrapeJobId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SCRAPE_JOB + bce.URI_PREFIX + ScrapeJobId
}
func getDeleteInstanceUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId
}
func getDeleteNotificationPolicyUri(NotifyRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NOTIFY_RULE + bce.URI_PREFIX + NotifyRuleId
}
func getDeletePodmonitorUri(PodMonitorName string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_POD_MONITOR + bce.URI_PREFIX + PodMonitorName
}
func getDeleteServiceMonitorUri(ServiceMonitorName string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SERVICE_MONITOR + bce.URI_PREFIX + ServiceMonitorName
}
func getGenerateInstanceTokenUri(InstanceId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE + bce.URI_PREFIX + InstanceId + bce.URI_PREFIX + CONSTANT_TOKEN
}
func getGetAlertDetailUri(AlertingRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE + bce.URI_PREFIX + AlertingRuleId
}
func getGetAlertEventDetailUri(EventId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_EVENT + bce.URI_PREFIX + EventId
}
func getGetClusterBindStatusUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_CLUSTER_BINDING
}
func getGetNotificationPolicyUri(NotifyRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NOTIFY_RULE + bce.URI_PREFIX + NotifyRuleId
}
func getListAlertEventsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_EVENT
}
func getListAlertTemplatesUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE_TEMPLATE
}
func getListAlertsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE
}
func getListBindableCloudProductsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_BCM_SCOPES
}
func getListInstancesUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_INSTANCE
}
func getListNotificationPoliciesUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NOTIFY_RULE
}
func getListPodMonitorsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_POD_MONITOR
}
func getListRelatedCloudProductsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_BCM_JOB + bce.URI_PREFIX + CONSTANT_SCOPES
}
func getListServiceMonitorsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SERVICE_MONITOR
}
func getRemoteReadUri(RemoteReadUrl string) string {
	return bce.URI_PREFIX + CONSTANT_PROMETHEUS + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_QUERY_RANGE
}
func getRemoteWriteUri(RemoteWriteUrl string) string {
	return bce.URI_PREFIX + CONSTANT_PROMETHEUS + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_WRITE
}
func getTogglePodMonitorServiceUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_POD_MONITOR_SERVICE
}
func getToggleServiceMonitorServiceUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SERVICE_MONITOR
}
func getUpdateAlertUri(AlertingRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_ALERTING_RULE + bce.URI_PREFIX + AlertingRuleId
}
func getUpdateNotificationPolicyUri(NotifyRuleId string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_NOTIFY_RULE + bce.URI_PREFIX + NotifyRuleId
}
func getUpdatePodMonitorUri(PodMonitorName string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_POD_MONITOR + bce.URI_PREFIX + PodMonitorName
}
func getUpdateRelatedCloudProductsUri() string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_BCM_JOB + bce.URI_PREFIX + CONSTANT_SCOPES
}
func getUpdateServiceMonitorUri(ServiceMonitorName string) string {
	return bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SERVICE_MONITOR + bce.URI_PREFIX + ServiceMonitorName
}
