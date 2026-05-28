package dns

type CreateRecordRequest struct {
	ZoneName    *string `json:"-"`
	ClientToken *string `json:"-"`
	Rr          *string `json:"rr,omitempty"`
	DnsType     *string `json:"type,omitempty"`
	Value       *string `json:"value,omitempty"`
	Ttl         *int32  `json:"ttl,omitempty"`
	Line        *string `json:"line,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
}
