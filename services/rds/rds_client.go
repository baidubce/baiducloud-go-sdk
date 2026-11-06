package rds

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const (
	VERSION_V1 = "v1"
)

// AddAccountApiAtTableGranularity
//
// PARAMS:
//   - request: the arguments to AddAccountApiAtTableGranularity
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) AddAccountApiAtTableGranularity(request *AddAccountApiAtTableGranularityRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddAccountApiAtTableGranularityUri(util.StringValue(request.InstanceId))).
		Do()
}

// BatchConfigurationOfInstances
//
// PARAMS:
//   - request: the arguments to BatchConfigurationOfInstances
//
// RETURNS:
//   - BatchConfigurationOfInstancesResponse: The return type of the BatchConfigurationOfInstances interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BatchConfigurationOfInstances(request *BatchConfigurationOfInstancesRequest) (*BatchConfigurationOfInstancesResponse, error) {
	result := &BatchConfigurationOfInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBatchConfigurationOfInstancesUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BindingTags
//
// PARAMS:
//   - request: the arguments to BindingTags
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) BindingTags(request *BindingTagsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBindingTagsUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// ChangeAccountPassword
//
// PARAMS:
//   - request: the arguments to ChangeAccountPassword
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ChangeAccountPassword(request *ChangeAccountPasswordRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getChangeAccountPasswordUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.AccountName))).
		WithBody(request).
		Do()
}

// CheckTheWhitelist
//
// PARAMS:
//   - request: the arguments to CheckTheWhitelist
//
// RETURNS:
//   - CheckTheWhitelistResponse: The return type of the CheckTheWhitelist interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CheckTheWhitelist(request *CheckTheWhitelistRequest) (*CheckTheWhitelistResponse, error) {
	result := &CheckTheWhitelistResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCheckTheWhitelistUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ClusterStatusCheck
//
// PARAMS:
//   - request: the arguments to ClusterStatusCheck
//
// RETURNS:
//   - ClusterStatusCheckResponse: The return type of the ClusterStatusCheck interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ClusterStatusCheck(request *ClusterStatusCheckRequest) (*ClusterStatusCheckResponse, error) {
	result := &ClusterStatusCheckResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterStatusCheckUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CopyParameterTemplate
//
// PARAMS:
//   - request: the arguments to CopyParameterTemplate
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CopyParameterTemplate(request *CopyParameterTemplateRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCopyParameterTemplateUri(VERSION_V1)).
		WithBody(request).
		Do()
}

// CreateADatabase
//
// PARAMS:
//   - request: the arguments to CreateADatabase
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateADatabase(request *CreateADatabaseRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateADatabaseUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateAPrimaryInstance
//
// PARAMS:
//   - request: the arguments to CreateAPrimaryInstance
//
// RETURNS:
//   - CreateAPrimaryInstanceResponse: The return type of the CreateAPrimaryInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAPrimaryInstance(request *CreateAPrimaryInstanceRequest) (*CreateAPrimaryInstanceResponse, error) {
	result := &CreateAPrimaryInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAPrimaryInstanceUri(VERSION_V1)).
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAProxyInstance
//
// PARAMS:
//   - request: the arguments to CreateAProxyInstance
//
// RETURNS:
//   - CreateAProxyInstanceResponse: The return type of the CreateAProxyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateAProxyInstance(request *CreateAProxyInstanceRequest) (*CreateAProxyInstanceResponse, error) {
	result := &CreateAProxyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAProxyInstanceUri(VERSION_V1)).
		WithQueryParam("rdsproxy", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateAccount
//
// PARAMS:
//   - request: the arguments to CreateAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) CreateAccount(request *CreateAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateAccountUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// CreateReadOnlyInstance
//
// PARAMS:
//   - request: the arguments to CreateReadOnlyInstance
//
// RETURNS:
//   - CreateReadOnlyInstanceResponse: The return type of the CreateReadOnlyInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CreateReadOnlyInstance(request *CreateReadOnlyInstanceRequest) (*CreateReadOnlyInstanceResponse, error) {
	result := &CreateReadOnlyInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCreateReadOnlyInstanceUri(VERSION_V1)).
		WithQueryParam("readReplica", "").
		WithQueryParamFilter("clientToken", util.StringValue(request.ClientToken)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAccount
//
// PARAMS:
//   - request: the arguments to DeleteAccount
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteAccount(request *DeleteAccountRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteAccountUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.AccountName))).
		Do()
}

