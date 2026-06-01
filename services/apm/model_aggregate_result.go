package apm

type AggregateResult struct {
	Sum          *float32 `json:"sum,omitempty"`
	SumPerSecond *float32 `json:"sumPerSecond,omitempty"`
}
