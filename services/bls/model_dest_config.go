package bls

type DestConfig struct {
	BOSPath                  *string   `json:"BOSPath,omitempty"`
	PartitionFormatTS        *string   `json:"partitionFormatTS,omitempty"`
	PartitionFormatLogStream *bool     `json:"partitionFormatLogStream,omitempty"`
	MaxObjectSize            *int32    `json:"maxObjectSize,omitempty"`
	CompressType             *string   `json:"compressType,omitempty"`
	DeliverInterval          *int32    `json:"deliverInterval,omitempty"`
	StorageFormat            *string   `json:"storageFormat,omitempty"`
	CsvHeadline              *bool     `json:"csvHeadline,omitempty"`
	CsvDelimiter             *string   `json:"csvDelimiter,omitempty"`
	CsvQuote                 *string   `json:"csvQuote,omitempty"`
	NullIdentifier           *string   `json:"nullIdentifier,omitempty"`
	SelectedColumnName       *string   `json:"selectedColumnName,omitempty"`
	SelectedColumnType       *string   `json:"selectedColumnType,omitempty"`
	FieldsName               []*string `json:"fieldsName,omitempty"`
	FieldsType               []*string `json:"fieldsType,omitempty"`
	ShipperType              *string   `json:"shipperType,omitempty"`
	KafkaConfig              *string   `json:"kafkaConfig,omitempty"`
	DestType                 *string   `json:"destType,omitempty"`
	LogStore                 *string   `json:"logStore,omitempty"`
	RateLimit                *int32    `json:"rateLimit,omitempty"`
	ClientCount              *int32    `json:"clientCount,omitempty"`
}
