

package eip

type CloseEipDirectAccessRequest struct {
    Eip *string `json:"-"`
    ClientToken *string `json:"-"`
}