// DeleteTheDatabase
//
// PARAMS:
//   - request: the arguments to DeleteTheDatabase
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTheDatabase(request *DeleteTheDatabaseRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTheDatabaseUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.DbName))).
		Do()
}

// DeleteTheSpecifiedBackupSet
//
// PARAMS:
//   - request: the arguments to DeleteTheSpecifiedBackupSet
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) DeleteTheSpecifiedBackupSet(request *DeleteTheSpecifiedBackupSetRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDeleteTheSpecifiedBackupSetUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.SnapshotId))).
		Do()
}

// DialTestInterface
//
// PARAMS:
//   - request: the arguments to DialTestInterface
//
// RETURNS:
//   - DialTestInterfaceResponse: The return type of the DialTestInterface interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DialTestInterface(request *DialTestInterfaceRequest) (*DialTestInterfaceResponse, error) {
	result := &DialTestInterfaceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDialTestInterfaceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DoesTheInstanceSupportEnablingAutomaticExpansion
//
// PARAMS:
//   - request: the arguments to DoesTheInstanceSupportEnablingAutomaticExpansion
//
// RETURNS:
//   - DoesTheInstanceSupportEnablingAutomaticExpansionResponse: The return type of the DoesTheInstanceSupportEnablingAutomaticExpansion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DoesTheInstanceSupportEnablingAutomaticExpansion(request *DoesTheInstanceSupportEnablingAutomaticExpansionRequest) (*DoesTheInstanceSupportEnablingAutomaticExpansionResponse, error) {
	result := &DoesTheInstanceSupportEnablingAutomaticExpansionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDoesTheInstanceSupportEnablingAutomaticExpansionUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ErrorLogDownloadDetails
//
// PARAMS:
//   - request: the arguments to ErrorLogDownloadDetails
//
// RETURNS:
//   - ErrorLogDownloadDetailsResponse: The return type of the ErrorLogDownloadDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ErrorLogDownloadDetails(request *ErrorLogDownloadDetailsRequest) (*ErrorLogDownloadDetailsResponse, error) {
	result := &ErrorLogDownloadDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getErrorLogDownloadDetailsUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.LogId), util.Int32Value(request.DownloadValidTimeInSec))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ErrorLogList
//
// PARAMS:
//   - request: the arguments to ErrorLogList
//
// RETURNS:
//   - ErrorLogListResponse: The return type of the ErrorLogList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ErrorLogList(request *ErrorLogListRequest) (*ErrorLogListResponse, error) {
	result := &ErrorLogListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getErrorLogListUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.Datetime))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetParameterTemplateDetails
//
// PARAMS:
//   - request: the arguments to GetParameterTemplateDetails
//
// RETURNS:
//   - GetParameterTemplateDetailsResponse: The return type of the GetParameterTemplateDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetParameterTemplateDetails(request *GetParameterTemplateDetailsRequest) (*GetParameterTemplateDetailsResponse, error) {
	result := &GetParameterTemplateDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetParameterTemplateDetailsUri(VERSION_V1, util.StringValue(request.TemplateId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetParameterTemplateList
//
// PARAMS:
//   - request: the arguments to GetParameterTemplateList
//
// RETURNS:
//   - GetParameterTemplateListResponse: The return type of the GetParameterTemplateList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetParameterTemplateList(request *GetParameterTemplateListRequest) (*GetParameterTemplateListResponse, error) {
	result := &GetParameterTemplateListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetParameterTemplateListUri(VERSION_V1)).
		WithQueryParamFilter("pageNo", "1").
		WithQueryParamFilter("pageSize", "10").
		WithQueryParamFilter("type", "system").
		WithQueryParamFilter("pageNo", util.Int32Value(request.PageNo)).
		WithQueryParamFilter("pageSize", util.Int32Value(request.PageSize)).
		WithQueryParamFilter("type", util.StringValue(request.Type)).
		WithQueryParamFilter("dbType", util.StringValue(request.DbType)).
		WithQueryParamFilter("dbVersion", util.StringValue(request.DbVersion)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPgLogDetails
//
// PARAMS:
//   - request: the arguments to GetPgLogDetails
//
// RETURNS:
//   - GetPgLogDetailsResponse: The return type of the GetPgLogDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPgLogDetails(request *GetPgLogDetailsRequest) (*GetPgLogDetailsResponse, error) {
	result := &GetPgLogDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPgLogDetailsUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("pglogId", util.StringValue(request.PglogId)).
		WithQueryParamFilter("downloadValidTimeInSec", util.Int32Value(request.DownloadValidTimeInSec)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPgLogList
//
// PARAMS:
//   - request: the arguments to GetPgLogList
//
// RETURNS:
//   - GetPgLogListResponse: The return type of the GetPgLogList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GetPgLogList(request *GetPgLogListRequest) (*GetPgLogListResponse, error) {
	result := &GetPgLogListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getGetPgLogListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("date", util.StringValue(request.Date)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// InstanceEnablesAutomaticRenewal
//
// PARAMS:
//   - request: the arguments to InstanceEnablesAutomaticRenewal
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) InstanceEnablesAutomaticRenewal(request *InstanceEnablesAutomaticRenewalRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceEnablesAutomaticRenewalUri(VERSION_V1)).
		WithQueryParam("autoRenew", "").
		WithBody(request).
		Do()
}

// ManuallyInitiateFullPhysicalBackup
//
// PARAMS:
//   - request: the arguments to ManuallyInitiateFullPhysicalBackup
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ManuallyInitiateFullPhysicalBackup(request *ManuallyInitiateFullPhysicalBackupRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getManuallyInitiateFullPhysicalBackupUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyAccountNotes
//
// PARAMS:
//   - request: the arguments to ModifyAccountNotes
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyAccountNotes(request *ModifyAccountNotesRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyAccountNotesUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.AccountName))).
		WithBody(request).
		Do()
}

