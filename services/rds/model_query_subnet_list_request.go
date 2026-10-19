package rds

type QuerySubnetListRequest struct {
	ZoneName *string `json:"-"`
	VpcId    *string `json:"-"`
}
