package ccr

type ReplicationRegistry struct {
	CreationTime *string             `json:"creationTime,omitempty"`
	Credential   *RegistryCredential `json:"credential,omitempty"`
	Description  *string             `json:"description,omitempty"`
	Id           *int32              `json:"id,omitempty"`
	Insecure     *bool               `json:"insecure,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Status       *string             `json:"status,omitempty"`
	CcrType      *string             `json:"type,omitempty"`
	UpdateTime   *string             `json:"updateTime,omitempty"`
	Url          *string             `json:"url,omitempty"`
	Region       *string             `json:"region,omitempty"`
}