// ModifyAccountPermissions
//
// PARAMS:
//   - request: the arguments to ModifyAccountPermissions
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyAccountPermissions(request *ModifyAccountPermissionsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyAccountPermissionsUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.AccountName))).
		WithBody(request).
		Do()
}

// ModifyBackupStrategy
//
// PARAMS:
//   - request: the arguments to ModifyBackupStrategy
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyBackupStrategy(request *ModifyBackupStrategyRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyBackupStrategyUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyBackupPolicy", "").
		WithBody(request).
		Do()
}

// ModifyConfigurationParameters
//
// PARAMS:
//   - request: the arguments to ModifyConfigurationParameters
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyConfigurationParameters(request *ModifyConfigurationParametersRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyConfigurationParametersUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyDatabasePort
//
// PARAMS:
//   - request: the arguments to ModifyDatabasePort
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDatabasePort(request *ModifyDatabasePortRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyDatabasePortUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// ModifyDescriptionInformation
//
// PARAMS:
//   - request: the arguments to ModifyDescriptionInformation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) ModifyDescriptionInformation(request *ModifyDescriptionInformationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyDescriptionInformationUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.DbName))).
		WithBody(request).
		Do()
}

// ObtainBinlogInformation
//
// PARAMS:
//   - request: the arguments to ObtainBinlogInformation
//
// RETURNS:
//   - ObtainBinlogInformationResponse: The return type of the ObtainBinlogInformation interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ObtainBinlogInformation(request *ObtainBinlogInformationRequest) (*ObtainBinlogInformationResponse, error) {
	result := &ObtainBinlogInformationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getObtainBinlogInformationUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.LogId), util.StringValue(request.DownloadValidTimeInSec))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ObtainSslEncryptionCertificate
//
// PARAMS:
//   - request: the arguments to ObtainSslEncryptionCertificate
//
// RETURNS:
//   - ObtainSslEncryptionCertificateResponse: The return type of the ObtainSslEncryptionCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ObtainSslEncryptionCertificate() (*ObtainSslEncryptionCertificateResponse, error) {
	result := &ObtainSslEncryptionCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getObtainSslEncryptionCertificateUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ObtainSslEncryptionInformation
//
// PARAMS:
//   - request: the arguments to ObtainSslEncryptionInformation
//
// RETURNS:
//   - ObtainSslEncryptionInformationResponse: The return type of the ObtainSslEncryptionInformation interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ObtainSslEncryptionInformation(request *ObtainSslEncryptionInformationRequest) (*ObtainSslEncryptionInformationResponse, error) {
	result := &ObtainSslEncryptionInformationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getObtainSslEncryptionInformationUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAccountDetails
//
// PARAMS:
//   - request: the arguments to QueryAccountDetails
//
// RETURNS:
//   - QueryAccountDetailsResponse: The return type of the QueryAccountDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAccountDetails(request *QueryAccountDetailsRequest) (*QueryAccountDetailsResponse, error) {
	result := &QueryAccountDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAccountDetailsUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.AccountName))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAccountList
