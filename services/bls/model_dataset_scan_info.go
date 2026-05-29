package bls

type DatasetScanInfo struct {
	Statistics  *Statistics `json:"statistics,omitempty"`
	IsTruncated *bool       `json:"isTruncated,omitempty"`
}
