package bcc

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_IMAGE = "image"

	CONSTANT_TAG = "tag"

	CONSTANT_VOLUME = "volume"

	CONSTANT_PROGRESS = "progress"

	CONSTANT_DISK = "disk"

	CONSTANT_QUOTA = "quota"

	CONSTANT_GET_PRICE = "getPrice"

	CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC = "getAvailableImagesBySpec"

	CONSTANT_RENAME = "rename"

	CONSTANT_IMPORT = "import"

	CONSTANT_OS = "os"

	CONSTANT_SHARED_USERS = "sharedUsers"
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
func getBindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getBindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
func getCancelRemoteCopyImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getCreateImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE
}
func getCreateVolumeUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME
}
func getDeleteImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getDetachVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetAvailableImagesBySpecUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_GET_AVAILABLE_IMAGES_BY_SPEC
}
func getGetCdsPriceUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_GET_PRICE
}
func getGetDiskQuotaUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_DISK + bce.URI_PREFIX + CONSTANT_QUOTA
}
func getGetImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getGetVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId
}
func getGetVolumeResizeProgressUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + CONSTANT_PROGRESS + bce.URI_PREFIX + VolumeId
}
func getImportImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_IMPORT
}
func getListImagesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE
}
func getListOsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_OS
}
func getListSharedUserUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_SHARED_USERS
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
func getRemoteCopyImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getRenameImageUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + CONSTANT_RENAME
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
func getShareImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getUnShareImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId
}
func getUnbindTagImageUri(version string, ImageId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_IMAGE + bce.URI_PREFIX + ImageId + bce.URI_PREFIX + CONSTANT_TAG
}
func getUnbindTagVolumeUri(version string, VolumeId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_VOLUME + bce.URI_PREFIX + VolumeId + bce.URI_PREFIX + CONSTANT_TAG
}
