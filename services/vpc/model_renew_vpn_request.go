

package vpc

type RenewVpnRequest struct {
    VpnId *string `json:"-"`
    ClientToken *string `json:"-"`
    Billing *Billing `json:"billing,omitempty"`
}