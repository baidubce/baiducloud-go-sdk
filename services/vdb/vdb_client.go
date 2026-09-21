package vdb

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AccountListUsingGET
//
// PARAMS:
//   - request: the arguments to AccountListUsingGET
//
// RETURNS:
//   - AccountListUsingGETResponse: The return type of the AccountListUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AccountListUsingGET(request *AccountListUsingGETRequest) (*AccountListUsingGETResponse, error) {
	result := &AccountListUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAccountListUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindEipUsingPOST
//
// PARAMS:
//   - request: the arguments to BindEipUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindEipUsingPOST(request *BindEipUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindEipUsingPOSTUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// CreateInstanceUsingPOST
//
// PARAMS:
//   - request: the arguments to CreateInstanceUsingPOST
//
// RETURNS:
//   - CreateInstanceUsingPOSTResponse: The return type of the CreateInstanceUsingPOST interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceUsingPOST(request *CreateInstanceUsingPOSTRequest) (*CreateInstanceUsingPOSTResponse, error) {
	result := &CreateInstanceUsingPOSTResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateInstanceUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteInstanceUsingDELETE
//
// PARAMS:
//   - request: the arguments to DeleteInstanceUsingDELETE
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceUsingDELETE(request *DeleteInstanceUsingDELETERequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteInstanceUsingDELETEUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		Do()
}

// DeleteRecordUsingDELETE
//
// PARAMS:
//   - request: the arguments to DeleteRecordUsingDELETE
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRecordUsingDELETE(request *DeleteRecordUsingDELETERequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRecordUsingDELETEUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithQueryParamFilter("batchId", util.StringValue(request.BatchId)).
		WithQueryParamFilter("backupId", util.StringValue(request.BackupId)).
		Do()
}

// DeleteRecyclerInstance
//
// PARAMS:
//   - request: the arguments to DeleteRecyclerInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteRecyclerInstance(request *DeleteRecyclerInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteRecyclerInstanceUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		Do()
}

// DescribeInstanceConfigs
//
// PARAMS:
//   - request: the arguments to DescribeInstanceConfigs
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DescribeInstanceConfigs(request *DescribeInstanceConfigsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDescribeInstanceConfigsUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		Do()
}

// GetConfigUsingGET
//
// PARAMS:
//   - request: the arguments to GetConfigUsingGET
//
// RETURNS:
//   - GetConfigUsingGETResponse: The return type of the GetConfigUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetConfigUsingGET(request *GetConfigUsingGETRequest) (*GetConfigUsingGETResponse, error) {
	result := &GetConfigUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetConfigUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFreeInstanceQuota
//
// PARAMS:
//   - request: the arguments to GetFreeInstanceQuota
//
// RETURNS:
//   - GetFreeInstanceQuotaResponse: The return type of the GetFreeInstanceQuota interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetFreeInstanceQuota() (*GetFreeInstanceQuotaResponse, error) {
	result := &GetFreeInstanceQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetFreeInstanceQuotaUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInstanceListUsingGET
