package bls

type DeleteLogStoreRequest struct {
	LogStoreName *string `json:"-"`
	Project      *string `json:"-"`
}
