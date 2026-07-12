package oos

type DagInstance struct {
	Id                *string                 `json:"id,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	Revision          *int64                  `json:"revision,omitempty"`
	CreatedTimestamp  *int64                  `json:"createdTimestamp,omitempty"`
	UpdatedTimestamp  *int64                  `json:"updatedTimestamp,omitempty"`
	FinishedTimestamp *int64                  `json:"finishedTimestamp,omitempty"`
	Namespace         *string                 `json:"namespace,omitempty"`
	State             *string                 `json:"state,omitempty"`
	DagSpec           *DagSpec                `json:"dagSpec,omitempty"`
	InitContext       *interface{}            `json:"initContext,omitempty"`
	Parallelism       *int32                  `json:"parallelism,omitempty"`
	Manually          *bool                   `json:"manually,omitempty"`
	WorkerSelectors   []*TagSelector          `json:"workerSelectors,omitempty"`
	Tasks             []*ExecutionTaskSummary `json:"tasks,omitempty"`
	User              *interface{}            `json:"user,omitempty"`
	OperatorActions   *interface{}            `json:"operatorActions,omitempty"`
	DagActions        *DagActionModel         `json:"dagActions,omitempty"`
	Event             *EventModel             `json:"event,omitempty"`
	ParentTask        *ExecutionTaskSummary   `json:"parentTask,omitempty"`
	Trigger           *string                 `json:"trigger,omitempty"`
	InitOperators     []*TaskOperatorSummary  `json:"initOperators,omitempty"`
	Reason            *string                 `json:"reason,omitempty"`
	ErrorCode         *string                 `json:"errorCode,omitempty"`
	CronDagName       *string                 `json:"cronDagName,omitempty"`
	EventDagName      *string                 `json:"eventDagName,omitempty"`
}
