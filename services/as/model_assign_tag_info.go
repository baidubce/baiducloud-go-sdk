package as

type AssignTagInfo struct {
	ResourceId  *string    `json:"resourceId,omitempty"`
	RelationTag *bool      `json:"relationTag,omitempty"`
	Tags        []*TagInfo `json:"tags,omitempty"`
}
