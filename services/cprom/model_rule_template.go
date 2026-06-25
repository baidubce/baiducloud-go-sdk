package cprom

type RuleTemplate struct {
	AlertName   *string `json:"alertName,omitempty"`
	Expr        *string `json:"expr,omitempty"`
	CpromFor    *string `json:"for,omitempty"`
	Description *string `json:"description,omitempty"`
}
