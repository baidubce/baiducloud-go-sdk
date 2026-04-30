package privatezone

type PrivateRecord struct {
	RecordId    *string `json:"recordId,omitempty"`
	Rr          *string `json:"rr,omitempty"`
	Value       *string `json:"value,omitempty"`
	Status      *string `json:"status,omitempty"`
	Type        *string `json:"type,omitempty"`
	Ttl         *int32  `json:"ttl,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
	Description *string `json:"description,omitempty"`
}
