package bls

type Response struct {
	Took         *int32                  `json:"took,omitempty"`
	TimedOut     *bool                   `json:"timed_out,omitempty"`
	Shards       *Shard                  `json:"_shards,omitempty"`
	Hits         *map[string][]Hit       `json:"hits,omitempty"`
	Aggregations *map[string]interface{} `json:"aggregations,omitempty"`
}
