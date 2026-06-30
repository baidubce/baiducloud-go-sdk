package bci

type DeleteInstanceRequest struct {
	InstanceId         *string `json:"-"`
	RelatedReleaseFlag *bool   `json:"-"`
}
