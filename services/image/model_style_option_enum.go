package image

// StyleOptionEnum the model 'StyleOptionEnum'
type StyleOptionEnum string

// List of StyleOptionEnum
const (
	StyleOptionEnumCartoon     StyleOptionEnum = "cartoon"
	StyleOptionEnumPencil      StyleOptionEnum = "pencil"
	StyleOptionEnumColorPencil StyleOptionEnum = "color_pencil"
	StyleOptionEnumWarm        StyleOptionEnum = "warm"
	StyleOptionEnumWave        StyleOptionEnum = "wave"
	StyleOptionEnumLavender    StyleOptionEnum = "lavender"
	StyleOptionEnumMononoke    StyleOptionEnum = "mononoke"
	StyleOptionEnumScream      StyleOptionEnum = "scream"
	StyleOptionEnumGothic      StyleOptionEnum = "gothic"
)

// All allowed values of StyleOptionEnum enum
var AllowedStyleOptionEnumEnumValues = []StyleOptionEnum{
	"cartoon",
	"pencil",
	"color_pencil",
	"warm",
	"wave",
	"lavender",
	"mononoke",
	"scream",
	"gothic",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StyleOptionEnum) IsValid() bool {
	for _, existing := range AllowedStyleOptionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
