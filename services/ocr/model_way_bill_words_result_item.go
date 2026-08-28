package ocr

type WayBillWordsResultItem struct {
	ImageInfo         *ImageInfo         `json:"image_info,omitempty"`
	BarCode           []*WaybillWordItem `json:"bar_code,omitempty"`
	WaybillNumber     []*WaybillWordItem `json:"waybill_number,omitempty"`
	ThreeSegmentCode  []*WaybillWordItem `json:"three_segment_code,omitempty"`
	RecipientName     []*WaybillWordItem `json:"recipient_name,omitempty"`
	SenderName        []*WaybillWordItem `json:"sender_name,omitempty"`
	RecipientAddr     []*WaybillWordItem `json:"recipient_addr,omitempty"`
	SenderAddr        []*WaybillWordItem `json:"sender_addr,omitempty"`
	RecipientPhone    []*WaybillWordItem `json:"recipient_phone,omitempty"`
	SenderPhone       []*WaybillWordItem `json:"sender_phone,omitempty"`
	VirtualNumber     []*WaybillWordItem `json:"virtual_number,omitempty"`
	VirtualNumberLast []*WaybillWordItem `json:"virtual_number_last,omitempty"`
	IsVirtualWaybill  []*WaybillWordItem `json:"is_virtual_waybill,omitempty"`
}
