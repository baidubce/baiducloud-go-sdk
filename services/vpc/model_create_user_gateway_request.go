package vpc

type CreateUserGatewayRequest struct {
	ClientToken *string `json:"-"`
	Name        *string `json:"name,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	Description *string `json:"description,omitempty"`
}
