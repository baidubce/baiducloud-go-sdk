package rds

type DeleteTheDatabaseRequest struct {
	InstanceId *string `json:"-"`
	DbName     *string `json:"-"`
}
