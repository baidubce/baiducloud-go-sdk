package cprom

type DeleteAlertRequest struct {
	AlertingRuleId *string `json:"-"`
	InstanceId     *string `json:"-"`
}
