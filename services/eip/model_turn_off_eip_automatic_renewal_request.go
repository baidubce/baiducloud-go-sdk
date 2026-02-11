

package eip

type TurnOffEipAutomaticRenewalRequest struct {
    Eip *string `json:"-"`
    ClientToken *string `json:"-"`
}