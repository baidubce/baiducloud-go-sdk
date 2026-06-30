package bci

type Container struct {
	Name            *string                   `json:"name,omitempty"`
	Image           *string                   `json:"image,omitempty"`
	Memory          *float32                  `json:"memory,omitempty"`
	Cpu             *float32                  `json:"cpu,omitempty"`
	Gpu             *float32                  `json:"gpu,omitempty"`
	WorkingDir      *string                   `json:"workingDir,omitempty"`
	ImagePullPolicy *string                   `json:"imagePullPolicy,omitempty"`
	Commands        []*string                 `json:"commands,omitempty"`
	Args            []*string                 `json:"args,omitempty"`
	VolumeMounts    []*VolumeMount            `json:"volumeMounts,omitempty"`
	Ports           []*Port                   `json:"ports,omitempty"`
	EnvironmentVars []*Environment            `json:"environmentVars,omitempty"`
	LivenessProbe   *Probe                    `json:"livenessProbe,omitempty"`
	ReadinessProbe  *Probe                    `json:"readinessProbe,omitempty"`
	StartupProbe    *Probe                    `json:"startupProbe,omitempty"`
	Stdin           *bool                     `json:"stdin,omitempty"`
	StdinOnce       *bool                     `json:"stdinOnce,omitempty"`
	Tty             *bool                     `json:"tty,omitempty"`
	SecurityContext *ContainerSecurityContext `json:"securityContext,omitempty"`
}
