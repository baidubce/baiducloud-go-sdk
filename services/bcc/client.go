package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_VOLUME = "volume"

	CONSTANT_TAG = "tag"

	CONSTANT_PROGRESS = "progress"

	CONSTANT_DISK = "disk"

	CONSTANT_QUOTA = "quota"

	CONSTANT_GET_PRICE = "getPrice"
)

// Client of bcc service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getAttachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getBindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getCreateVolumeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getDetachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetCdsPriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_GET_PRICE
}
func getGetDiskQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_DISK + bce.URI_PREFIX + CONSTANT_QUOTA
}
func getGetVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetVolumeResizeProgressUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_PROGRESS + bce.URI_PREFIX + VolumeId
}
func getListVolumesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getModifyCdsAttributeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getModifyVolumeChargeTypeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getPurchaseReservedVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getReleaseVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getRenameVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getResizeVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getRollbackVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getUnbindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
