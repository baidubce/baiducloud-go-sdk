package image

type IngredientRequest struct {
	Image  *string `json:"image,omitempty"`
	Url    *string `json:"url,omitempty"`
	TopNum *int32  `json:"top_num,omitempty"`
}
