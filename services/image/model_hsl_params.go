package image

type HslParams struct {
	HslBrightness *float64 `json:"hsl_brightness,omitempty"`
	HslSaturation *float64 `json:"hsl_saturation,omitempty"`
	HslGamut      *int32   `json:"hsl_gamut,omitempty"`
	HslHue        *float64 `json:"hsl_hue,omitempty"`
}
