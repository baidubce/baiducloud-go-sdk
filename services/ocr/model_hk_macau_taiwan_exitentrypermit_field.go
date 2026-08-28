package ocr

type HkMacauTaiwanExitentrypermitField struct {
	Word        *string                                  `json:"word,omitempty"`
	Location    *HkMacauTaiwanExitentrypermitLocation    `json:"location,omitempty"`
	Probability *HkMacauTaiwanExitentrypermitProbability `json:"probability,omitempty"`
}
