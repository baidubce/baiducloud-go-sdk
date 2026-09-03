package aigw

type DeleteConsumerRequest struct {
	InstanceId *string `json:"-"`
	ConsumerId *string `json:"-"`
	KeyType    *string `json:"-"`
	XRegion    *string `json:"-"`
}
