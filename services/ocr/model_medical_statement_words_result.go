package ocr

type MedicalStatementWordsResult struct {
	AdmissionDate          *MedicalStatementFieldValue `json:"AdmissionDate,omitempty"`
	DischargeDate          *MedicalStatementFieldValue `json:"DischargeDate,omitempty"`
	Name                   *MedicalStatementFieldValue `json:"Name,omitempty"`
	AmountInFiguers        *MedicalStatementFieldValue `json:"AmountInFiguers,omitempty"`
	SelfPaymentAmount      *MedicalStatementFieldValue `json:"SelfPaymentAmount,omitempty"`
	MedicalInsuranceAmount *MedicalStatementFieldValue `json:"MedicalInsuranceAmount,omitempty"`
}
