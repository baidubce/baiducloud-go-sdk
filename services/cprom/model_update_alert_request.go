package cprom

type UpdateAlertRequest struct {
	AlertingRuleId *string            `json:"-"`
	InstanceId     *string            `json:"-"`
	AlertName      *string            `json:"alertName,omitempty"`
	Expr           *string            `json:"expr,omitempty"`
	CpromFor       *string            `json:"for,omitempty"`
	Description    *string            `json:"description,omitempty"`
	NotifyRuleId   *string            `json:"notifyRuleId,omitempty"`
	Severity       *string            `json:"severity,omitempty"`
	Enable         *bool              `json:"enable,omitempty"`
	Labels         *map[string]string `json:"labels,omitempty"`
	Annotations    *map[string]string `json:"annotations,omitempty"`
}
