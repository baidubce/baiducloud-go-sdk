package scs

type GetSystemParameterListRequest struct {
	Engine        *string `json:"-"`
	EngineVersion *string `json:"-"`
	ClusterType   *string `json:"-"`
}
