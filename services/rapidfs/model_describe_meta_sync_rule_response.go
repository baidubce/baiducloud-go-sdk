package rapidfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeMetaSyncRuleResponse struct {
	bce.BaseResponse
	MetaSyncRuleInfo *MetaSyncRuleInfo `json:"metaSyncRuleInfo,omitempty"`
}
