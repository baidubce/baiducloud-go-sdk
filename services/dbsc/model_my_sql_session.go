package dbsc

type MySQLSession struct {
	Command  *string `json:"command,omitempty"`
	Db       *string `json:"db,omitempty"`
	Host     *string `json:"host,omitempty"`
	Id       *int32  `json:"id,omitempty"`
	Sqlstmt  *string `json:"sqlstmt,omitempty"`
	State    *string `json:"state,omitempty"`
	Time     *int32  `json:"time,omitempty"`
	TrxState *string `json:"trxState,omitempty"`
	TrxTime  *int32  `json:"trxTime,omitempty"`
	User     *string `json:"user,omitempty"`
}