//
// PARAMS:
//   - request: the arguments to GetInstanceListUsingGET
//
// RETURNS:
//   - GetInstanceListUsingGETResponse: The return type of the GetInstanceListUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstanceListUsingGET(request *GetInstanceListUsingGETRequest) (*GetInstanceListUsingGETResponse, error) {
	result := &GetInstanceListUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetInstanceListUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithQueryParamFilter("instanceType", util.StringValue(request.InstanceType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNodeSpecListUsingGET
//
// PARAMS:
//   - request: the arguments to GetNodeSpecListUsingGET
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) GetNodeSpecListUsingGET(request *GetNodeSpecListUsingGETRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetNodeSpecListUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		Do()
}

// GetPriceUsingPOST
//
// PARAMS:
//   - request: the arguments to GetPriceUsingPOST
//
// RETURNS:
//   - GetPriceUsingPOSTResponse: The return type of the GetPriceUsingPOST interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPriceUsingPOST(request *GetPriceUsingPOSTRequest) (*GetPriceUsingPOSTResponse, error) {
	result := &GetPriceUsingPOSTResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGetPriceUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetQuotaUsingGET
//
// PARAMS:
//   - request: the arguments to GetQuotaUsingGET
//
// RETURNS:
//   - GetQuotaUsingGETResponse: The return type of the GetQuotaUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetQuotaUsingGET(request *GetQuotaUsingGETRequest) (*GetQuotaUsingGETResponse, error) {
	result := &GetQuotaUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetQuotaUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTLSCertificateUsingGET
//
// PARAMS:
//   - request: the arguments to GetTLSCertificateUsingGET
//
// RETURNS:
//   - GetTLSCertificateUsingGETResponse: The return type of the GetTLSCertificateUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTLSCertificateUsingGET(request *GetTLSCertificateUsingGETRequest) (*GetTLSCertificateUsingGETResponse, error) {
	result := &GetTLSCertificateUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTLSCertificateUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTLSInfoUsingGET
//
// PARAMS:
//   - request: the arguments to GetTLSInfoUsingGET
//
// RETURNS:
//   - GetTLSInfoUsingGETResponse: The return type of the GetTLSInfoUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetTLSInfoUsingGET(request *GetTLSInfoUsingGETRequest) (*GetTLSInfoUsingGETResponse, error) {
	result := &GetTLSInfoUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetTLSInfoUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Getinstancelistusingget1
//
// PARAMS:
//   - request: the arguments to Getinstancelistusingget1
//
// RETURNS:
//   - Getinstancelistusingget1Response: The return type of the Getinstancelistusingget1 interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Getinstancelistusingget1(request *Getinstancelistusingget1Request) (*Getinstancelistusingget1Response, error) {
	result := &Getinstancelistusingget1Response{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetinstancelistusingget1Uri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceDetailUsingGET
//
// PARAMS:
//   - request: the arguments to InstanceDetailUsingGET
//
// RETURNS:
//   - InstanceDetailUsingGETResponse: The return type of the InstanceDetailUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) InstanceDetailUsingGET(request *InstanceDetailUsingGETRequest) (*InstanceDetailUsingGETResponse, error) {
	result := &InstanceDetailUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceDetailUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListRecordsUsingGET
//
// PARAMS:
//   - request: the arguments to ListRecordsUsingGET
//
// RETURNS:
//   - ListRecordsUsingGETResponse: The return type of the ListRecordsUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRecordsUsingGET(request *ListRecordsUsingGETRequest) (*ListRecordsUsingGETResponse, error) {
	result := &ListRecordsUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListRecordsUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithQueryParamFilter("listOrder", util.StringValue(request.ListOrder)).
		WithQueryParamFilter("page", util.StringValue(request.Page)).
		WithQueryParamFilter("pageSize", util.StringValue(request.PageSize)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ManualBackupUsingPOST
//
// PARAMS:
//   - request: the arguments to ManualBackupUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ManualBackupUsingPOST(request *ManualBackupUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getManualBackupUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// ModifyInstanceConfig
//
// PARAMS:
//   - request: the arguments to ModifyInstanceConfig
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyInstanceConfigUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// ModifyPasswordUsingPOST
//
// PARAMS:
//   - request: the arguments to ModifyPasswordUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyPasswordUsingPOST(request *ModifyPasswordUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getModifyPasswordUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// ModifyPublicAccess
//
// PARAMS:
//   - request: the arguments to ModifyPublicAccess
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyPublicAccess(request *ModifyPublicAccessRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyPublicAccessUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyTLSUsingPUT
//
// PARAMS:
//   - request: the arguments to ModifyTLSUsingPUT
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyTLSUsingPUT(request *ModifyTLSUsingPUTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyTLSUsingPUTUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// PasswordUsingGET
//
// PARAMS:
//   - request: the arguments to PasswordUsingGET
//
// RETURNS:
//   - PasswordUsingGETResponse: The return type of the PasswordUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PasswordUsingGET(request *PasswordUsingGETRequest) (*PasswordUsingGETResponse, error) {
	result := &PasswordUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPasswordUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("username", util.StringValue(request.Username)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RecoverInstanceUsingPOST
//
// PARAMS:
//   - request: the arguments to RecoverInstanceUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RecoverInstanceUsingPOST(request *RecoverInstanceUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRecoverInstanceUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		Do()
}

// RecoverUsingPOST
//
// PARAMS:
//   - request: the arguments to RecoverUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RecoverUsingPOST(request *RecoverUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRecoverUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// ResizeInstanceUsingPOST
//
// PARAMS:
//   - request: the arguments to ResizeInstanceUsingPOST
//
// RETURNS:
//   - ResizeInstanceUsingPOSTResponse: The return type of the ResizeInstanceUsingPOST interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeInstanceUsingPOST(request *ResizeInstanceUsingPOSTRequest) (*ResizeInstanceUsingPOSTResponse, error) {
	result := &ResizeInstanceUsingPOSTResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getResizeInstanceUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetCommentUsingPOST
//
// PARAMS:
//   - request: the arguments to SetCommentUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetCommentUsingPOST(request *SetCommentUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSetCommentUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// SetConfigUsingPOST
//
// PARAMS:
//   - request: the arguments to SetConfigUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetConfigUsingPOST(request *SetConfigUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSetConfigUsingPOSTUri(VERSION_V1)).
		WithQueryParamFilter("instanceId", util.StringValue(request.InstanceId)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// UnbindEipUsingPOST
//
// PARAMS:
//   - request: the arguments to UnbindEipUsingPOST
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UnbindEipUsingPOST(request *UnbindEipUsingPOSTRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUnbindEipUsingPOSTUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		Do()
}

// UpdateInstanceDomain
//
// PARAMS:
//   - request: the arguments to UpdateInstanceDomain
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceDomain(request *UpdateInstanceDomainRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateInstanceDomainUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// UpdateInstanceName
//
// PARAMS:
//   - request: the arguments to UpdateInstanceName
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUpdateInstanceNameUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithBody(request).
		Do()
}

// ZoneListUsingGET
//
// PARAMS:
//   - request: the arguments to ZoneListUsingGET
//
// RETURNS:
//   - ZoneListUsingGETResponse: The return type of the ZoneListUsingGET interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ZoneListUsingGET(request *ZoneListUsingGETRequest) (*ZoneListUsingGETResponse, error) {
	result := &ZoneListUsingGETResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getZoneListUsingGETUri(VERSION_V1)).
		WithQueryParamFilter("from", util.StringValue(request.From)).
		WithQueryParamFilter("engineType", util.StringValue(request.EngineType)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
