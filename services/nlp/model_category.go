package nlp

type Category struct {
	Level1 *string `json:"level_1,omitempty"`
	Level2 *string `json:"level_2,omitempty"`
	Level3 *string `json:"level_3,omitempty"`
}
