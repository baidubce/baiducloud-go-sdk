package pfs

import "github.com/baidubce/baiducloud-go-sdk/bce"

type InstanceAvailableVersionsResponse struct {
	bce.BaseResponse
	RequestId            *string             `json:"requestId,omitempty"`
	AvailableVersions    []*AvailableVersion `json:"availableVersions,omitempty"`
	PremiumVersionNumber *string             `json:"premiumVersionNumber,omitempty"`
	VersionPinned        *bool               `json:"versionPinned,omitempty"`
}
