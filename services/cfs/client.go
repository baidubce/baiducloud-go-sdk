package cfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "cfs." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_ACCESS_GROUP = "accessGroup"

	CONSTANT_CFS = "cfs"

	CONSTANT_BATCH_CREATE_ACCESS_RULE = "batchCreateAccessRule"

	CONSTANT_CLIENTS = "clients"
)

// Client of cfs service is a kind of BceClient, so derived from BceClient
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

func getBatchCreationOfPermissionGroupRulesUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + CONSTANT_BATCH_CREATE_ACCESS_RULE
}
func getCreateFileSystemUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS
}
func getCreateMountingTargetUri(version string, FsId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FsId
}
func getCreatePermissionGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP
}
func getCreatePermissionGroupRulesUri(version string, AgName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName
}
func getDeletePermissionGroupUri(version string, AgName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName
}
func getDeletePermissionGroupRuleUri(version string, AgName string, ArId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName + bce.URI_PREFIX + ArId
}
func getDropFileSystemUri(version string, FsId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FsId
}
func getDropMountTargetUri(version string, FsId string, MountId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FsId + bce.URI_PREFIX + MountId
}
func getModifyTheMountTargetPermissionGroupUri(version string, FsId string, MountID string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FsId + bce.URI_PREFIX + MountID
}
func getQueryFileSystemUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS
}
func getQueryMountedClientUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + CONSTANT_CLIENTS
}
func getQueryMountingTargetUri(version string, FId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FId
}
func getQueryPermissionGroupUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP
}
func getQueryPermissionGroupRulesUri(version string, AgName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName
}
func getUpdateFileSystemUri(version string, FsId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS + bce.URI_PREFIX + FsId
}
func getUpdateFileSystemLabelsUri(version string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_CFS
}
func getUpdatePermissionGroupUri(version string, AgName string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName
}
func getUpdatePermissionGroupRulesUri(version string, AgName string, ArId string) string {
	return bce.URI_PREFIX + version + bce.URI_PREFIX + CONSTANT_ACCESS_GROUP + bce.URI_PREFIX + AgName + bce.URI_PREFIX + ArId
}
