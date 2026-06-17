package bls

type CreateTaskRequest struct {
	Name   *string     `json:"name,omitempty"`
	Config *TaskConfig `json:"config,omitempty"`
	Hosts  []*Host     `json:"hosts,omitempty"`
	Tags   []*Tag      `json:"tags,omitempty"`
}
