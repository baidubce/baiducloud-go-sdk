package pfs

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// CancelL2BucketLink
//
// PARAMS:
//   - request: the arguments to CancelL2BucketLink
//
// RETURNS:
//   - CancelL2BucketLinkResponse: The return type of the CancelL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CancelL2BucketLink(request *CancelL2BucketLinkRequest) (*CancelL2BucketLinkResponse, error) {
	result := &CancelL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCancelL2BucketLinkUri()).
		WithQueryParamFilter("action", "CancelL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFileset
//
// PARAMS:
//   - request: the arguments to CreateFileset
//
// RETURNS:
//   - CreateFilesetResponse: The return type of the CreateFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateFileset(request *CreateFilesetRequest) (*CreateFilesetResponse, error) {
	result := &CreateFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateFilesetUri()).
		WithQueryParamFilter("action", "CreateFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateL2BucketLink
//
// PARAMS:
//   - request: the arguments to CreateL2BucketLink
//
// RETURNS:
//   - CreateL2BucketLinkResponse: The return type of the CreateL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateL2BucketLink(request *CreateL2BucketLinkRequest) (*CreateL2BucketLinkResponse, error) {
	result := &CreateL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateL2BucketLinkUri()).
		WithQueryParamFilter("action", "CreateL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateL2Policy
//
// PARAMS:
//   - request: the arguments to CreateL2Policy
//
// RETURNS:
//   - CreateL2PolicyResponse: The return type of the CreateL2Policy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateL2Policy(request *CreateL2PolicyRequest) (*CreateL2PolicyResponse, error) {
	result := &CreateL2PolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateL2PolicyUri()).
		WithQueryParamFilter("action", "CreateL2Policy").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreatePfs
//
// PARAMS:
//   - request: the arguments to CreatePfs
//
// RETURNS:
//   - CreatePfsResponse: The return type of the CreatePfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePfs(request *CreatePfsRequest) (*CreatePfsResponse, error) {
	result := &CreatePfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreatePfsUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteFileset
//
// PARAMS:
//   - request: the arguments to DeleteFileset
//
// RETURNS:
//   - DeleteFilesetResponse: The return type of the DeleteFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteFileset(request *DeleteFilesetRequest) (*DeleteFilesetResponse, error) {
	result := &DeleteFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteFilesetUri()).
		WithQueryParamFilter("action", "DeleteFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteL2BucketLink
//
// PARAMS:
//   - request: the arguments to DeleteL2BucketLink
//
// RETURNS:
//   - DeleteL2BucketLinkResponse: The return type of the DeleteL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteL2BucketLink(request *DeleteL2BucketLinkRequest) (*DeleteL2BucketLinkResponse, error) {
	result := &DeleteL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteL2BucketLinkUri()).
		WithQueryParamFilter("action", "DeleteL2BucketLink").
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteL2Policy
//
// PARAMS:
//   - request: the arguments to DeleteL2Policy
//
// RETURNS:
//   - DeleteL2PolicyResponse: The return type of the DeleteL2Policy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteL2Policy(request *DeleteL2PolicyRequest) (*DeleteL2PolicyResponse, error) {
	result := &DeleteL2PolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDeleteL2PolicyUri()).
		WithQueryParamFilter("action", "DeleteL2Policy").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeletePfs
//
// PARAMS:
//   - request: the arguments to DeletePfs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeletePfs(request *DeletePfsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeletePfsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		Do()
}

// DescFileset
//
// PARAMS:
//   - request: the arguments to DescFileset
//
// RETURNS:
//   - DescFilesetResponse: The return type of the DescFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescFileset(request *DescFilesetRequest) (*DescFilesetResponse, error) {
	result := &DescFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescFilesetUri()).
		WithQueryParamFilter("action", "DescribeFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescL2BucketLink
//
// PARAMS:
//   - request: the arguments to DescL2BucketLink
//
// RETURNS:
//   - DescL2BucketLinkResponse: The return type of the DescL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescL2BucketLink(request *DescL2BucketLinkRequest) (*DescL2BucketLinkResponse, error) {
	result := &DescL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescL2BucketLinkUri()).
		WithQueryParamFilter("action", "DescribeL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescL2Policy
//
// PARAMS:
//   - request: the arguments to DescL2Policy
//
// RETURNS:
//   - DescL2PolicyResponse: The return type of the DescL2Policy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescL2Policy(request *DescL2PolicyRequest) (*DescL2PolicyResponse, error) {
	result := &DescL2PolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDescL2PolicyUri()).
		WithQueryParamFilter("action", "DescribeL2Policy").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DescPfs
//
// PARAMS:
//   - request: the arguments to DescPfs
//
// RETURNS:
//   - DescPfsResponse: The return type of the DescPfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DescPfs(request *DescPfsRequest) (*DescPfsResponse, error) {
	result := &DescPfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescPfsUri()).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceListClients
//
// PARAMS:
//   - request: the arguments to InstanceListClients
//
// RETURNS:
//   - InstanceListClientsResponse: The return type of the InstanceListClients interface.
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceListClients(request *InstanceListClientsRequest) (*InstanceListClientsResponse, error) {
	result := &InstanceListClientsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceListClientsUri()).
		WithQueryParamFilter("action", "InstanceListClients").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListFileset
