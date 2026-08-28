package ocr

type VehicleRegistrationCertificateRequest struct {
	Image *string `json:"image,omitempty"`
	Url   *string `json:"url,omitempty"`
}
