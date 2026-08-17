package ocr

type MedicalInvoiceWordsResult struct {
	BusinessNum         *MedicalInvoiceFieldValue             `json:"BusinessNum,omitempty"`
	InvoiceNum          *MedicalInvoiceFieldValue             `json:"InvoiceNum,omitempty"`
	HospitalNum         *MedicalInvoiceFieldValue             `json:"HospitalNum,omitempty"`
	HospitalName        *MedicalInvoiceFieldValue             `json:"HospitalName,omitempty"`
	RecordNum           *MedicalInvoiceFieldValue             `json:"RecordNum,omitempty"`
	HospitalDay         *MedicalInvoiceFieldValue             `json:"HospitalDay,omitempty"`
	AdmissionDate       *MedicalInvoiceFieldValue             `json:"AdmissionDate,omitempty"`
	DischargeDate       *MedicalInvoiceFieldValue             `json:"DischargeDate,omitempty"`
	DischargeDepartment *MedicalInvoiceFieldValue             `json:"DischargeDepartment,omitempty"`
	Name                *MedicalInvoiceFieldValue             `json:"Name,omitempty"`
	Sex                 *MedicalInvoiceFieldValue             `json:"Sex,omitempty"`
	HospitalType        *MedicalInvoiceFieldValue             `json:"HospitalType,omitempty"`
	SocialSecurityNum   *MedicalInvoiceFieldValue             `json:"SocialSecurityNum,omitempty"`
	InsuranceType       *MedicalInvoiceFieldValue             `json:"InsuranceType,omitempty"`
	ChargingUnit        *MedicalInvoiceFieldValue             `json:"ChargingUnit,omitempty"`
	Payee               *MedicalInvoiceFieldValue             `json:"Payee,omitempty"`
	Date                *MedicalInvoiceFieldValue             `json:"Date,omitempty"`
	AmountInWords       *MedicalInvoiceFieldValue             `json:"AmountInWords,omitempty"`
	AmountInFiguers     *MedicalInvoiceFieldValue             `json:"AmountInFiguers,omitempty"`
	InsurancePayment    *MedicalInvoiceFieldValue             `json:"InsurancePayment,omitempty"`
	PersonalPayment     *MedicalInvoiceFieldValue             `json:"PersonalPayment,omitempty"`
	PrepayAmount        *MedicalInvoiceFieldValue             `json:"PrepayAmount,omitempty"`
	PaymentAmount       *MedicalInvoiceFieldValue             `json:"PaymentAmount,omitempty"`
	RefundAmount        *MedicalInvoiceFieldValue             `json:"RefundAmount,omitempty"`
	ClinicNum           *MedicalInvoiceFieldValue             `json:"ClinicNum,omitempty"`
	CostCategories      [][]*MedicalInvoiceCostCategoryItem   `json:"CostCategories,omitempty"`
	CostDetail          [][]*MedicalInvoiceCostDetailItem     `json:"CostDetail,omitempty"`
	RegionSupplement    []*MedicalInvoiceRegionSupplementItem `json:"RegionSupplement,omitempty"`
}
