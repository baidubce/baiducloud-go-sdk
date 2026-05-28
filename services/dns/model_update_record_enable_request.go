package dns

type UpdateRecordEnableRequest struct {
	ZoneName    *string `json:"-"`
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
