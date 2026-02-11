

package eip

type UnbindEipRequest struct {
    Eip *string `json:"-"`
    ClientToken *string `json:"-"`
}