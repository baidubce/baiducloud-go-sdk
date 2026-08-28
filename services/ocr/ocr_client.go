package ocr

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// AccountOpening
//
// PARAMS:
//   - request: the arguments to AccountOpening
//
// RETURNS:
//   - AccountOpeningResponse: The return type of the AccountOpening interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AccountOpening(request *AccountOpeningRequest) (*AccountOpeningResponse, error) {
	result := &AccountOpeningResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAccountOpeningUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Accurate
//
// PARAMS:
//   - request: the arguments to Accurate
//
// RETURNS:
//   - AccurateResponse: The return type of the Accurate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Accurate(request *AccurateRequest) (*AccurateResponse, error) {
	result := &AccurateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAccurateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AccurateBasic
//
// PARAMS:
//   - request: the arguments to AccurateBasic
//
// RETURNS:
//   - AccurateBasicResponse: The return type of the AccurateBasic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AccurateBasic(request *AccurateBasicRequest) (*AccurateBasicResponse, error) {
	result := &AccurateBasicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAccurateBasicUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AirTicket
//
// PARAMS:
//   - request: the arguments to AirTicket
//
// RETURNS:
//   - AirTicketResponse: The return type of the AirTicket interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AirTicket(request *AirTicketRequest) (*AirTicketResponse, error) {
	result := &AirTicketResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAirTicketUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BankReceiptNew
//
// PARAMS:
//   - request: the arguments to BankReceiptNew
//
// RETURNS:
//   - BankReceiptNewResponse: The return type of the BankReceiptNew interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BankReceiptNew(request *BankReceiptNewRequest) (*BankReceiptNewResponse, error) {
	result := &BankReceiptNewResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBankReceiptNewUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Bankcard
//
// PARAMS:
//   - request: the arguments to Bankcard
//
// RETURNS:
//   - BankcardResponse: The return type of the Bankcard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Bankcard(request *BankcardRequest) (*BankcardResponse, error) {
	result := &BankcardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBankcardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BirthCertificate
//
// PARAMS:
//   - request: the arguments to BirthCertificate
//
// RETURNS:
//   - BirthCertificateResponse: The return type of the BirthCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BirthCertificate(request *BirthCertificateRequest) (*BirthCertificateResponse, error) {
	result := &BirthCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBirthCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusTicket
//
// PARAMS:
//   - request: the arguments to BusTicket
//
// RETURNS:
//   - BusTicketResponse: The return type of the BusTicket interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusTicket(request *BusTicketRequest) (*BusTicketResponse, error) {
	result := &BusTicketResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusTicketUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusinessLicense
//
// PARAMS:
//   - request: the arguments to BusinessLicense
//
// RETURNS:
//   - BusinessLicenseResponse: The return type of the BusinessLicense interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusinessLicense(request *BusinessLicenseRequest) (*BusinessLicenseResponse, error) {
	result := &BusinessLicenseResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusinessLicenseUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusinesslicenseDetailed
//
// PARAMS:
//   - request: the arguments to BusinesslicenseDetailed
//
// RETURNS:
//   - BusinesslicenseDetailedResponse: The return type of the BusinesslicenseDetailed interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusinesslicenseDetailed(request *BusinesslicenseDetailedRequest) (*BusinesslicenseDetailedResponse, error) {
	result := &BusinesslicenseDetailedResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusinesslicenseDetailedUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusinesslicenseStandard
//
// PARAMS:
//   - request: the arguments to BusinesslicenseStandard
//
// RETURNS:
//   - BusinesslicenseStandardResponse: The return type of the BusinesslicenseStandard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusinesslicenseStandard(request *BusinesslicenseStandardRequest) (*BusinesslicenseStandardResponse, error) {
	result := &BusinesslicenseStandardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusinesslicenseStandardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusinesslicenseVerificationDetailed
//
// PARAMS:
//   - request: the arguments to BusinesslicenseVerificationDetailed
//
// RETURNS:
//   - BusinesslicenseVerificationDetailedResponse: The return type of the BusinesslicenseVerificationDetailed interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusinesslicenseVerificationDetailed(request *BusinesslicenseVerificationDetailedRequest) (*BusinesslicenseVerificationDetailedResponse, error) {
	result := &BusinesslicenseVerificationDetailedResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusinesslicenseVerificationDetailedUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BusinesslicenseVerificationStandard
