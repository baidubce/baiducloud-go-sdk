package bls

type ProcessConfig struct {
	Regex            *string `json:"regex,omitempty"`
	Separator        *string `json:"separator,omitempty"`
	CustomSeparator  *string `json:"customSeparator,omitempty"`
	Quote            *string `json:"quote,omitempty"`
	KvKeyIndex       *int32  `json:"kvKeyIndex,omitempty"`
	KvValueIndex     *int32  `json:"kvValueIndex,omitempty"`
	SampleLog        *string `json:"sampleLog,omitempty"`
	Keys             *string `json:"keys,omitempty"`
	DataType         *string `json:"dataType,omitempty"`
	DiscardOnFailure *bool   `json:"discardOnFailure,omitempty"`
	KeepOriginal     *bool   `json:"keepOriginal,omitempty"`
}
