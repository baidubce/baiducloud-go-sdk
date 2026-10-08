package scs

// Engine the model 'Engine'
type Engine string

// List of Engine
const (
	EngineRedis  Engine = "redis"
	EnginePegadb Engine = "PegaDB"
)

// All allowed values of Engine enum
var AllowedEngineEnumValues = []Engine{
	"redis",
	"PegaDB",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Engine) IsValid() bool {
	for _, existing := range AllowedEngineEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
