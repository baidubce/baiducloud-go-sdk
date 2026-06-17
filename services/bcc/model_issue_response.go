package bcc

type IssueResponse struct {
	IssueName        *string `json:"issueName,omitempty"`
	IssueAlias       *string `json:"issueAlias,omitempty"`
	IssueEffect      *string `json:"issueEffect,omitempty"`
	IssueDescription *string `json:"issueDescription,omitempty"`
	IssueOccurTime   *string `json:"issueOccurTime,omitempty"`
	IssueSource      *string `json:"issueSource,omitempty"`
}