//
// PARAMS:
//   - request: the arguments to QueryAccountList
//
// RETURNS:
//   - QueryAccountListResponse: The return type of the QueryAccountList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAccountList(request *QueryAccountListRequest) (*QueryAccountListResponse, error) {
	result := &QueryAccountListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAccountListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryAndModifyHistory
//
// PARAMS:
//   - request: the arguments to QueryAndModifyHistory
//
// RETURNS:
//   - QueryAndModifyHistoryResponse: The return type of the QueryAndModifyHistory interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryAndModifyHistory(request *QueryAndModifyHistoryRequest) (*QueryAndModifyHistoryResponse, error) {
	result := &QueryAndModifyHistoryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryAndModifyHistoryUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryBackupDetails
//
// PARAMS:
//   - request: the arguments to QueryBackupDetails
//
// RETURNS:
//   - QueryBackupDetailsResponse: The return type of the QueryBackupDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryBackupDetails(request *QueryBackupDetailsRequest) (*QueryBackupDetailsResponse, error) {
	result := &QueryBackupDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryBackupDetailsUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.BackupId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryBackupList
//
// PARAMS:
//   - request: the arguments to QueryBackupList
//
// RETURNS:
//   - QueryBackupListResponse: The return type of the QueryBackupList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryBackupList(request *QueryBackupListRequest) (*QueryBackupListResponse, error) {
	result := &QueryBackupListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryBackupListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryInstanceList
