package ocr

type HouseholdRegisterResult struct {
	HouseholdNum    *HouseholdWordItem `json:"HouseholdNum,omitempty"`
	Name            *HouseholdWordItem `json:"Name,omitempty"`
	Relationship    *HouseholdWordItem `json:"Relationship,omitempty"`
	Sex             *HouseholdWordItem `json:"Sex,omitempty"`
	BirthAddress    *HouseholdWordItem `json:"BirthAddress,omitempty"`
	Nation          *HouseholdWordItem `json:"Nation,omitempty"`
	Birthday        *HouseholdWordItem `json:"Birthday,omitempty"`
	CardNo          *HouseholdWordItem `json:"CardNo,omitempty"`
	FormerName      *HouseholdWordItem `json:"FormerName,omitempty"`
	Hometown        *HouseholdWordItem `json:"Hometown,omitempty"`
	OtherAddress    *HouseholdWordItem `json:"OtherAddress,omitempty"`
	Belief          *HouseholdWordItem `json:"Belief,omitempty"`
	Height          *HouseholdWordItem `json:"Height,omitempty"`
	BloodType       *HouseholdWordItem `json:"BloodType,omitempty"`
	Education       *HouseholdWordItem `json:"Education,omitempty"`
	MaritalStatus   *HouseholdWordItem `json:"MaritalStatus,omitempty"`
	VeteranStatus   *HouseholdWordItem `json:"VeteranStatus,omitempty"`
	WorkAddress     *HouseholdWordItem `json:"WorkAddress,omitempty"`
	Career          *HouseholdWordItem `json:"Career,omitempty"`
	WWToCity        *HouseholdWordItem `json:"WWToCity,omitempty"`
	WWHere          *HouseholdWordItem `json:"WWHere,omitempty"`
	Date            *HouseholdWordItem `json:"Date,omitempty"`
	HouseholdType   *HouseholdWordItem `json:"HouseholdType,omitempty"`
	HouseholderName *HouseholdWordItem `json:"HouseholderName,omitempty"`
	Address         *HouseholdWordItem `json:"Address,omitempty"`
	IssueDate       *HouseholdWordItem `json:"IssueDate,omitempty"`
}
