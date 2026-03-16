

package vpc

type BindEipRequest struct {
    VpnId *string `json:"-"`
    ClientToken *string `json:"-"`
    Eip *string `json:"eip,omitempty"`
}