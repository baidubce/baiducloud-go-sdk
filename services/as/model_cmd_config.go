package as

type CmdConfig struct {
	HasDecreaseCmd *bool   `json:"hasDecreaseCmd,omitempty"`
	DecCmdStrategy *string `json:"decCmdStrategy,omitempty"`
	DecCmdData     *string `json:"decCmdData,omitempty"`
	DecCmdTimeout  *int32  `json:"decCmdTimeout,omitempty"`
	DecCmdManual   *bool   `json:"decCmdManual,omitempty"`
	HasIncreaseCmd *bool   `json:"hasIncreaseCmd,omitempty"`
	IncCmdStrategy *string `json:"incCmdStrategy,omitempty"`
	IncCmdData     *string `json:"incCmdData,omitempty"`
	IncCmdTimeout  *int32  `json:"incCmdTimeout,omitempty"`
	IncCmdManual   *bool   `json:"incCmdManual,omitempty"`
}
