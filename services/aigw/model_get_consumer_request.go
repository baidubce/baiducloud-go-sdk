package aigw

type GetConsumerRequest struct {
	InstanceId *string `json:"-"`
	ConsumerId *string `json:"-"`
	KeyType    *string `json:"-"`
	XRegion    *string `json:"-"`
}
