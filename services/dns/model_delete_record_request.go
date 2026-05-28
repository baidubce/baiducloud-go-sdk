package dns

type DeleteRecordRequest struct {
	ZoneName    *string `json:"-"`
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
