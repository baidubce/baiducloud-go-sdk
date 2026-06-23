package cprom

type GetAlertDetailRequest struct {
	AlertingRuleId *string `json:"-"`
	InstanceId     *string `json:"-"`
}