//
// PARAMS:
//   - request: the arguments to QueryInstanceList
//
// RETURNS:
//   - QueryInstanceListResponse: The return type of the QueryInstanceList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryInstanceList(request *QueryInstanceListRequest) (*QueryInstanceListResponse, error) {
	result := &QueryInstanceListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryInstanceListUri(VERSION_V1)).
		WithQueryParamFilter("marker", util.StringValue(request.Marker)).
		WithQueryParamFilter("maxKeys", util.Int32Value(request.MaxKeys)).
		WithQueryParamFilter("vnetIp", util.StringValue(request.VnetIp)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryParameterList
//
// PARAMS:
//   - request: the arguments to QueryParameterList
//
// RETURNS:
//   - QueryParameterListResponse: The return type of the QueryParameterList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryParameterList(request *QueryParameterListRequest) (*QueryParameterListResponse, error) {
	result := &QueryParameterListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryParameterListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("keyword", "keyword").
		WithQueryParamFilter("keyword", util.StringValue(request.Keyword)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuerySubnetList
//
// PARAMS:
//   - request: the arguments to QuerySubnetList
//
// RETURNS:
//   - QuerySubnetListResponse: The return type of the QuerySubnetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuerySubnetList(request *QuerySubnetListRequest) (*QuerySubnetListResponse, error) {
	result := &QuerySubnetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuerySubnetListUri(VERSION_V1)).
		WithQueryParamFilter("zoneName", util.StringValue(request.ZoneName)).
		WithQueryParamFilter("vpcId", util.StringValue(request.VpcId)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTaskList
//
// PARAMS:
//   - request: the arguments to QueryTaskList
//
// RETURNS:
//   - QueryTaskListResponse: The return type of the QueryTaskList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTaskList(request *QueryTaskListRequest) (*QueryTaskListResponse, error) {
	result := &QueryTaskListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQueryTaskListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheDetailsOfTheSpecifiedInstance
//
// PARAMS:
//   - request: the arguments to QueryTheDetailsOfTheSpecifiedInstance
//
// RETURNS:
//   - QueryTheDetailsOfTheSpecifiedInstanceResponse: The return type of the QueryTheDetailsOfTheSpecifiedInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheDetailsOfTheSpecifiedInstance(request *QueryTheDetailsOfTheSpecifiedInstanceRequest) (*QueryTheDetailsOfTheSpecifiedInstanceResponse, error) {
	result := &QueryTheDetailsOfTheSpecifiedInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheDetailsOfTheSpecifiedInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QueryTheListOfAvailableZones
//
// PARAMS:
//   - request: the arguments to QueryTheListOfAvailableZones
//
// RETURNS:
//   - QueryTheListOfAvailableZonesResponse: The return type of the QueryTheListOfAvailableZones interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QueryTheListOfAvailableZones() (*QueryTheListOfAvailableZonesResponse, error) {
	result := &QueryTheListOfAvailableZonesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryTheListOfAvailableZonesUri(VERSION_V1)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReleaseInstance
//
// PARAMS:
//   - request: the arguments to ReleaseInstance
//
// RETURNS:
//   - ReleaseInstanceResponse: The return type of the ReleaseInstance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ReleaseInstance(request *ReleaseInstanceRequest) (*ReleaseInstanceResponse, error) {
	result := &ReleaseInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getReleaseInstanceUri(VERSION_V1)).
		WithQueryParamFilter("instanceIds", util.StringValue(request.InstanceIds)).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RenewalExample
//
// PARAMS:
//   - request: the arguments to RenewalExample
//
// RETURNS:
//   - RenewalExampleResponse: The return type of the RenewalExample interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RenewalExample(request *RenewalExampleRequest) (*RenewalExampleResponse, error) {
	result := &RenewalExampleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRenewalExampleUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RestartTheInstance
//
// PARAMS:
//   - request: the arguments to RestartTheInstance
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestartTheInstance(request *RestartTheInstanceRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestartTheInstanceUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("reboot", "").
		WithBody(request).
		Do()
}

// RestoreLibraryTablesBasedOnBackupSets
//
// PARAMS:
//   - request: the arguments to RestoreLibraryTablesBasedOnBackupSets
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestoreLibraryTablesBasedOnBackupSets(request *RestoreLibraryTablesBasedOnBackupSetsRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestoreLibraryTablesBasedOnBackupSetsUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// RestoreTheDatabaseTableByTimePoint
//
// PARAMS:
//   - request: the arguments to RestoreTheDatabaseTableByTimePoint
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) RestoreTheDatabaseTableByTimePoint(request *RestoreTheDatabaseTableByTimePointRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRestoreTheDatabaseTableByTimePointUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// RetrieveDatabaseList
//
// PARAMS:
//   - request: the arguments to RetrieveDatabaseList
//
// RETURNS:
//   - RetrieveDatabaseListResponse: The return type of the RetrieveDatabaseList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveDatabaseList(request *RetrieveDatabaseListRequest) (*RetrieveDatabaseListResponse, error) {
	result := &RetrieveDatabaseListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveDatabaseListUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveInstanceAutoScalingConfig
//
// PARAMS:
//   - request: the arguments to RetrieveInstanceAutoScalingConfig
//
// RETURNS:
//   - RetrieveInstanceAutoScalingConfigResponse: The return type of the RetrieveInstanceAutoScalingConfig interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveInstanceAutoScalingConfig(request *RetrieveInstanceAutoScalingConfigRequest) (*RetrieveInstanceAutoScalingConfigResponse, error) {
	result := &RetrieveInstanceAutoScalingConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveInstanceAutoScalingConfigUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveTheBinlogList
//
// PARAMS:
//   - request: the arguments to RetrieveTheBinlogList
//
// RETURNS:
//   - RetrieveTheBinlogListResponse: The return type of the RetrieveTheBinlogList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RetrieveTheBinlogList(request *RetrieveTheBinlogListRequest) (*RetrieveTheBinlogListResponse, error) {
	result := &RetrieveTheBinlogListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRetrieveTheBinlogListUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.Datetime))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetEncryptionPolicyApi
//
// PARAMS:
//   - request: the arguments to SetEncryptionPolicyApi
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) SetEncryptionPolicyApi(request *SetEncryptionPolicyApiRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSetEncryptionPolicyApiUri(util.StringValue(request.InstanceId))).
		WithQueryParam("updateEncryptPolicy", "").
		WithBody(request).
		Do()
}

// SlowLogDownloadDetails
//
// PARAMS:
//   - request: the arguments to SlowLogDownloadDetails
//
// RETURNS:
//   - SlowLogDownloadDetailsResponse: The return type of the SlowLogDownloadDetails interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SlowLogDownloadDetails(request *SlowLogDownloadDetailsRequest) (*SlowLogDownloadDetailsResponse, error) {
	result := &SlowLogDownloadDetailsResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSlowLogDownloadDetailsUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.LogId), util.Int32Value(request.DownloadValidTimeInSec))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SlowLogList
//
// PARAMS:
//   - request: the arguments to SlowLogList
//
// RETURNS:
//   - SlowLogListResponse: The return type of the SlowLogList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SlowLogList(request *SlowLogListRequest) (*SlowLogListResponse, error) {
	result := &SlowLogListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSlowLogListUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.Datetime))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TableGranularityQueryUserDetailsApi
