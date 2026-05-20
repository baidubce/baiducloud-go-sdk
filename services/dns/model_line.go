package dns

type Line struct {
	Id                 *string   `json:"id,omitempty"`
	Name               *string   `json:"name,omitempty"`
	Lines              []*string `json:"lines,omitempty"`
	RelatedZoneCount   *int32    `json:"relatedZoneCount,omitempty"`
	RelatedRecordCount *int32    `json:"relatedRecordCount,omitempty"`
}
