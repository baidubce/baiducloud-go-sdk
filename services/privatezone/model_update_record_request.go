package privatezone

type UpdateRecordRequest struct {
	RecordId        *string `json:"-"`
	ClientToken     *string `json:"-"`
	Rr              *string `json:"rr,omitempty"`
	Value           *string `json:"value,omitempty"`
	PrivatezoneType *string `json:"type,omitempty"`
	Ttl             *int32  `json:"ttl,omitempty"`
	Priority        *int32  `json:"priority,omitempty"`
	Description     *string `json:"description,omitempty"`
}
