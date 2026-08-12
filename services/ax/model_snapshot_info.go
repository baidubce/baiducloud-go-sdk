package ax

type SnapshotInfo struct {
	SnapshotID *string   `json:"snapshotID,omitempty"`
	Names      []*string `json:"names,omitempty"`
}
