package csn

type CreateAssociationRelationRequest struct {
	CsnRtId     *string `json:"-"`
	ClientToken *string `json:"-"`
	AttachId    *string `json:"attachId,omitempty"`
	Description *string `json:"description,omitempty"`
}
