package ocr

type SmartStructLineInfo struct {
	ObjectId             *int32               `json:"object_id,omitempty"`
	BlockId              *int32               `json:"block_id,omitempty"`
	Word                 *string              `json:"word,omitempty"`
	LineClass            *string              `json:"line_class,omitempty"`
	LineClassProbability *float64             `json:"line_class_probability,omitempty"`
	LineProbability      *float64             `json:"line__probability,omitempty"`
	Left                 *float64             `json:"left,omitempty"`
	Top                  *float64             `json:"top,omitempty"`
	Width                *float64             `json:"width,omitempty"`
	Height               *float64             `json:"height,omitempty"`
	LineLocation         *SmartStructLocation `json:"line_location,omitempty"`
}
