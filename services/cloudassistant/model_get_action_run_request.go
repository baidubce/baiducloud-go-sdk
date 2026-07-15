package cloudassistant

type GetActionRunRequest struct {
	Id            *string `json:"-"`
	WithLog       *string `json:"-"`
	PageNo        *int32  `json:"-"`
	PageSize      *int32  `json:"-"`
	ChildRunState *string `json:"-"`
	Locale        *string `json:"-"`
}
