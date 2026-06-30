package iam

type ListStrategiesRequest struct {
	PolicyType *string `json:"-"`
	NameFilter *string `json:"-"`
}
