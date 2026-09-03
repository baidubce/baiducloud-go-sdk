package image

type IColorParams struct {
	Shadow            *float64     `json:"shadow,omitempty"`
	SmartRemoveFog    *float64     `json:"smart_remove_fog,omitempty"`
	Tint              *float64     `json:"tint,omitempty"`
	SkinColorRefresh  *int32       `json:"skin_color_refresh,omitempty"`
	AiColor           *int32       `json:"ai_color,omitempty"`
	SmartExposure     *float64     `json:"smart_exposure,omitempty"`
	Saturation        *float64     `json:"saturation,omitempty"`
	Highlight         *float64     `json:"highlight,omitempty"`
	BgEnhance         *float64     `json:"bg_enhance,omitempty"`
	White             *float64     `json:"white,omitempty"`
	SharpenAmount     *float64     `json:"sharpen_amount,omitempty"`
	Temperature       *float64     `json:"temperature,omitempty"`
	LutValue          *float64     `json:"lut_value,omitempty"`
	AutoWhitebalance  *float64     `json:"auto_whitebalance,omitempty"`
	SharpenRadius     *float64     `json:"sharpen_radius,omitempty"`
	Black             *float64     `json:"black,omitempty"`
	HslParams         []*HslParams `json:"hsl_params,omitempty"`
	AutoExposure      *float64     `json:"auto_exposure,omitempty"`
	Brightness        *float64     `json:"brightness,omitempty"`
	Exposure          *float64     `json:"exposure,omitempty"`
	Contrast          *float64     `json:"contrast,omitempty"`
	Vibrance          *float64     `json:"vibrance,omitempty"`
	SmartWhitebalance *float64     `json:"smart_whitebalance,omitempty"`
	RemoveFog         *float64     `json:"remove_fog,omitempty"`
	LutId             *string      `json:"lut_id,omitempty"`
}
