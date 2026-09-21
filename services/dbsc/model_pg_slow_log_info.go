package dbsc

type PGSlowLogInfo struct {
	Product        *string `json:"product,omitempty"`
	AppID          *string `json:"appID,omitempty"`
	AppName        *string `json:"appName,omitempty"`
	AppShortID     *string `json:"appShortID,omitempty"`
	ClusterID      *string `json:"clusterID,omitempty"`
	NodeID         *string `json:"nodeID,omitempty"`
	Uuid           *string `json:"uuid,omitempty"`
	Pid            *int64  `json:"pid,omitempty"`
	ClientIP       *string `json:"clientIP,omitempty"`
	CurrentDB      *string `json:"currentDB,omitempty"`
	CurrentUser    *string `json:"currentUser,omitempty"`
	Duration       *int64  `json:"duration,omitempty"`
	Start          *string `json:"start,omitempty"`
	End            *string `json:"end,omitempty"`
	Statement      *string `json:"statement,omitempty"`
	Fingerprint    *string `json:"fingerprint,omitempty"`
	FingerprintMD5 *string `json:"fingerprintMD5,omitempty"`
}
