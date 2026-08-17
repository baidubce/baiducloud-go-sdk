package ocr

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// HealthReport
//
// PARAMS:
//   - request: the arguments to HealthReport
//
// RETURNS:
//   - HealthReportResponse: The return type of the HealthReport interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HealthReport(request *HealthReportRequest) (*HealthReportResponse, error) {
	result := &HealthReportResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHealthReportUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalDetail
//
// PARAMS:
//   - request: the arguments to MedicalDetail
//
// RETURNS:
//   - MedicalDetailResponse: The return type of the MedicalDetail interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalDetail(request *MedicalDetailRequest) (*MedicalDetailResponse, error) {
	result := &MedicalDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalDetailUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalInvoice
//
// PARAMS:
//   - request: the arguments to MedicalInvoice
//
// RETURNS:
//   - MedicalInvoiceResponse: The return type of the MedicalInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalInvoice(request *MedicalInvoiceRequest) (*MedicalInvoiceResponse, error) {
	result := &MedicalInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalPrescription
//
// PARAMS:
//   - request: the arguments to MedicalPrescription
//
// RETURNS:
//   - MedicalPrescriptionResponse: The return type of the MedicalPrescription interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalPrescription(request *MedicalPrescriptionRequest) (*MedicalPrescriptionResponse, error) {
	result := &MedicalPrescriptionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalPrescriptionUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalRecord
//
// PARAMS:
//   - request: the arguments to MedicalRecord
//
// RETURNS:
//   - MedicalRecordResponse: The return type of the MedicalRecord interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalRecord(request *MedicalRecordRequest) (*MedicalRecordResponse, error) {
	result := &MedicalRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalRecordUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalReportDetection
//
// PARAMS:
//   - request: the arguments to MedicalReportDetection
//
// RETURNS:
//   - MedicalReportDetectionResponse: The return type of the MedicalReportDetection interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalReportDetection(request *MedicalReportDetectionRequest) (*MedicalReportDetectionResponse, error) {
	result := &MedicalReportDetectionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalReportDetectionUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalStatement
//
// PARAMS:
//   - request: the arguments to MedicalStatement
//
// RETURNS:
//   - MedicalStatementResponse: The return type of the MedicalStatement interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalStatement(request *MedicalStatementRequest) (*MedicalStatementResponse, error) {
	result := &MedicalStatementResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalStatementUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MedicalSummary
//
// PARAMS:
//   - request: the arguments to MedicalSummary
//
// RETURNS:
//   - MedicalSummaryResponse: The return type of the MedicalSummary interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MedicalSummary(request *MedicalSummaryRequest) (*MedicalSummaryResponse, error) {
	result := &MedicalSummaryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMedicalSummaryUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
