package privatezone

type AddParsingRecordsRequest struct {
	ZoneId      *string `json:"-"`
	ClientToken *string `json:"-"`
	Rr          *string `json:"rr,omitempty"`
	Value       *string `json:"value,omitempty"`
	Type        *string `json:"type,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
	Ttl         *int32  `json:"ttl,omitempty"`
	Description *string `json:"description,omitempty"`
}
