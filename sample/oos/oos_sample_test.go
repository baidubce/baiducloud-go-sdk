package oossample

import (
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/bce"
)

func TestOOSV2SamplesReturnAfterRequestError(t *testing.T) {
	bce.DEFAULT_RETRY_POLICY = bce.NewBackOffRetryPolicy(0, 0, 0)

	tests := []struct {
		name string
		call func()
	}{
		{"check template", CheckTemplateV2},
		{"create execution", CreateExecutionV2},
		{"create template", CreateTemplateV2},
		{"delete template", DeleteTemplateV2},
		{"get execution detail", GetExecutionDetailV2},
		{"get execution list", GetExecutionListV2},
		{"get operator list", GetOperatorListV2},
		{"get task children list", GetTaskChildrenListV2},
		{"get task detail", GetTaskDetailV2},
		{"get template detail", GetTemplateDetailV2},
		{"get template list", GetTemplateListV2},
		{"update template", UpdateTemplateV2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.call()
		})
	}
}
