package privatezone

type SetParsingRecordStatusRequest struct {
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
	Action      *string `json:"-"`
}
