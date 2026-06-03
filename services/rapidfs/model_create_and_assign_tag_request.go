package rapidfs

type CreateAndAssignTagRequest struct {
	Action       *string        `json:"-"`
	ClientToken  *string        `json:"-"`
	TagResources []*TagResource `json:"tagResources,omitempty"`
}
