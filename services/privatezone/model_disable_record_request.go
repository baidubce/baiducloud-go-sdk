package privatezone

type DisableRecordRequest struct {
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
