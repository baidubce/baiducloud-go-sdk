package oos

type GetTemplateDetailV2Request struct {
	Namespace *string `json:"-"`
	Id        *string `json:"-"`
	Name      *string `json:"-"`
	Type      *string `json:"-"`
	Locale    *string `json:"-"`
}
