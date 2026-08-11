package ax

type QuerySandboxesRequest struct {
	Limit      *int32             `json:"limit,omitempty"`
	NextToken  *string            `json:"nextToken,omitempty"`
	SandboxIds []*string          `json:"sandboxIds,omitempty"`
	ImagePaths []*string          `json:"imagePaths,omitempty"`
	Metadata   *map[string]string `json:"metadata,omitempty"`
	State      []*string          `json:"state,omitempty"`
}
