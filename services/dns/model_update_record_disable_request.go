package dns

type UpdateRecordDisableRequest struct {
	ZoneName    *string `json:"-"`
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
