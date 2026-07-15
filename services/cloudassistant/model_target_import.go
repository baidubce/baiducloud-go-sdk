package cloudassistant

type TargetImport struct {
	KeywordType *string   `json:"keywordType,omitempty"`
	Instances   []*string `json:"instances,omitempty"`
}
