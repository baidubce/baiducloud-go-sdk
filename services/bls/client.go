package bls

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bls." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_ALARM = "alarm"

	CONSTANT_RECORD = "record"

	CONSTANT_LOGSTORE = "logstore"

	CONSTANT_DOWNLOAD = "download"

	CONSTANT_LIST = "list"

	CONSTANT_BLS = "bls"

	CONSTANT_LOGSHIPPER = "logshipper"

	CONSTANT_PROJECT = "project"

	CONSTANT_INDEX = "index"

	CONSTANT_EXECUTION = "execution"

	CONSTANT_LOGSTREAM = "logstream"

	CONSTANT_STATUS = "status"

	CONSTANT_BATCH = "batch"

	CONSTANT_POLICY = "policy"

	CONSTANT_FIELD_CAPS = "_field_caps"

	CONSTANT_FASTQUERY = "fastquery"

	CONSTANT_TASK = "task"

	CONSTANT_LOGHISTOGRAM = "loghistogram"

	CONSTANT_LOGRECORD = "logrecord"

	CONSTANT_ENABLE = "enable"

	CONSTANT_LINK = "link"

	CONSTANT_CONDITION = "condition"

	CONSTANT_VALIDATE = "validate"

	CONSTANT_ASYNC_SEARCH = "_async_search"

	CONSTANT_DISABLE = "disable"

	CONSTANT_RESOLVE = "_resolve"

	CONSTANT_TERMS_ENUM = "_terms_enum"
)

// Client of bls service is a kind of BceClient, so derived from BceClient
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

func getAsyncSearchUri(Name string) string {
	return bce.URI_PREFIX + Name + bce.URI_PREFIX + CONSTANT_ASYNC_SEARCH
}
func getBatchGetLogStoreUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_BATCH
}
func getBulkDeleteLogShipperUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER
}
func getBulkSetLogShipperStatusUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + CONSTANT_STATUS + bce.URI_PREFIX + CONSTANT_BATCH
}
func getCreateAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY
}
func getCreateDownloadTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_DOWNLOAD
}
func getCreateFastQueryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_FASTQUERY
}
func getCreateIndexUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_INDEX
}
func getCreateLogShipperUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER
}
func getCreateLogStoreUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE
}
func getCreateLogStoreTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getCreateLogStoreViewUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getCreateProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT
}
func getCreateTaskUri() string {
	return bce.URI_PREFIX + CONSTANT_TASK
}
func getDeleteAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY
}
func getDeleteDownloadTaskUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_DOWNLOAD + bce.URI_PREFIX + Uuid
}
func getDeleteFastQueryUri(version string, FastQueryName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_FASTQUERY + bce.URI_PREFIX + FastQueryName
}
func getDeleteIndexUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_INDEX
}
func getDeleteLogStoreUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName
}
func getDeleteLogStoreTemplatesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getDeleteLogStoreViewUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getDeleteProjectUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + Uuid
}
func getDeleteSingleLogShipperUri(version string, LogShipperID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + LogShipperID
}
func getDescribeAlarmPolicyUri(version string, PolicyName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY
}
func getDescribeAlarmRecordUri(version string, AlarmId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_RECORD
}
func getDescribeDownloadTaskUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_DOWNLOAD + bce.URI_PREFIX + Uuid
}
func getDescribeFastQueryUri(version string, FastQueryName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_FASTQUERY + bce.URI_PREFIX + FastQueryName
}
func getDescribeIndexUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_INDEX
}
func getDescribeLogStoreUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName
}
func getDescribeLogStoreTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getDescribeLogStoreTemplatesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getDescribeLogStoreViewUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getDescribeProjectUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + Uuid
}
func getDisableAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + CONSTANT_DISABLE
}
func getEnableAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + CONSTANT_ENABLE
}
func getFieldCapsUri(Name string) string {
	return bce.URI_PREFIX + Name + bce.URI_PREFIX + CONSTANT_FIELD_CAPS
}
func getGetDownloadTaskLinkUri(version string, Uuid string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_DOWNLOAD + bce.URI_PREFIX + CONSTANT_LINK + bce.URI_PREFIX + Uuid
}
func getGetLogShipperUri(version string, LogShipperID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + LogShipperID
}
func getListAlarmExecutionStatsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_EXECUTION + bce.URI_PREFIX + CONSTANT_LIST
}
func getListAlarmExecutionsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_EXECUTION + bce.URI_PREFIX + CONSTANT_LIST
}
func getListAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY + bce.URI_PREFIX + CONSTANT_LIST
}
func getListAlarmRecordUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_RECORD + bce.URI_PREFIX + CONSTANT_LIST
}
func getListDownloadTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_DOWNLOAD + bce.URI_PREFIX + CONSTANT_LIST
}
func getListFastQueryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_FASTQUERY
}
func getListLogShipperUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER
}
func getListLogShipperRecordUri(version string, LogShipperID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + LogShipperID + bce.URI_PREFIX + CONSTANT_RECORD
}
func getListLogStoreUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_LIST
}
func getListLogStoreViewUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getListLogStreamUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGSTREAM
}
func getListProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT + bce.URI_PREFIX + CONSTANT_LIST
}
func getPullLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getPushLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getQueryLogHistogramUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGHISTOGRAM
}
func getQueryLogRecordUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_LOGRECORD
}
func getResolveIndexUri(Name string) string {
	return bce.URI_PREFIX + CONSTANT_RESOLVE + bce.URI_PREFIX + CONSTANT_INDEX + bce.URI_PREFIX + Name
}
func getSetSingleLogShipperStatusUri(version string, LogShipperID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + LogShipperID + bce.URI_PREFIX + CONSTANT_STATUS
}
func getTermsEnumUri(Name string) string {
	return bce.URI_PREFIX + Name + bce.URI_PREFIX + CONSTANT_TERMS_ENUM
}
func getUpdateAlarmPolicyUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_POLICY
}
func getUpdateFastQueryUri(version string, Name string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_FASTQUERY + bce.URI_PREFIX + Name
}
func getUpdateIndexUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName + bce.URI_PREFIX + CONSTANT_INDEX
}
func getUpdateLogShipperUri(version string, LogShipperID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSHIPPER + bce.URI_PREFIX + LogShipperID
}
func getUpdateLogStoreUri(version string, LogStoreName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + LogStoreName
}
func getUpdateLogStoreTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getUpdateLogStoreViewUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_BLS
}
func getUpdateProjectUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_PROJECT
}
func getUpdateTaskUri(TaskId string) string {
	return bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + TaskId
}
func getValidateAlarmConditionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ALARM + bce.URI_PREFIX + CONSTANT_CONDITION + bce.URI_PREFIX + CONSTANT_VALIDATE
}
func getValidateAlarmPolicySQLUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_LOGSTORE + bce.URI_PREFIX + CONSTANT_VALIDATE
}
