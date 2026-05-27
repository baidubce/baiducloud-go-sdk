package csn

type DeleteStudyRelationRequest struct {
	CsnRtId     *string `json:"-"`
	AttachId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
