package ocr

type SmartStructRelations struct {
	KvRelations    []*SmartStructKVRelation   `json:"kv_relations,omitempty"`
	TableRelations *SmartStructTableRelations `json:"table_relations,omitempty"`
}
