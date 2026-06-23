package cprom

type RuleDetail struct {
	AlertId      *string            `json:"alertId,omitempty"`
	AlertName    *string            `json:"alertName,omitempty"`
	Expr         *string            `json:"expr,omitempty"`
	CpromFor     *string            `json:"for,omitempty"`
	Description  *string            `json:"description,omitempty"`
	Enable       *bool              `json:"enable,omitempty"`
	NotifyRuleId *string            `json:"notifyRuleId,omitempty"`
	Severity     *string            `json:"severity,omitempty"`
	Labels       *map[string]string `json:"labels,omitempty"`
	Annotations  *map[string]string `json:"annotations,omitempty"`
}
