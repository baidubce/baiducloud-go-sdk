package cprom

type GenerateInstanceTokenRequest struct {
	InstanceId *string `json:"-"`
	Action     *string `json:"-"`
	Token      *string `json:"token,omitempty"`
}
