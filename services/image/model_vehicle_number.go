package image

type VehicleNumber struct {
	Car       *int32 `json:"car,omitempty"`
	Truck     *int32 `json:"truck,omitempty"`
	Bus       *int32 `json:"bus,omitempty"`
	Motorbike *int32 `json:"motorbike,omitempty"`
	Tricycle  *int32 `json:"tricycle,omitempty"`
	Carplate  *int32 `json:"carplate,omitempty"`
}
