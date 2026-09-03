package image

type AiRetouchingCreateTaskRequest struct {
	Image               *string              `json:"image,omitempty"`
	Url                 *string              `json:"url,omitempty"`
	CallbackData        *string              `json:"callback_data,omitempty"`
	IColorParams        *IColorParams        `json:"IColorParams,omitempty"`
	AllHumanOptions     *AllHumanOptions     `json:"AllHumanOptions,omitempty"`
	PartialHumanOptions *PartialHumanOptions `json:"PartialHumanOptions,omitempty"`
	PartialTemplates    *PartialTemplates    `json:"PartialTemplates,omitempty"`
	TransformOptions    *TransformOptions    `json:"transform_options,omitempty"`
}
