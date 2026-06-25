package cprom

type MentionedUserConfig struct {
	AtAll   *bool     `json:"atAll,omitempty"`
	UserIds []*string `json:"userIds,omitempty"`
}
