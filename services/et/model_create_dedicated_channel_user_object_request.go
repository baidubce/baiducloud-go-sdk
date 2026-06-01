package et

type CreateDedicatedChannelUserObjectRequest struct {
	EtId            *string   `json:"-"`
	EtChannelId     *string   `json:"-"`
	ClientToken     *string   `json:"-"`
	AuthorizedUsers []*string `json:"authorizedUsers,omitempty"`
}
