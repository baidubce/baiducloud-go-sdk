package cprom

type LabelSelector struct {
	MatchLabels *map[string]string `json:"matchLabels,omitempty"`
}
