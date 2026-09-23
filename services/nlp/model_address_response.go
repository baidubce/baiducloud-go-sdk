package nlp

import "github.com/baidubce/baiducloud-go-sdk/bce"

type AddressResponse struct {
	bce.BaseResponse
	ErrorCode    *int32   `json:"error_code,omitempty"`
	ErrorMsg     *string  `json:"error_msg,omitempty"`
	LogId        *int64   `json:"log_id,omitempty"`
	Text         *string  `json:"text,omitempty"`
	Province     *string  `json:"province,omitempty"`
	ProvinceCode *string  `json:"province_code,omitempty"`
	City         *string  `json:"city,omitempty"`
	CityCode     *string  `json:"city_code,omitempty"`
	County       *string  `json:"county,omitempty"`
	CountyCode   *string  `json:"county_code,omitempty"`
	Town         *string  `json:"town,omitempty"`
	TownCode     *string  `json:"town_code,omitempty"`
	Person       *string  `json:"person,omitempty"`
	Detail       *string  `json:"detail,omitempty"`
	Phonenum     *string  `json:"phonenum,omitempty"`
	Lat          *float32 `json:"lat,omitempty"`
	Lng          *float32 `json:"lng,omitempty"`
}
