package oos

type GetExecutionDetailV2Request struct {
	Namespace *string `json:"-"`
	Id        *string `json:"-"`
	WithLog   *string `json:"-"`
	Locale    *string `json:"-"`
}
