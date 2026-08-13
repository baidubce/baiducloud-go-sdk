package aigw

type ResponseHeaders struct {
	Enabled *bool           `json:"enabled,omitempty"`
	Headers []*CustomHeader `json:"headers,omitempty"`
}
