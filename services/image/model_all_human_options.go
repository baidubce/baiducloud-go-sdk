package image

type AllHumanOptions struct {
	BodyHeighten     *float64 `json:"body_heighten,omitempty"`
	RemoveBgFlaw     *float64 `json:"remove_bg_flaw,omitempty"`
	LegLong          *float64 `json:"leg_long,omitempty"`
	AllSkinColorSame *float64 `json:"all_skin_color_same,omitempty"`
	RemovePureBgFlaw *float64 `json:"remove_pure_bg_flaw,omitempty"`
}
