package ocr

type UsedVehicleInvoiceWordResult struct {
	InvoiceHeader                   *string                     `json:"InvoiceHeader,omitempty"`
	InvoiceCode                     *UsedVehicleInvoiceWordItem `json:"InvoiceCode,omitempty"`
	InvoiceNum                      *UsedVehicleInvoiceWordItem `json:"InvoiceNum,omitempty"`
	InvoiceDate                     *UsedVehicleInvoiceWordItem `json:"InvoiceDate,omitempty"`
	TaxCode                         *UsedVehicleInvoiceWordItem `json:"TaxCode,omitempty"`
	Purchaser                       *UsedVehicleInvoiceWordItem `json:"Purchaser,omitempty"`
	PurchaserCode                   *UsedVehicleInvoiceWordItem `json:"PurchaserCode,omitempty"`
	PurchaserAddress                *UsedVehicleInvoiceWordItem `json:"PurchaserAddress,omitempty"`
	PurchaserPhone                  *UsedVehicleInvoiceWordItem `json:"PurchaserPhone,omitempty"`
	Saler                           *UsedVehicleInvoiceWordItem `json:"Saler,omitempty"`
	SalerCode                       *UsedVehicleInvoiceWordItem `json:"SalerCode,omitempty"`
	SalerAddress                    *UsedVehicleInvoiceWordItem `json:"SalerAddress,omitempty"`
	SalerPhone                      *UsedVehicleInvoiceWordItem `json:"SalerPhone,omitempty"`
	LicensePlateNum                 *UsedVehicleInvoiceWordItem `json:"LicensePlateNum,omitempty"`
	RegistrationCode                *UsedVehicleInvoiceWordItem `json:"RegistrationCode,omitempty"`
	VehicleType                     *UsedVehicleInvoiceWordItem `json:"VehicleType,omitempty"`
	VinNum                          *UsedVehicleInvoiceWordItem `json:"VinNum,omitempty"`
	ManuModel                       *UsedVehicleInvoiceWordItem `json:"ManuModel,omitempty"`
	TransferVehicleManagementOffice *UsedVehicleInvoiceWordItem `json:"TransferVehicleManagementOffice,omitempty"`
	TotalCarPrice                   *UsedVehicleInvoiceWordItem `json:"TotalCarPrice,omitempty"`
	TotalCarPriceLow                *UsedVehicleInvoiceWordItem `json:"TotalCarPriceLow,omitempty"`
	UsedCarMarket                   *UsedVehicleInvoiceWordItem `json:"UsedCarMarket,omitempty"`
	TaxNum                          *UsedVehicleInvoiceWordItem `json:"TaxNum,omitempty"`
	TaxAddress                      *UsedVehicleInvoiceWordItem `json:"TaxAddress,omitempty"`
	TaxPhone                        *UsedVehicleInvoiceWordItem `json:"TaxPhone,omitempty"`
	SheetNum                        *UsedVehicleInvoiceWordItem `json:"SheetNum,omitempty"`
	InvoiceNumDigit                 *UsedVehicleInvoiceWordItem `json:"InvoiceNumDigit,omitempty"`
}
