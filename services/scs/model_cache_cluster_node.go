package scs

type CacheClusterNode struct {
	InstanceId *string `json:"instanceId,omitempty"`
	FlavorInGB *string `json:"flavorInGB,omitempty"`
	HashName   *string `json:"hashName,omitempty"`
	Domain     *string `json:"domain,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ShardId    *string `json:"shardId,omitempty"`
}