//
// PARAMS:
//   - request: the arguments to BusinesslicenseVerificationStandard
//
// RETURNS:
//   - BusinesslicenseVerificationStandardResponse: The return type of the BusinesslicenseVerificationStandard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) BusinesslicenseVerificationStandard(request *BusinesslicenseVerificationStandardRequest) (*BusinesslicenseVerificationStandardResponse, error) {
	result := &BusinesslicenseVerificationStandardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBusinesslicenseVerificationStandardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CorrectEduCreateTask
//
// PARAMS:
//   - request: the arguments to CorrectEduCreateTask
//
// RETURNS:
//   - CorrectEduCreateTaskResponse: The return type of the CorrectEduCreateTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CorrectEduCreateTask(request *CorrectEduCreateTaskRequest) (*CorrectEduCreateTaskResponse, error) {
	result := &CorrectEduCreateTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCorrectEduCreateTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CorrectEduGetResult
//
// PARAMS:
//   - request: the arguments to CorrectEduGetResult
//
// RETURNS:
//   - CorrectEduGetResultResponse: The return type of the CorrectEduGetResult interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CorrectEduGetResult(request *CorrectEduGetResultRequest) (*CorrectEduGetResultResponse, error) {
	result := &CorrectEduGetResultResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCorrectEduGetResultUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DivorceCertificate
//
// PARAMS:
//   - request: the arguments to DivorceCertificate
//
// RETURNS:
//   - DivorceCertificateResponse: The return type of the DivorceCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DivorceCertificate(request *DivorceCertificateRequest) (*DivorceCertificateResponse, error) {
	result := &DivorceCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDivorceCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DocAnalysis
//
// PARAMS:
//   - request: the arguments to DocAnalysis
//
// RETURNS:
//   - DocAnalysisResponse: The return type of the DocAnalysis interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DocAnalysis(request *DocAnalysisRequest) (*DocAnalysisResponse, error) {
	result := &DocAnalysisResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDocAnalysisUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DocAnalysisOffice
//
// PARAMS:
//   - request: the arguments to DocAnalysisOffice
//
// RETURNS:
//   - DocAnalysisOfficeResponse: The return type of the DocAnalysisOffice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DocAnalysisOffice(request *DocAnalysisOfficeRequest) (*DocAnalysisOfficeResponse, error) {
	result := &DocAnalysisOfficeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDocAnalysisOfficeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DocClassify
//
// PARAMS:
//   - request: the arguments to DocClassify
//
// RETURNS:
//   - DocClassifyResponse: The return type of the DocClassify interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DocClassify(request *DocClassifyRequest) (*DocClassifyResponse, error) {
	result := &DocClassifyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDocClassifyUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DocCropEnhance
//
// PARAMS:
//   - request: the arguments to DocCropEnhance
//
// RETURNS:
//   - DocCropEnhanceResponse: The return type of the DocCropEnhance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DocCropEnhance(request *DocCropEnhanceRequest) (*DocCropEnhanceResponse, error) {
	result := &DocCropEnhanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDocCropEnhanceUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DrivingLicense
//
// PARAMS:
//   - request: the arguments to DrivingLicense
//
// RETURNS:
//   - DrivingLicenseResponse: The return type of the DrivingLicense interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DrivingLicense(request *DrivingLicenseRequest) (*DrivingLicenseResponse, error) {
	result := &DrivingLicenseResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDrivingLicenseUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Facade
//
// PARAMS:
//   - request: the arguments to Facade
//
// RETURNS:
//   - FacadeResponse: The return type of the Facade interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Facade(request *FacadeRequest) (*FacadeResponse, error) {
	result := &FacadeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFacadeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FerryTicket
//
// PARAMS:
//   - request: the arguments to FerryTicket
//
// RETURNS:
//   - FerryTicketResponse: The return type of the FerryTicket interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FerryTicket(request *FerryTicketRequest) (*FerryTicketResponse, error) {
	result := &FerryTicketResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFerryTicketUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ForeignResidentIdCard
