package rapidfs

type CreateAndAssignTagRequest struct {
	ClientToken  *string        `json:"-"`
	TagResources []*TagResource `json:"tagResources,omitempty"`
}
