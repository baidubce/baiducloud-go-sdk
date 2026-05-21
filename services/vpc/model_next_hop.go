package vpc

type NextHop struct {
	NexthopId   *string `json:"nexthopId,omitempty"`
	NexthopType *string `json:"nexthopType,omitempty"`
	PathType    *string `json:"pathType,omitempty"`
}
