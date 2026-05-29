package cfs

type QueryPermissionGroupRulesRequest struct {
	AgName  *string `json:"-"`
	ArId    *string `json:"-"`
	Marker  *string `json:"-"`
	MaxKeys *int32  `json:"-"`
}
