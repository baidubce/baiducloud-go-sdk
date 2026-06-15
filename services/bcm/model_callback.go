package bcm

type Callback struct {
	Url     *string  `json:"url,omitempty"`
	Mention *Mention `json:"mention,omitempty"`
}
