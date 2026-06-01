package apm

type ExceptionDetail struct {
	Id         *string `json:"id,omitempty"`
	Stacktrace *string `json:"stacktrace,omitempty"`
	Service    *string `json:"service,omitempty"`
}