//
// PARAMS:
//   - request: the arguments to TableGranularityQueryUserDetailsApi
//
// RETURNS:
//   - TableGranularityQueryUserDetailsApiResponse: The return type of the TableGranularityQueryUserDetailsApi interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TableGranularityQueryUserDetailsApi(request *TableGranularityQueryUserDetailsApiRequest) (*TableGranularityQueryUserDetailsApiResponse, error) {
	result := &TableGranularityQueryUserDetailsApiResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTableGranularityQueryUserDetailsApiUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TableGranularityQueryUserListApi
//
// PARAMS:
//   - request: the arguments to TableGranularityQueryUserListApi
//
// RETURNS:
//   - TableGranularityQueryUserListApiResponse: The return type of the TableGranularityQueryUserListApi interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TableGranularityQueryUserListApi(request *TableGranularityQueryUserListApiRequest) (*TableGranularityQueryUserListApiResponse, error) {
	result := &TableGranularityQueryUserListApiResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTableGranularityQueryUserListApiUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TableGranularityUpdateUserDescriptionApi
//
// PARAMS:
//   - request: the arguments to TableGranularityUpdateUserDescriptionApi
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) TableGranularityUpdateUserDescriptionApi(request *TableGranularityUpdateUserDescriptionApiRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTableGranularityUpdateUserDescriptionApiUri()).
		WithBody(request).
		Do()
}

// TargetRegionMappingListApi
//
// PARAMS:
//   - request: the arguments to TargetRegionMappingListApi
//
// RETURNS:
//   - TargetRegionMappingListApiResponse: The return type of the TargetRegionMappingListApi interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TargetRegionMappingListApi(request *TargetRegionMappingListApiRequest) (*TargetRegionMappingListApiResponse, error) {
	result := &TargetRegionMappingListApiResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTargetRegionMappingListApiUri(util.StringValue(request.InstanceId))).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TransformationExample
//
// PARAMS:
//   - request: the arguments to TransformationExample
//
// RETURNS:
//   - TransformationExampleResponse: The return type of the TransformationExample interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TransformationExample(request *TransformationExampleRequest) (*TransformationExampleResponse, error) {
	result := &TransformationExampleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTransformationExampleUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("resize", "").
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutomaticStorageExpansionConfiguration
//
// PARAMS:
//   - request: the arguments to UpdateAutomaticStorageExpansionConfiguration
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateAutomaticStorageExpansionConfiguration(request *UpdateAutomaticStorageExpansionConfigurationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateAutomaticStorageExpansionConfigurationUri(VERSION_V1, util.StringValue(request.InstanceId), util.StringValue(request.Action))).
		WithBody(request).
		Do()
}

// UpdateInstanceConnectionInformation
//
// PARAMS:
//   - request: the arguments to UpdateInstanceConnectionInformation
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceConnectionInformation(request *UpdateInstanceConnectionInformationRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceConnectionInformationUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyEndpoint", "").
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
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceNameUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("rename", "").
		WithBody(request).
		Do()
}

// UpdateInstancePublicNetworkAccessStatus
//
// PARAMS:
//   - request: the arguments to UpdateInstancePublicNetworkAccessStatus
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstancePublicNetworkAccessStatus(request *UpdateInstancePublicNetworkAccessStatusRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstancePublicNetworkAccessStatusUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifyPublicAccess", "").
		WithBody(request).
		Do()
}

// UpdateInstanceSynchronizationMode
//
// PARAMS:
//   - request: the arguments to UpdateInstanceSynchronizationMode
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceSynchronizationMode(request *UpdateInstanceSynchronizationModeRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateInstanceSynchronizationModeUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParam("modifySyncMode", "").
		WithBody(request).
		Do()
}

// UpdateTimeWindow
//
// PARAMS:
//   - request: the arguments to UpdateTimeWindow
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateTimeWindow(request *UpdateTimeWindowRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateTimeWindowUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithBody(request).
		Do()
}

// UpdateWhitelist
//
// PARAMS:
//   - request: the arguments to UpdateWhitelist
//
// RETURNS:

// - error: nil if success otherwise the specific error
func (c *Client) UpdateWhitelist(request *UpdateWhitelistRequest) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUpdateWhitelistUri(VERSION_V1, util.StringValue(request.InstanceId))).
		WithQueryParamFilter("action", util.StringValue(request.Action)).
		WithBody(request).
		Do()
}
