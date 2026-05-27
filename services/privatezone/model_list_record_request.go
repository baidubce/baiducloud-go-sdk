package privatezone

type ListRecordRequest struct {
	ZoneId     *string `json:"-"`
	Marker     *string `json:"-"`
	MaxKeys    *int32  `json:"-"`
	Rr         *string `json:"-"`
	SearchMode *string `json:"-"`
	Type       *string `json:"-"`
	Value      *string `json:"-"`
}
