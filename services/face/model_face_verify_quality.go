package face

type FaceVerifyQuality struct {
	Occlusion    *FaceVerifyOcclusion `json:"occlusion,omitempty"`
	Blur         *float64             `json:"blur,omitempty"`
	Illumination *float64             `json:"illumination,omitempty"`
	Completeness *int64               `json:"completeness,omitempty"`
}
