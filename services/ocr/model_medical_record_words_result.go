package ocr

type MedicalRecordWordsResult struct {
	RecordNum           *MedicalRecordFieldValue `json:"RecordNum,omitempty"`
	Name                *MedicalRecordFieldValue `json:"Name,omitempty"`
	Sex                 *MedicalRecordFieldValue `json:"Sex,omitempty"`
	Birthday            *MedicalRecordFieldValue `json:"Birthday,omitempty"`
	Age                 *MedicalRecordFieldValue `json:"Age,omitempty"`
	Career              *MedicalRecordFieldValue `json:"Career,omitempty"`
	MaritalStatus       *MedicalRecordFieldValue `json:"MaritalStatus,omitempty"`
	Nation              *MedicalRecordFieldValue `json:"Nation,omitempty"`
	ID                  *MedicalRecordFieldValue `json:"ID,omitempty"`
	Nationality         *MedicalRecordFieldValue `json:"Nationality,omitempty"`
	AdmissionDepartment *MedicalRecordFieldValue `json:"AdmissionDepartment,omitempty"`
	DischargeDepartment *MedicalRecordFieldValue `json:"DischargeDepartment,omitempty"`
	HospitalDay         *MedicalRecordFieldValue `json:"HospitalDay,omitempty"`
	Allergy             *MedicalRecordFieldValue `json:"Allergy,omitempty"`
	BloodType           *MedicalRecordFieldValue `json:"BloodType,omitempty"`
}
