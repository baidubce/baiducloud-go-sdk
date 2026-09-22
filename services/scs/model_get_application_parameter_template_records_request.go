package scs

type GetApplicationParameterTemplateRecordsRequest struct {
	TemplateShowId *string `json:"-"`
	Marker         *string `json:"-"`
	MaxKeys        *int32  `json:"-"`
}
