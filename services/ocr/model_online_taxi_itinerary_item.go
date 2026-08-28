package ocr

type OnlineTaxiItineraryItem struct {
	ItemId           *string `json:"ItemId,omitempty"`
	PickupTime       *string `json:"PickupTime,omitempty"`
	PickupDate       *string `json:"PickupDate,omitempty"`
	CarType          *string `json:"CarType,omitempty"`
	Distance         *string `json:"Distance,omitempty"`
	StartPlace       *string `json:"StartPlace,omitempty"`
	DestinationPlace *string `json:"DestinationPlace,omitempty"`
	City             *string `json:"City,omitempty"`
	Fare             *string `json:"Fare,omitempty"`
	ItemProvider     *string `json:"item_provider,omitempty"`
}
