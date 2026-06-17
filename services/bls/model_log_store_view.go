package bls

type LogStoreView struct {
	Project          *string     `json:"project,omitempty"`
	Name             *string     `json:"name,omitempty"`
	Logstores        []*LogStore `json:"logstores,omitempty"`
	CreatedTimestamp *string     `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp *string     `json:"updatedTimestamp,omitempty"`
}
