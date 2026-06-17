package bcm

type Filter struct {
	Key    *string   `json:"key,omitempty"`
	Op     *string   `json:"op,omitempty"`
	Value  *string   `json:"value,omitempty"`
	Values []*string `json:"values,omitempty"`
}
