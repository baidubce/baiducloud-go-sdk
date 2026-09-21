package dbsc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "dbsc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_DIAGNOSIS = "diagnosis"

	CONSTANT_MYSQL = "mysql"

	CONSTANT_SQLFILTER = "sqlfilter"

	CONSTANT_DELETE = "delete"

	CONSTANT_SESSION = "session"

	CONSTANT_KILL = "kill"

	CONSTANT_LIST = "list"

	CONSTANT_V1 = "v1"

	CONSTANT_ACTION = "action"

	CONSTANT_API = "api"

	CONSTANT_REDIS = "redis"

	CONSTANT_SLOWLOG = "slowlog"

	CONSTANT_TREND = "trend"

	CONSTANT_STATS = "stats"

	CONSTANT_DURATION = "duration"

	CONSTANT_DEADLOCK = "deadlock"

	CONSTANT_LATEST = "latest"

	CONSTANT_MONGODB = "mongodb"

	CONSTANT_SPACE = "space"

	CONSTANT_COLLECTION = "collection"

	CONSTANT_PEGA = "pega"

	CONSTANT_SCHEMA = "schema"

	CONSTANT_TABLE = "table"

	CONSTANT_INDEX = "index"

	CONSTANT_BIG_KEY = "big-key"

	CONSTANT_TASK = "task"

	CONSTANT_TEMPLATE = "template"

	CONSTANT_DATABASE = "database"

	CONSTANT_POSTGRESQL = "postgresql"

	CONSTANT_HISTORY = "history"

	CONSTANT_SUMMARY = "summary"

	CONSTANT_RESULT = "result"

	CONSTANT_ALLOWED = "allowed"
)

// Client of dbsc service is a kind of BceClient, so derived from BceClient
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

func getCheckMysqlRateLimitSupportUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER + bce.URI_PREFIX + CONSTANT_ALLOWED
}
func getCreateMysqlRateLimitTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER
}
func getCreateRedisBigKeyAnalysisTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_BIG_KEY + bce.URI_PREFIX + CONSTANT_TASK
}
func getDeleteMysqlRateLimitTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER + bce.URI_PREFIX + CONSTANT_DELETE
}
func getDeleteRedisBigKeyAnalysisTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_BIG_KEY + bce.URI_PREFIX + CONSTANT_TASK
}
func getGetMongodbCollectionIndexesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SCHEMA + bce.URI_PREFIX + CONSTANT_COLLECTION + bce.URI_PREFIX + CONSTANT_INDEX
}
func getGetMongodbCollectionSpaceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_COLLECTION
}
func getGetMongodbCollectionSpaceTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_COLLECTION + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetMongodbDatabaseSpaceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_DATABASE
}
func getGetMongodbDatabaseSpaceTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_DATABASE + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetMongodbSlowLogTimeDistributionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_STATS + bce.URI_PREFIX + CONSTANT_DURATION
}
func getGetMongodbSlowLogTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetMongodbSlowQueryTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getGetMongodbSpaceSummaryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_SUMMARY
}
func getGetMysqlActiveSessionsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SESSION + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetMysqlDatabaseSpaceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_DATABASE
}
func getGetMysqlDeadlockInfoUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_DEADLOCK + bce.URI_PREFIX + CONSTANT_LATEST
}
func getGetMysqlKillSessionHistoryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SESSION + bce.URI_PREFIX + CONSTANT_KILL + bce.URI_PREFIX + CONSTANT_HISTORY
}
func getGetMysqlRateLimitTaskDetailUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER
}
func getGetMysqlSlowLogTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getGetMysqlSlowLogTimeDistributionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_STATS + bce.URI_PREFIX + CONSTANT_DURATION
}
func getGetMysqlSlowLogTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetMysqlSpaceSummaryUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_SUMMARY
}
func getGetMysqlTableIndexesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SCHEMA + bce.URI_PREFIX + CONSTANT_TABLE + bce.URI_PREFIX + CONSTANT_INDEX
}
func getGetMysqlTableSpaceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SPACE + bce.URI_PREFIX + CONSTANT_TABLE
}
func getGetPegadbSlowLogTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_PEGA + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_SUMMARY
}
func getGetPegadbSlowLogTimeDistributionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_PEGA + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_STATS + bce.URI_PREFIX + CONSTANT_DURATION
}
func getGetPegadbSlowLogTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_PEGA + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetPostgresqlSlowLogTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_POSTGRESQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TEMPLATE
}
func getGetPostgresqlSlowLogTimeDistributionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_POSTGRESQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_STATS + bce.URI_PREFIX + CONSTANT_DURATION
}
func getGetPostgresqlSlowLogTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_POSTGRESQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TREND
}
func getGetRedisBigKeyAnalysisResultUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_BIG_KEY + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_LIST
}
func getGetRedisSlowLogTemplateUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_SUMMARY
}
func getGetRedisSlowLogTimeDistributionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_STATS + bce.URI_PREFIX + CONSTANT_DURATION
}
func getGetRedisSlowLogTrendUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_TREND
}
func getKillMysqlSessionUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SESSION + bce.URI_PREFIX + CONSTANT_KILL
}
func getListMongodbSlowLogsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MONGODB + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getListMysqlRateLimitTasksUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER + bce.URI_PREFIX + CONSTANT_LIST
}
func getListMysqlSlowLogsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getListPegadbSlowLogsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_PEGA + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getListPostgresqlSlowLogsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_POSTGRESQL + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getListRedisBigKeyAnalysisTasksUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_BIG_KEY + bce.URI_PREFIX + CONSTANT_TASK + bce.URI_PREFIX + CONSTANT_RESULT
}
func getListRedisSlowLogsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_API + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_REDIS + bce.URI_PREFIX + CONSTANT_SLOWLOG + bce.URI_PREFIX + CONSTANT_LIST
}
func getStartStopMysqlInstanceFlowLimitingTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_DIAGNOSIS + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER + bce.URI_PREFIX + CONSTANT_ACTION
}
func getUpdateMysqlRateLimitTaskUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_MYSQL + bce.URI_PREFIX + CONSTANT_SQLFILTER
}
