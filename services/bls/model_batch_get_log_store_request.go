package bls

type BatchGetLogStoreRequest struct {
	LogStores []*LogStoreBatchRequest `json:"LogStores,omitempty"`
}
