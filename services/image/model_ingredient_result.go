package image

type IngredientResult struct {
	Name  *string  `json:"name,omitempty"`
	Score *float64 `json:"score,omitempty"`
}
