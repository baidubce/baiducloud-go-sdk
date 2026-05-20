package dns

type PublicRecord struct {
	Id          *string `json:"id,omitempty"`
	Rr          *string `json:"rr,omitempty"`
	Status      *string `json:"status,omitempty"`
	DnsType     *string `json:"type,omitempty"`
	Value       *string `json:"value,omitempty"`
	Ttl         *int32  `json:"ttl,omitempty"`
	Line        *string `json:"line,omitempty"`
	Description *string `json:"description,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
}