//
// PARAMS:
//   - request: the arguments to ListFileset
//
// RETURNS:
//   - ListFilesetResponse: The return type of the ListFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListFileset(request *ListFilesetRequest) (*ListFilesetResponse, error) {
	result := &ListFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListFilesetUri()).
		WithQueryParamFilter("action", "ListFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListL2BucketLink
//
// PARAMS:
//   - request: the arguments to ListL2BucketLink
//
// RETURNS:
//   - ListL2BucketLinkResponse: The return type of the ListL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListL2BucketLink(request *ListL2BucketLinkRequest) (*ListL2BucketLinkResponse, error) {
	result := &ListL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListL2BucketLinkUri()).
		WithQueryParamFilter("action", "ListL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListL2Policy
//
// PARAMS:
//   - request: the arguments to ListL2Policy
//
// RETURNS:
//   - ListL2PolicyResponse: The return type of the ListL2Policy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListL2Policy(request *ListL2PolicyRequest) (*ListL2PolicyResponse, error) {
	result := &ListL2PolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getListL2PolicyUri()).
		WithQueryParamFilter("action", "ListL2Policy").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListPfs
//
// PARAMS:
//   - request: the arguments to ListPfs
//
// RETURNS:
//   - ListPfsResponse: The return type of the ListPfs interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListPfs(request *ListPfsRequest) (*ListPfsResponse, error) {
	result := &ListPfsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListPfsUri()).
		WithQueryParamFilter("manner", "marker").
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("manner", util.StringValue(request.Manner)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("filterTag", util.StringValue(request.FilterTag)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LstPerL2BktLnkExecLog
//
// PARAMS:
//   - request: the arguments to LstPerL2BktLnkExecLog
//
// RETURNS:
//   - LstPerL2BktLnkExecLogResponse: The return type of the LstPerL2BktLnkExecLog interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LstPerL2BktLnkExecLog(request *LstPerL2BktLnkExecLogRequest) (*LstPerL2BktLnkExecLogResponse, error) {
	result := &LstPerL2BktLnkExecLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLstPerL2BktLnkExecLogUri()).
		WithQueryParamFilter("action", "ListPeriodL2BucketLinkExecuteLog").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MountTargetListClients
//
// PARAMS:
//   - request: the arguments to MountTargetListClients
//
// RETURNS:
//   - MountTargetListClientsResponse: The return type of the MountTargetListClients interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MountTargetListClients(request *MountTargetListClientsRequest) (*MountTargetListClientsResponse, error) {
	result := &MountTargetListClientsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMountTargetListClientsUri()).
		WithQueryParamFilter("action", "MountTargetListClients").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PauseL2BucketLink
//
// PARAMS:
//   - request: the arguments to PauseL2BucketLink
//
// RETURNS:
//   - PauseL2BucketLinkResponse: The return type of the PauseL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PauseL2BucketLink(request *PauseL2BucketLinkRequest) (*PauseL2BucketLinkResponse, error) {
	result := &PauseL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPauseL2BucketLinkUri()).
		WithQueryParamFilter("action", "PauseL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QryL2PolExecDetail
//
// PARAMS:
//   - request: the arguments to QryL2PolExecDetail
//
// RETURNS:
//   - QryL2PolExecDetailResponse: The return type of the QryL2PolExecDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QryL2PolExecDetail(request *QryL2PolExecDetailRequest) (*QryL2PolExecDetailResponse, error) {
	result := &QryL2PolExecDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQryL2PolExecDetailUri()).
		WithQueryParamFilter("action", "QueryL2PolicyExecuteDetail").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QryL2PolExecLog
//
// PARAMS:
//   - request: the arguments to QryL2PolExecLog
//
// RETURNS:
//   - QryL2PolExecLogResponse: The return type of the QryL2PolExecLog interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QryL2PolExecLog(request *QryL2PolExecLogRequest) (*QryL2PolExecLogResponse, error) {
	result := &QryL2PolExecLogResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQryL2PolExecLogUri()).
		WithQueryParamFilter("action", "QueryL2PolicyExecuteLog").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResumeL2BucketLink
//
// PARAMS:
//   - request: the arguments to ResumeL2BucketLink
//
// RETURNS:
//   - ResumeL2BucketLinkResponse: The return type of the ResumeL2BucketLink interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResumeL2BucketLink(request *ResumeL2BucketLinkRequest) (*ResumeL2BucketLinkResponse, error) {
	result := &ResumeL2BucketLinkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResumeL2BucketLinkUri()).
		WithQueryParamFilter("action", "ResumeL2BucketLink").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdPerL2BktLnkInfo
//
// PARAMS:
//   - request: the arguments to UpdPerL2BktLnkInfo
//
// RETURNS:
//   - UpdPerL2BktLnkInfoResponse: The return type of the UpdPerL2BktLnkInfo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdPerL2BktLnkInfo(request *UpdPerL2BktLnkInfoRequest) (*UpdPerL2BktLnkInfoResponse, error) {
	result := &UpdPerL2BktLnkInfoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdPerL2BktLnkInfoUri()).
		WithQueryParamFilter("action", "UpdatePeriodL2BucketLinkInfo").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFileset
//
// PARAMS:
//   - request: the arguments to UpdateFileset
//
// RETURNS:
//   - UpdateFilesetResponse: The return type of the UpdateFileset interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateFileset(request *UpdateFilesetRequest) (*UpdateFilesetResponse, error) {
	result := &UpdateFilesetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateFilesetUri()).
		WithQueryParamFilter("action", "UpdateFileset").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateL2Policy
//
// PARAMS:
//   - request: the arguments to UpdateL2Policy
//
// RETURNS:
//   - UpdateL2PolicyResponse: The return type of the UpdateL2Policy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateL2Policy(request *UpdateL2PolicyRequest) (*UpdateL2PolicyResponse, error) {
	result := &UpdateL2PolicyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateL2PolicyUri()).
		WithQueryParamFilter("action", "UpdateL2Policy").
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithHeader("Version", "v2").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdatePFSTag
//
// PARAMS:
//   - request: the arguments to UpdatePFSTag
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdatePFSTag(request *UpdatePFSTagRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdatePFSTagUri()).
		WithBody(request).
		Do()
}
