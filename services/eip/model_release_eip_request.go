

package eip

type ReleaseEipRequest struct {
    Eip *string `json:"-"`
    ReleaseToRecycle *bool `json:"-"`
    ClientToken *string `json:"-"`
}