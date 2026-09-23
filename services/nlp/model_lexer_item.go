package nlp

type LexerItem struct {
	Item       *string           `json:"item,omitempty"`
	Ne         *string           `json:"ne,omitempty"`
	Pos        *string           `json:"pos,omitempty"`
	ByteOffset *int32            `json:"byte_offset,omitempty"`
	ByteLength *int32            `json:"byte_length,omitempty"`
	Uri        *string           `json:"uri,omitempty"`
	Formal     *string           `json:"formal,omitempty"`
	BasicWords []*string         `json:"basic_words,omitempty"`
	LocDetails []*LexerLocDetail `json:"loc_details,omitempty"`
}
