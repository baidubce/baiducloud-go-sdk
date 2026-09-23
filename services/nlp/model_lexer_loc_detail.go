package nlp

type LexerLocDetail struct {
	NlpType    *string `json:"type,omitempty"`
	ByteOffset *int32  `json:"byte_offset,omitempty"`
	ByteLength *int32  `json:"byte_length,omitempty"`
}
