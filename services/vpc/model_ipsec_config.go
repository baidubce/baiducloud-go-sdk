package vpc

type IpsecConfig struct {
	IpsecEncAlg   *string `json:"ipsecEncAlg,omitempty"`
	IpsecAuthAlg  *string `json:"ipsecAuthAlg,omitempty"`
	IpsecPfs      *string `json:"ipsecPfs,omitempty"`
	IpsecLifetime *string `json:"ipsecLifetime,omitempty"`
}
