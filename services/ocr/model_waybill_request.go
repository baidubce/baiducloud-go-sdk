package ocr

type WaybillRequest struct {
	Image                    *string `json:"image,omitempty"`
	Url                      *string `json:"url,omitempty"`
	IsIdentifyVirtualWaybill *bool   `json:"is_identify_virtual_waybill,omitempty"`
}
