package ocr

type BirthCertificateResult struct {
	BabyBirthday      *BabyBirthday      `json:"BabyBirthday,omitempty"`
	BirthProvince     *BirthProvince     `json:"BirthProvince,omitempty"`
	BirthCity         *BirthCity         `json:"BirthCity,omitempty"`
	BirthCounty       *BirthCounty       `json:"BirthCounty,omitempty"`
	BirthWeight       *BirthWeight       `json:"BirthWeight,omitempty"`
	BirthLength       *BirthLength       `json:"BirthLength,omitempty"`
	GestationalAge    *GestationalAge    `json:"GestationalAge,omitempty"`
	BabyName          *BabyName          `json:"BabyName,omitempty"`
	BabySex           *BabySex           `json:"BabySex,omitempty"`
	Code              *Code              `json:"Code,omitempty"`
	Hospital          *Hospital          `json:"Hospital,omitempty"`
	FatherName        *FatherName        `json:"FatherName,omitempty"`
	FatherID          *FatherID          `json:"FatherID,omitempty"`
	FatherNationality *FatherNationality `json:"FatherNationality,omitempty"`
	FatherEthnic      *FatherEthnic      `json:"FatherEthnic,omitempty"`
	FatherAddress     *FatherAddress     `json:"FatherAddress,omitempty"`
	FatherAge         *FatherAge         `json:"FatherAge,omitempty"`
	MotherName        *MotherName        `json:"MotherName,omitempty"`
	MotherID          *MotherID          `json:"MotherID,omitempty"`
	MotherNationality *MotherNationality `json:"MotherNationality,omitempty"`
	MotherEthnic      *MotherEthnic      `json:"MotherEthnic,omitempty"`
	MotherAddress     *MotherAddress     `json:"MotherAddress,omitempty"`
	MotherAge         *MotherAge         `json:"MotherAge,omitempty"`
}
