package bcm

type UpdateAlarmMaskingStatesRequest struct {
	Ids   []*string `json:"ids,omitempty"`
	State *string   `json:"state,omitempty"`
}
