package cprom

type User struct {
	UserId      *string `json:"userId,omitempty"`
	UserName    *string `json:"userName,omitempty"`
	UserType    *string `json:"userType,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Email       *string `json:"email,omitempty"`
}
