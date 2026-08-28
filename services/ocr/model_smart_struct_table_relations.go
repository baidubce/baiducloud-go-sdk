package ocr

type SmartStructTableRelations struct {
	KkRelations []*SmartStructKVRelation `json:"kk_relations,omitempty"`
	KvRelations []*SmartStructKVRelation `json:"kv_relations,omitempty"`
	VvRelations []*SmartStructKVRelation `json:"vv_relations,omitempty"`
}
