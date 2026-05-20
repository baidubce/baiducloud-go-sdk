package dns

type ModifyTheParsingRecordStatusRequest struct {
	ZoneName    *string `json:"-"`
	RecordId    *string `json:"-"`
	Action      *string `json:"-"`
	ClientToken *string `json:"-"`
}
