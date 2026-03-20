package vpc

type IkeConfig struct {
	IkeVersion  *string `json:"ikeVersion,omitempty"`
	IkeMode     *string `json:"ikeMode,omitempty"`
	IkeEncAlg   *string `json:"ikeEncAlg,omitempty"`
	IkeAuthAlg  *string `json:"ikeAuthAlg,omitempty"`
	IkePfs      *string `json:"ikePfs,omitempty"`
	IkeLifeTime *string `json:"ikeLifeTime,omitempty"`
}
