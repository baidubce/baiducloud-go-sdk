package ocr

type OnlineTaxiItineraryWordsResult struct {
	ServiceProvider *string                    `json:"ServiceProvider,omitempty"`
	StartTime       *string                    `json:"StartTime,omitempty"`
	EndTime         *string                    `json:"EndTime,omitempty"`
	Phone           *string                    `json:"Phone,omitempty"`
	ApplicationDate *string                    `json:"ApplicationDate,omitempty"`
	TotalFare       *string                    `json:"TotalFare,omitempty"`
	ItemNum         *string                    `json:"ItemNum,omitempty"`
	ServiceType     *string                    `json:"ServiceType,omitempty"`
	Items           []*OnlineTaxiItineraryItem `json:"items,omitempty"`
}
