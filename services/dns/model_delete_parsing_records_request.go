package dns

type DeleteParsingRecordsRequest struct {
	ZoneName    *string `json:"-"`
	RecordId    *string `json:"-"`
	ClientToken *string `json:"-"`
}
