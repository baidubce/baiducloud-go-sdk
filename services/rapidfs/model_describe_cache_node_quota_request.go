package rapidfs

type DescribeCacheNodeQuotaRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
}
