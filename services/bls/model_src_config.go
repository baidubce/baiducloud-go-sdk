package bls

type SrcConfig struct {
	SrcType        *string                 `json:"srcType,omitempty"`
	LogType        *string                 `json:"logType,omitempty"`
	SrcDir         *string                 `json:"srcDir,omitempty"`
	MatchedPattern *string                 `json:"matchedPattern,omitempty"`
	IgnorePattern  *string                 `json:"ignorePattern,omitempty"`
	TimeFormat     *string                 `json:"timeFormat,omitempty"`
	Ttl            *int32                  `json:"ttl,omitempty"`
	UseMultiline   *bool                   `json:"useMultiline,omitempty"`
	MultilineRegex *string                 `json:"multilineRegex,omitempty"`
	RecursiveDir   *bool                   `json:"recursiveDir,omitempty"`
	ProcessType    *string                 `json:"processType,omitempty"`
	ProcessConfig  *ProcessConfig          `json:"processConfig,omitempty"`
	LogTime        *string                 `json:"logTime,omitempty"`
	TimestampKey   *string                 `json:"timestampKey,omitempty"`
	DateFormat     *string                 `json:"dateFormat,omitempty"`
	FilterExpr     *string                 `json:"filterExpr,omitempty"`
	AdditionConfig *map[string]interface{} `json:"additionConfig,omitempty"`
	MetaEnv        []*string               `json:"metaEnv,omitempty"`
	MetaLabel      []*string               `json:"metaLabel,omitempty"`
	MetaContainer  []*string               `json:"metaContainer,omitempty"`
	MetaToFields   *bool                   `json:"metaToFields,omitempty"`
	HarvesterLimit *int32                  `json:"HarvesterLimit,omitempty"`
}
