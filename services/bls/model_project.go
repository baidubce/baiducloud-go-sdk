package bls

type Project struct {
	Uuid          *string `json:"uuid,omitempty"`
	Name          *string `json:"name,omitempty"`
	Description   *string `json:"description,omitempty"`
	Top           *bool   `json:"top,omitempty"`
	LogStoreCount *int32  `json:"logStoreCount,omitempty"`
	CreatedTime   *string `json:"createdTime,omitempty"`
	UpdatedTime   *string `json:"updatedTime,omitempty"`
}
