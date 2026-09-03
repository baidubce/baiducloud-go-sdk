package image

type DishResult struct {
	Name        *string    `json:"name,omitempty"`
	Calorie     *string    `json:"calorie,omitempty"`
	Probability *string    `json:"probability,omitempty"`
	HasCalorie  *bool      `json:"has_calorie,omitempty"`
	BaikeInfo   *BaikeInfo `json:"baike_info,omitempty"`
}
