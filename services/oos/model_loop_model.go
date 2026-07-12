package oos

type LoopModel struct {
	InitContext     *interface{}   `json:"initContext,omitempty"`
	WorkerSelectors []*TagSelector `json:"workerSelectors,omitempty"`
}
