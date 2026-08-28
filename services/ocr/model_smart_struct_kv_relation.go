package ocr

type SmartStructKVRelation struct {
	RootNode  *int32   `json:"root_node,omitempty"`
	LeafNodes []*int32 `json:"leaf_nodes,omitempty"`
}
