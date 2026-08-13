package aigw

type RegexRewriteConfig struct {
	Match   *string `json:"match,omitempty"`
	Rewrite *string `json:"rewrite,omitempty"`
}
