package cprom

type DeleteInstanceRequest struct {
	InstanceId    *string `json:"-"`
	DeleteGrafana *bool   `json:"-"`
}
