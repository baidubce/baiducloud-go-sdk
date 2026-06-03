package ccr

type TriggerTagScanRequest struct {
	InstanceId     *string `json:"-"`
	ProjectName    *string `json:"-"`
	RepositoryName *string `json:"-"`
	TagName        *string `json:"-"`
}
