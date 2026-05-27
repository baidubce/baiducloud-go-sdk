package privatezone

type EnableRecordRequest struct {
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
	Action      *string `json:"-"`
}
