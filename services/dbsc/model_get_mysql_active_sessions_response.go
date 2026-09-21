package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlActiveSessionsResponse struct {
	bce.BaseResponse
	Items              []*MySQLSession            `json:"items,omitempty"`
	DatabaseStatistics []*MysqlSessionDBSummary   `json:"databaseStatistics,omitempty"`
	HostStatistics     []*MysqlSessionHostSummary `json:"hostStatistics,omitempty"`
	UserStatistics     []*MysqlSessionUserSummary `json:"userStatistics,omitempty"`
	Summary            []*MySQLSessionSummary     `json:"summary,omitempty"`
}
