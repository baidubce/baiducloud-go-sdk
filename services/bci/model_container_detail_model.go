package bci

type ContainerDetailModel struct {
	Name            *string          `json:"name,omitempty"`
	Image           *string          `json:"image,omitempty"`
	Cpu             *float32         `json:"cpu,omitempty"`
	Gpu             *float32         `json:"gpu,omitempty"`
	Memory          *float32         `json:"memory,omitempty"`
	WorkingDir      *string          `json:"workingDir,omitempty"`
	ImagePullPolicy *string          `json:"imagePullPolicy,omitempty"`
	Commands        []*string        `json:"commands,omitempty"`
	Args            []*string        `json:"args,omitempty"`
	Ports           []*Port          `json:"ports,omitempty"`
	VolumeMounts    []*VolumeMount   `json:"volumeMounts,omitempty"`
	Envs            []*Environment   `json:"envs,omitempty"`
	CreateTime      *string          `json:"createTime,omitempty"`
	UpdateTime      *string          `json:"updateTime,omitempty"`
	DeleteTime      *string          `json:"deleteTime,omitempty"`
	PreviousState   *ContainerStatus `json:"previousState,omitempty"`
	CurrentState    *ContainerStatus `json:"currentState,omitempty"`
	Ready           *bool            `json:"ready,omitempty"`
	RestartCount    *int32           `json:"restartCount,omitempty"`
}
