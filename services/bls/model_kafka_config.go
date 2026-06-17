package bls

// KafkaConfig the model 'KafkaConfig'
type KafkaConfig string

// List of KafkaConfig
const (
	KafkaConfigBrokers    KafkaConfig = "brokers"
	KafkaConfigTopic      KafkaConfig = "topic"
	KafkaConfigMaxretries KafkaConfig = "maxRetries"
)

// All allowed values of KafkaConfig enum
var AllowedKafkaConfigEnumValues = []KafkaConfig{
	"brokers",
	"topic",
	"maxRetries",
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KafkaConfig) IsValid() bool {
	for _, existing := range AllowedKafkaConfigEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}