//
// PARAMS:
//   - request: the arguments to ForeignResidentIdCard
//
// RETURNS:
//   - ForeignResidentIdCardResponse: The return type of the ForeignResidentIdCard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ForeignResidentIdCard(request *ForeignResidentIdCardRequest) (*ForeignResidentIdCardResponse, error) {
	result := &ForeignResidentIdCardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getForeignResidentIdCardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ForgeryDetection
//
// PARAMS:
//   - request: the arguments to ForgeryDetection
//
// RETURNS:
//   - ForgeryDetectionResponse: The return type of the ForgeryDetection interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ForgeryDetection(request *ForgeryDetectionRequest) (*ForgeryDetectionResponse, error) {
	result := &ForgeryDetectionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getForgeryDetectionUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FourFactorsVerification
//
// PARAMS:
//   - request: the arguments to FourFactorsVerification
//
// RETURNS:
//   - FourFactorsVerificationResponse: The return type of the FourFactorsVerification interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FourFactorsVerification(request *FourFactorsVerificationRequest) (*FourFactorsVerificationResponse, error) {
	result := &FourFactorsVerificationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFourFactorsVerificationUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// General
//
// PARAMS:
//   - request: the arguments to General
//
// RETURNS:
//   - GeneralResponse: The return type of the General interface.
//   - error: nil if success otherwise the specific error
func (c *Client) General(request *GeneralRequest) (*GeneralResponse, error) {
	result := &GeneralResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGeneralUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GeneralBasic
//
// PARAMS:
//   - request: the arguments to GeneralBasic
//
// RETURNS:
//   - GeneralBasicResponse: The return type of the GeneralBasic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GeneralBasic(request *GeneralBasicRequest) (*GeneralBasicResponse, error) {
	result := &GeneralBasicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGeneralBasicUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Handwriting
//
// PARAMS:
//   - request: the arguments to Handwriting
//
// RETURNS:
//   - HandwritingResponse: The return type of the Handwriting interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Handwriting(request *HandwritingRequest) (*HandwritingResponse, error) {
	result := &HandwritingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHandwritingUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HandwritingCompositionCreateTask
//
// PARAMS:
//   - request: the arguments to HandwritingCompositionCreateTask
//
// RETURNS:
//   - HandwritingCompositionCreateTaskResponse: The return type of the HandwritingCompositionCreateTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HandwritingCompositionCreateTask(request *HandwritingCompositionCreateTaskRequest) (*HandwritingCompositionCreateTaskResponse, error) {
	result := &HandwritingCompositionCreateTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHandwritingCompositionCreateTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HandwritingCompositionGetResult
//
// PARAMS:
//   - request: the arguments to HandwritingCompositionGetResult
//
// RETURNS:
//   - HandwritingCompositionGetResultResponse: The return type of the HandwritingCompositionGetResult interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HandwritingCompositionGetResult(request *HandwritingCompositionGetResultRequest) (*HandwritingCompositionGetResultResponse, error) {
	result := &HandwritingCompositionGetResultResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHandwritingCompositionGetResultUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

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

// HkMacauTaiwanExitentrypermit
//
// PARAMS:
//   - request: the arguments to HkMacauTaiwanExitentrypermit
//
// RETURNS:
//   - HkMacauTaiwanExitentrypermitResponse: The return type of the HkMacauTaiwanExitentrypermit interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HkMacauTaiwanExitentrypermit(request *HkMacauTaiwanExitentrypermitRequest) (*HkMacauTaiwanExitentrypermitResponse, error) {
	result := &HkMacauTaiwanExitentrypermitResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHkMacauTaiwanExitentrypermitUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HkMacauTaiwanpermit
//
// PARAMS:
//   - request: the arguments to HkMacauTaiwanpermit
//
// RETURNS:
//   - HkMacauTaiwanpermitResponse: The return type of the HkMacauTaiwanpermit interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HkMacauTaiwanpermit(request *HkMacauTaiwanpermitRequest) (*HkMacauTaiwanpermitResponse, error) {
	result := &HkMacauTaiwanpermitResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHkMacauTaiwanpermitUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// HouseholdRegister
//
// PARAMS:
//   - request: the arguments to HouseholdRegister
//
// RETURNS:
//   - HouseholdRegisterResponse: The return type of the HouseholdRegister interface.
//   - error: nil if success otherwise the specific error
func (c *Client) HouseholdRegister(request *HouseholdRegisterRequest) (*HouseholdRegisterResponse, error) {
	result := &HouseholdRegisterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHouseholdRegisterUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Idcard
//
// PARAMS:
//   - request: the arguments to Idcard
//
// RETURNS:
//   - IdcardResponse: The return type of the Idcard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Idcard(request *IdcardRequest) (*IdcardResponse, error) {
	result := &IdcardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getIdcardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Invoice
//
// PARAMS:
//   - request: the arguments to Invoice
//
// RETURNS:
//   - InvoiceResponse: The return type of the Invoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Invoice(request *InvoiceRequest) (*InvoiceResponse, error) {
	result := &InvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LicensePlate
//
// PARAMS:
//   - request: the arguments to LicensePlate
//
// RETURNS:
//   - LicensePlateResponse: The return type of the LicensePlate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LicensePlate(request *LicensePlateRequest) (*LicensePlateResponse, error) {
	result := &LicensePlateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLicensePlateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MarriageCertificate
//
// PARAMS:
//   - request: the arguments to MarriageCertificate
//
// RETURNS:
//   - MarriageCertificateResponse: The return type of the MarriageCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MarriageCertificate(request *MarriageCertificateRequest) (*MarriageCertificateResponse, error) {
	result := &MarriageCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMarriageCertificateUri()).
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

// Meter
//
// PARAMS:
//   - request: the arguments to Meter
//
// RETURNS:
//   - MeterResponse: The return type of the Meter interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Meter(request *MeterRequest) (*MeterResponse, error) {
	result := &MeterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMeterUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MixedMultiVehicle
//
// PARAMS:
//   - request: the arguments to MixedMultiVehicle
//
// RETURNS:
//   - MixedMultiVehicleResponse: The return type of the MixedMultiVehicle interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MixedMultiVehicle(request *MixedMultiVehicleRequest) (*MixedMultiVehicleResponse, error) {
	result := &MixedMultiVehicleResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMixedMultiVehicleUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MultiIdcard
//
// PARAMS:
//   - request: the arguments to MultiIdcard
//
// RETURNS:
//   - MultiIdcardResponse: The return type of the MultiIdcard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MultiIdcard(request *MultiIdcardRequest) (*MultiIdcardResponse, error) {
	result := &MultiIdcardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMultiIdcardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MultipleInvoice
//
// PARAMS:
//   - request: the arguments to MultipleInvoice
//
// RETURNS:
//   - MultipleInvoiceResponse: The return type of the MultipleInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MultipleInvoice(request *MultipleInvoiceRequest) (*MultipleInvoiceResponse, error) {
	result := &MultipleInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMultipleInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Numbers
//
// PARAMS:
//   - request: the arguments to Numbers
//
// RETURNS:
//   - NumbersResponse: The return type of the Numbers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Numbers(request *NumbersRequest) (*NumbersResponse, error) {
	result := &NumbersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNumbersUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OnlineTaxiItinerary
//
// PARAMS:
//   - request: the arguments to OnlineTaxiItinerary
//
// RETURNS:
//   - OnlineTaxiItineraryResponse: The return type of the OnlineTaxiItinerary interface.
//   - error: nil if success otherwise the specific error
func (c *Client) OnlineTaxiItinerary(request *OnlineTaxiItineraryRequest) (*OnlineTaxiItineraryResponse, error) {
	result := &OnlineTaxiItineraryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getOnlineTaxiItineraryUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OverseasPassport
//
// PARAMS:
//   - request: the arguments to OverseasPassport
//
// RETURNS:
//   - OverseasPassportResponse: The return type of the OverseasPassport interface.
//   - error: nil if success otherwise the specific error
func (c *Client) OverseasPassport(request *OverseasPassportRequest) (*OverseasPassportResponse, error) {
	result := &OverseasPassportResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getOverseasPassportUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaddleVlParserTask
//
// PARAMS:
//   - request: the arguments to PaddleVlParserTask
//
// RETURNS:
//   - PaddleVlParserTaskResponse: The return type of the PaddleVlParserTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PaddleVlParserTask(request *PaddleVlParserTaskRequest) (*PaddleVlParserTaskResponse, error) {
	result := &PaddleVlParserTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPaddleVlParserTaskUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaddleVlParserTaskQuery
//
// PARAMS:
//   - request: the arguments to PaddleVlParserTaskQuery
//
// RETURNS:
//   - PaddleVlParserTaskQueryResponse: The return type of the PaddleVlParserTaskQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PaddleVlParserTaskQuery(request *PaddleVlParserTaskQueryRequest) (*PaddleVlParserTaskQueryResponse, error) {
	result := &PaddleVlParserTaskQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPaddleVlParserTaskQueryUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaperCutEdu
//
// PARAMS:
//   - request: the arguments to PaperCutEdu
//
// RETURNS:
//   - PaperCutEduResponse: The return type of the PaperCutEdu interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PaperCutEdu(request *PaperCutEduRequest) (*PaperCutEduResponse, error) {
	result := &PaperCutEduResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPaperCutEduUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaperCutEduVlmCreateTask
//
// PARAMS:
//   - request: the arguments to PaperCutEduVlmCreateTask
//
// RETURNS:
//   - PaperCutEduVlmCreateTaskResponse: The return type of the PaperCutEduVlmCreateTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PaperCutEduVlmCreateTask(request *PaperCutEduVlmCreateTaskRequest) (*PaperCutEduVlmCreateTaskResponse, error) {
	result := &PaperCutEduVlmCreateTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPaperCutEduVlmCreateTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PaperCutEduVlmGetResult
//
// PARAMS:
//   - request: the arguments to PaperCutEduVlmGetResult
//
// RETURNS:
//   - PaperCutEduVlmGetResultResponse: The return type of the PaperCutEduVlmGetResult interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PaperCutEduVlmGetResult(request *PaperCutEduVlmGetResultRequest) (*PaperCutEduVlmGetResultResponse, error) {
	result := &PaperCutEduVlmGetResultResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPaperCutEduVlmGetResultUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParserTask
//
// PARAMS:
//   - request: the arguments to ParserTask
//
// RETURNS:
//   - ParserTaskResponse: The return type of the ParserTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ParserTask(request *ParserTaskRequest) (*ParserTaskResponse, error) {
	result := &ParserTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getParserTaskUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParserTaskQuery
//
// PARAMS:
//   - request: the arguments to ParserTaskQuery
//
// RETURNS:
//   - ParserTaskQueryResponse: The return type of the ParserTaskQuery interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ParserTaskQuery(request *ParserTaskQueryRequest) (*ParserTaskQueryResponse, error) {
	result := &ParserTaskQueryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getParserTaskQueryUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Passport
//
// PARAMS:
//   - request: the arguments to Passport
//
// RETURNS:
//   - PassportResponse: The return type of the Passport interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Passport(request *PassportRequest) (*PassportResponse, error) {
	result := &PassportResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPassportUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Qrcode
//
// PARAMS:
//   - request: the arguments to Qrcode
//
// RETURNS:
//   - QrcodeResponse: The return type of the Qrcode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Qrcode(request *QrcodeRequest) (*QrcodeResponse, error) {
	result := &QrcodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQrcodeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// QuotaInvoice
//
// PARAMS:
//   - request: the arguments to QuotaInvoice
//
// RETURNS:
//   - QuotaInvoiceResponse: The return type of the QuotaInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) QuotaInvoice(request *QuotaInvoiceRequest) (*QuotaInvoiceResponse, error) {
	result := &QuotaInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getQuotaInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RealEstateCertificate
//
// PARAMS:
//   - request: the arguments to RealEstateCertificate
//
// RETURNS:
//   - RealEstateCertificateResponse: The return type of the RealEstateCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RealEstateCertificate(request *RealEstateCertificateRequest) (*RealEstateCertificateResponse, error) {
	result := &RealEstateCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRealEstateCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveHandwriting
//
// PARAMS:
//   - request: the arguments to RemoveHandwriting
//
// RETURNS:
//   - RemoveHandwritingResponse: The return type of the RemoveHandwriting interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoveHandwriting(request *RemoveHandwritingRequest) (*RemoveHandwritingResponse, error) {
	result := &RemoveHandwritingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveHandwritingUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RoadTransportCertificate
//
// PARAMS:
//   - request: the arguments to RoadTransportCertificate
//
// RETURNS:
//   - RoadTransportCertificateResponse: The return type of the RoadTransportCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RoadTransportCertificate(request *RoadTransportCertificateRequest) (*RoadTransportCertificateResponse, error) {
	result := &RoadTransportCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRoadTransportCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Seal
//
// PARAMS:
//   - request: the arguments to Seal
//
// RETURNS:
//   - SealResponse: The return type of the Seal interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Seal(request *SealRequest) (*SealResponse, error) {
	result := &SealResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSealUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ShoppingReceipt
//
// PARAMS:
//   - request: the arguments to ShoppingReceipt
//
// RETURNS:
//   - ShoppingReceiptResponse: The return type of the ShoppingReceipt interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ShoppingReceipt(request *ShoppingReceiptRequest) (*ShoppingReceiptResponse, error) {
	result := &ShoppingReceiptResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getShoppingReceiptUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SmartStruct
//
// PARAMS:
//   - request: the arguments to SmartStruct
//
// RETURNS:
//   - SmartStructResponse: The return type of the SmartStruct interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SmartStruct(request *SmartStructRequest) (*SmartStructResponse, error) {
	result := &SmartStructResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSmartStructUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SocialSecurityCard
//
// PARAMS:
//   - request: the arguments to SocialSecurityCard
//
// RETURNS:
//   - SocialSecurityCardResponse: The return type of the SocialSecurityCard interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SocialSecurityCard(request *SocialSecurityCardRequest) (*SocialSecurityCardResponse, error) {
	result := &SocialSecurityCardResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSocialSecurityCardUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Table
//
// PARAMS:
//   - request: the arguments to Table
//
// RETURNS:
//   - TableResponse: The return type of the Table interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Table(request *TableRequest) (*TableResponse, error) {
	result := &TableResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTableUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TaxiReceipt
//
// PARAMS:
//   - request: the arguments to TaxiReceipt
//
// RETURNS:
//   - TaxiReceiptResponse: The return type of the TaxiReceipt interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TaxiReceipt(request *TaxiReceiptRequest) (*TaxiReceiptResponse, error) {
	result := &TaxiReceiptResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTaxiReceiptUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ThreeFactorsVerification
//
// PARAMS:
//   - request: the arguments to ThreeFactorsVerification
//
// RETURNS:
//   - ThreeFactorsVerificationResponse: The return type of the ThreeFactorsVerification interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ThreeFactorsVerification(request *ThreeFactorsVerificationRequest) (*ThreeFactorsVerificationResponse, error) {
	result := &ThreeFactorsVerificationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getThreeFactorsVerificationUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TollInvoice
//
// PARAMS:
//   - request: the arguments to TollInvoice
//
// RETURNS:
//   - TollInvoiceResponse: The return type of the TollInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TollInvoice(request *TollInvoiceRequest) (*TollInvoiceResponse, error) {
	result := &TollInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTollInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TrainTicket
//
// PARAMS:
//   - request: the arguments to TrainTicket
//
// RETURNS:
//   - TrainTicketResponse: The return type of the TrainTicket interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TrainTicket(request *TrainTicketRequest) (*TrainTicketResponse, error) {
	result := &TrainTicketResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTrainTicketUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TwoFactorsVerification
//
// PARAMS:
//   - request: the arguments to TwoFactorsVerification
//
// RETURNS:
//   - TwoFactorsVerificationResponse: The return type of the TwoFactorsVerification interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TwoFactorsVerification(request *TwoFactorsVerificationRequest) (*TwoFactorsVerificationResponse, error) {
	result := &TwoFactorsVerificationResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTwoFactorsVerificationUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UsedVehicleInvoice
//
// PARAMS:
//   - request: the arguments to UsedVehicleInvoice
//
// RETURNS:
//   - UsedVehicleInvoiceResponse: The return type of the UsedVehicleInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UsedVehicleInvoice(request *UsedVehicleInvoiceRequest) (*UsedVehicleInvoiceResponse, error) {
	result := &UsedVehicleInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUsedVehicleInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VatInvoice
//
// PARAMS:
//   - request: the arguments to VatInvoice
//
// RETURNS:
//   - VatInvoiceResponse: The return type of the VatInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VatInvoice(request *VatInvoiceRequest) (*VatInvoiceResponse, error) {
	result := &VatInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVatInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleCertificate
//
// PARAMS:
//   - request: the arguments to VehicleCertificate
//
// RETURNS:
//   - VehicleCertificateResponse: The return type of the VehicleCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleCertificate(request *VehicleCertificateRequest) (*VehicleCertificateResponse, error) {
	result := &VehicleCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleInvoice
//
// PARAMS:
//   - request: the arguments to VehicleInvoice
//
// RETURNS:
//   - VehicleInvoiceResponse: The return type of the VehicleInvoice interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleInvoice(request *VehicleInvoiceRequest) (*VehicleInvoiceResponse, error) {
	result := &VehicleInvoiceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleInvoiceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleLicense
//
// PARAMS:
//   - request: the arguments to VehicleLicense
//
// RETURNS:
//   - VehicleLicenseResponse: The return type of the VehicleLicense interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleLicense(request *VehicleLicenseRequest) (*VehicleLicenseResponse, error) {
	result := &VehicleLicenseResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleLicenseUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleRegCertificate
//
// PARAMS:
//   - request: the arguments to VehicleRegCertificate
//
// RETURNS:
//   - VehicleRegCertificateResponse: The return type of the VehicleRegCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleRegCertificate(request *VehicleRegCertificateRequest) (*VehicleRegCertificateResponse, error) {
	result := &VehicleRegCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleRegCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleRegistrationCertificate
//
// PARAMS:
//   - request: the arguments to VehicleRegistrationCertificate
//
// RETURNS:
//   - VehicleRegistrationCertificateResponse: The return type of the VehicleRegistrationCertificate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleRegistrationCertificate(request *VehicleRegistrationCertificateRequest) (*VehicleRegistrationCertificateResponse, error) {
	result := &VehicleRegistrationCertificateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleRegistrationCertificateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VinCode
//
// PARAMS:
//   - request: the arguments to VinCode
//
// RETURNS:
//   - VinCodeResponse: The return type of the VinCode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VinCode(request *VinCodeRequest) (*VinCodeResponse, error) {
	result := &VinCodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVinCodeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Waybill
//
// PARAMS:
//   - request: the arguments to Waybill
//
// RETURNS:
//   - WaybillResponse: The return type of the Waybill interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Waybill(request *WaybillRequest) (*WaybillResponse, error) {
	result := &WaybillResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getWaybillUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// WebImage
//
// PARAMS:
//   - request: the arguments to WebImage
//
// RETURNS:
//   - WebImageResponse: The return type of the WebImage interface.
//   - error: nil if success otherwise the specific error
func (c *Client) WebImage(request *WebImageRequest) (*WebImageResponse, error) {
	result := &WebImageResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getWebImageUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// WebImageLoc
//
// PARAMS:
//   - request: the arguments to WebImageLoc
//
// RETURNS:
//   - WebImageLocResponse: The return type of the WebImageLoc interface.
//   - error: nil if success otherwise the specific error
func (c *Client) WebImageLoc(request *WebImageLocRequest) (*WebImageLocResponse, error) {
	result := &WebImageLocResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getWebImageLocUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// WeightNote
//
// PARAMS:
//   - request: the arguments to WeightNote
//
// RETURNS:
//   - WeightNoteResponse: The return type of the WeightNote interface.
//   - error: nil if success otherwise the specific error
func (c *Client) WeightNote(request *WeightNoteRequest) (*WeightNoteResponse, error) {
	result := &WeightNoteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getWeightNoteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
