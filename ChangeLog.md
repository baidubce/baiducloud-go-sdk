2026-07-24 Version: v0.0.112
- 涉及产品: BCC，解绑自动快照策略删除deleteAutoSnapshot字段
- 涉及产品: BCC，跨地域复制快照删除uuid字段
- 涉及产品: BCC，修改自定义镜像名称新增imageId字段

2026-07-22 Version: v0.0.111
- 涉及产品: CLOUDASSISTANT，SDK版本更新

2026-07-21 Version: v0.0.110
- 涉及产品: AS，SDK版本更新

2026-07-21 Version: v0.0.109
- 涉及产品: IAM，创建永久API Key响应新增id、tokenId、services等字段
- 涉及产品: IAM，更新/查询/解码永久API Key响应新增访问控制及时间字段
- 涉及产品: IAM，获取永久API Key列表及删除接口新增状态与分页字段
- 涉及产品: IAM，创建短期API Key响应新增userId、token及有效期字段

2026-07-17 Version: v0.0.108
执行命令接口新增实例与网络相关字段
- 涉及产品: CLOUDASSISTANT，执行命令新增实例ID、名称、内外网IP及带宽字段
- 涉及产品: CLOUDASSISTANT，执行命令新增ref字段
- 涉及产品: CLOUDASSISTANT，命令列表查询删除instanceType字段

2026-07-17 Version: v0.0.107
- 涉及产品: CLOUDASSISTANT，SDK版本更新

2026-07-16 Version: v0.0.106
- 涉及产品: BLB，releaseAppBlb释放实例接口query参数新增clientToken字段

2026-07-16 Version: v0.0.105
- 涉及产品: BLB，TCP/HTTP/UDP/SSL/HTTPS监听器创建与更新接口新增clientToken参数
- 涉及产品: BLB，添加、更新、释放后端服务器及释放监听器接口同步变更

2026-07-16 Version: v0.0.104
- 涉及产品: CLOUDASSISTANT，SDK版本更新

2026-07-15 Version: v0.0.103
- 涉及产品: BLB，创建IP组删除memberId、portList、healthCheckPortType、status字段
- 涉及产品: BLB，查询IP组删除healthCheckHost字段
- 涉及产品: BLB，创建/更新IP组成员删除ip、portList等冗余字段
- 涉及产品: BLB，新增计费变更及变配接口支持

2026-07-15 Version: v0.0.102
- 涉及产品: CLOUDASSISTANT，新增命令的创建、查询、修改、删除接口
- 涉及产品: CLOUDASSISTANT，新增命令执行及执行记录查询接口
- 涉及产品: CLOUDASSISTANT，新增bsm-agent状态批量查询接口

2026-07-14 Version: v0.0.101
- 涉及产品: AS，新增伸缩组创建、查询、删除及列表接口
- 涉及产品: AS，新增伸缩规则创建、执行、查询、删除接口
- 涉及产品: AS，新增手动扩缩容、调整节点数及伸缩日志查询接口
- 涉及产品: AS，新增节点保护、托管状态修改及节点移入移出接口

2026-07-14 Version: v0.0.100
- 涉及产品: VPC，SDK版本更新

2026-07-14 Version: v0.0.99
VPC专线网关健康检查接口新增三个字段
- 涉及产品: VPC，创建专线网关健康检查新增dcphyId、channelId、subnetId字段

2026-07-14 Version: v0.0.98
新增OAuth2客户端、IdP配置及Agent身份管理接口响应字段
- 涉及产品: AGENTIDENTITY，createOauth2Client/updateOauth2Client/getOauth2Client新增clientId、clientSecret等字段
- 涉及产品: AGENTIDENTITY，createIdpConfiguration/updateIdpConfiguration/getIdpConfiguration/disableIdpConfiguration/enableIdpConfiguration新增idpType、discoveryUrl等字段
- 涉及产品: AGENTIDENTITY，createAgent/getAgent/updateAgent新增bceDomainId、extra等字段，凭证提供方接口新增完整OIDC配置字段
- 涉及产品: AGENTIDENTITY，createUserPool/getUserPool/createUser/getUser/updateUser新增统计、端点及来源字段

2026-07-12 Version: v0.0.97
- 涉及产品: OOS，新增模板的创建、更新、删除、校验及查看接口
- 涉及产品: OOS，新增执行的创建、查询详情及列表接口
- 涉及产品: OOS，新增任务详情及子执行列表查询接口
- 涉及产品: OOS，新增操作符列表查询接口

2026-07-10 Version: v0.0.96
- 涉及产品: BLB，后端服务器增删改查接口新增desc字段

2026-07-10 Version: v0.0.95
- 涉及产品: PFS，createPfs接口新增zone（可用区）字段

2026-07-07 Version: v0.0.94
- 涉及产品: AGENTIDENTITY，新增Agent创建、查询、更新、删除接口
- 涉及产品: AGENTIDENTITY，新增用户池及用户的增删改查与批量获取接口
- 涉及产品: AGENTIDENTITY，新增IdP配置创建、更新、启用、禁用及OAuth2回调接口
- 涉及产品: AGENTIDENTITY，新增OAuth2客户端、凭证提供方及Token获取接口

2026-07-07 Version: v0.0.93
- 涉及产品: BCI，SDK版本更新

2026-07-06 Version: v0.0.92
- 涉及产品: AIHC，新增训练任务创建、查询、停止、删除、更新接口
- 涉及产品: AIHC，新增训练任务日志、事件、监控、节点查询接口
- 涉及产品: AIHC，新增获取WebTerminal地址及批量停止训练任务接口

2026-06-30 Version: v0.0.91
- 涉及产品: BCI，新增创建和列举镜像缓存接口
- 涉及产品: BCI，新增创建、查询、删除实例接口
- 涉及产品: BCI，支持批量删除实例及镜像缓存

2026-06-30 Version: v0.0.90
- 涉及产品: IAM，SDK版本更新

2026-06-30 Version: v0.0.89
- 涉及产品: IAM，SDK版本更新

2026-06-29 Version: v0.0.88
- 涉及产品: BLS，SDK版本更新

2026-06-26 Version: v0.0.87
- 涉及产品: BLS，pullLogRecord/queryLogRecord/queryLogHistogram新增logStoreType字段

2026-06-25 Version: v0.0.86
- 涉及产品: VPC，创建专线网关健康检查删除healthCheckPort字段

2026-06-25 Version: v0.0.85
- 涉及产品: CPROM，新增告警增删改查及事件认领接口
- 涉及产品: CPROM，新增监控实例创建、释放、Token生成及列表查询接口
- 涉及产品: CPROM，新增Pod/Service Monitor及自定义采集任务管理接口
- 涉及产品: CPROM，新增通知策略增删改查及CCE集群关联接口

2026-06-22 Version: v0.0.84
- 涉及产品: BLS，快速查询创建/详情/列表接口新增logStoreType字段

2026-06-22 Version: v0.0.83
- 涉及产品: BCC，新增修改快照信息、查询快照用量接口
- 涉及产品: BCC，查询实例券价格新增预付/后付类目价及成交价字段

2026-06-18 Version: v0.0.82
- 涉及产品: BCC，新增配置磁盘释放保护功能

2026-06-17 Version: v0.0.81
- 涉及产品: BCC，新增授权事件与验收事件接口
- 涉及产品: BCC，新增预期与非预期事件查询接口
- 涉及产品: BCC，新增预授权规则的创建、修改、删除接口
- 涉及产品: BCC，新增获取授权规则列表接口
- 涉及产品: STS，新增AssumeRole角色扮演接口
- 涉及产品: STS，新增GetSessionToken获取临时Token接口

2026-06-17 Version: v0.0.80
- 涉及产品: BCM，修复describeMetricDataLatest等3个接口中timestmap字段拼写错误，更正为timestamp

2026-06-17 Version: v0.0.79
- 涉及产品: BCC，新增专属集群创建、查询、变配、续费及标签管理接口
- 涉及产品: BCC，新增预留实例券创建、调整、续费、转移及自动续费管理接口
- 涉及产品: BCC，新增EHC集群创建、修改、删除及列表查询接口
- 涉及产品: BCC，新增可用区、套餐规格、价格查询及任务管理接口
2026-06-17 Version: v0.0.81
- 涉及产品: BLS，日志投递任务状态字段由desiredStatus改为status
- 涉及产品: BLS，报警策略/历史详情接口参数从path移至query
- 涉及产品: BLS，启用/禁用/删除报警策略支持批量操作，字段name改为names
- 涉及产品: BLS，日志集模版接口移除createdTimestamp和updatedTimestamp字段

2026-06-17 Version: v0.0.80
- 涉及产品: BLB，创建服务器组删除privateIp、portList等8个字段
- 涉及产品: BLB，修改后端服务器删除privateIp、status等9个字段
- 涉及产品: BLB，添加后端服务器删除privateIp、policyId等9个字段

2026-06-17 Version: v0.0.79
- 涉及产品: BLS，新增日志集、日志流、索引的创建/查询/更新/删除接口
- 涉及产品: BLS，新增日志投递任务管理及批量操作接口
- 涉及产品: BLS，新增报警策略管理、执行统计及历史记录查询接口
- 涉及产品: BLS，新增快速查询、日志视图、日志集模版管理接口

2026-06-16 Version: v0.0.78
- 涉及产品: BLB，SDK版本更新

2026-06-16 Version: v0.0.77
- 涉及产品: CFW，新增查询/创建/修改网络型CFW策略接口
- 涉及产品: CFW，完善应用型CFW策略的增删改查接口
- 涉及产品: CFW，支持批量创建/修改/删除应用型CFW规则

2026-06-16 Version: v0.0.76
新增BCC云服务器批量接口支持抢占实例、快照、安全组等功能
- 涉及产品: BCC，新增抢占实例创建、定价及订单取消接口
- 涉及产品: BCC，新增快照全生命周期管理及跨地域复制、共享接口
- 涉及产品: BCC，新增安全组及规则的创建、授权、撤销、删除接口
- 涉及产品: BCC，新增密钥对与部署集的完整管理接口
BLS新增下载任务相关接口并更新日志组字段
- 涉及产品: BLS，新增创建、查询、列表、删除、获取链接五个下载任务接口
- 涉及产品: BLS，updateProject接口description字段类型改为string并新增top字段

2026-06-15 Version: v0.0.75
- 涉及产品: BCM，新增报警模版的创建、查询、更新、删除、导入导出接口
- 涉及产品: BCM，新增实例组的创建、查询、更新、删除及实例管理接口
- 涉及产品: BCM，新增通知模版的创建、查询、更新、删除接口
- 涉及产品: BCM，新增系统模版规则查询及通知对象列表获取接口

2026-06-15 Version: v0.0.74
- 涉及产品: AIHC，SDK版本更新

2026-06-15 Version: v0.0.73
- 涉及产品: ET，支持删除物理专线资源

2026-06-15 Version: v0.0.72
- 涉及产品: BCC，新增实例生命周期管理（创建、启停、重启、重装、释放）
- 涉及产品: BCC，新增实例计费变更（续费、转包年包月、转按量、自动续费）
- 涉及产品: BCC，新增实例网络管理（变更VPC/子网、IPv6增删、网卡查询）
- 涉及产品: BCC，新增实例属性管理（安全组、标签、角色、密码、描述绑定）
- 涉及产品: RAPIDFS，元数据同步规则列表及详情接口字段IntervalMinutes更正为intervalMinutes
- 涉及产品: RAPIDFS，缓存管理规则任务列表字段cacheRuleJobInfos和cacheRuleJobId更正为cacheJobInfos和cacheJobId

2026-06-12 Version: v0.0.71
- 涉及产品: AIHC，SDK版本更新

2026-06-11 Version: v0.0.70
- 涉及产品: AIHC，新增模型增删改查及版本管理接口
- 涉及产品: AIHC，新增数据集增删改查及版本管理接口

2026-06-11 Version: v0.0.69
新增BCC自定义镜像管理相关接口
- 涉及产品: BCC，新增镜像创建、查询、删除、重命名接口
- 涉及产品: BCC，新增镜像共享、取消共享及共享用户查询接口
- 涉及产品: BCC，新增镜像导入、跨地域复制及取消复制接口
- 涉及产品: BCC，新增镜像标签绑定、OS查询及可用镜像查询接口

2026-06-10 Version: v0.0.68
- 涉及产品: BCC，支持CDS磁盘创建、挂载、卸载、释放等生命周期管理
- 涉及产品: BCC，支持磁盘扩容、类型变更、计费变更及续费操作
- 涉及产品: BCC，支持磁盘标签绑定解绑、重命名及属性修改
- 涉及产品: BCC，支持磁盘详情查询、列表查询、价格查询及变配进度查询

2026-06-09 Version: v0.0.67
- 涉及产品: CCR，新增创建、查询、更新命名空间接口

2026-06-08 Version: v0.0.66
- 涉及产品: BLB，createAppBlb创建实例接口变更
- 涉及产品: BLB，describeAppBlbs查询实例新增query.type字段

2026-06-04 Version: v0.0.65
- 涉及产品: PRIVATEZONE，新增DNS解析器增删改查及转发规则管理接口
- 涉及产品: PRIVATEZONE，新增VPC与转发规则关联/解关联接口
- 涉及产品: PRIVATEZONE，unbindVpc和enableRecord接口删除query.action字段

2026-06-03 Version: v0.0.64
- 涉及产品: RAPIDFS，元数据同步、缓存规则、实例管理等接口统一移除query.action参数

2026-06-03 Version: v0.0.63
- 涉及产品: CCR，SDK版本更新

2026-06-03 Version: v0.0.62
- 涉及产品: PFS，新增数据流动任务的创建、查询、暂停、启动、取消、删除及日志查询
- 涉及产品: PFS，新增生命周期规则的创建、更新、删除、查询及执行日志查询

2026-06-03 Version: v0.0.61
- 涉及产品: CCR，新增实例、命名空间、镜像仓库的增删查改接口
- 涉及产品: CCR，新增触发器、实例同步规则的创建、执行、管理接口
- 涉及产品: CCR，新增镜像迁移规则及执行记录查询接口
- 涉及产品: CCR，新增Helm Chart、机器人账号及网络配置管理接口

2026-06-03 Version: v0.0.60
新增RAPIDFS缓存实例及数据源管理等系列接口
- 涉及产品: RAPIDFS，新增实例创建及扩缩容接口
- 涉及产品: RAPIDFS，新增数据源导入及管理接口
- 涉及产品: RAPIDFS，新增缓存规则及节点管理接口
- 涉及产品: RAPIDFS，新增元数据同步规则管理接口

2026-06-02 Version: v0.0.59
新增PFS Fileset相关管理接口
- 涉及产品: PFS，支持Fileset的创建、删除与更新操作
- 涉及产品: PFS，支持Fileset列表查询与详情查询

2026-06-02 Version: v0.0.58
- 涉及产品: PFS，支持查询PFS实例挂载客户端列表
- 涉及产品: PFS，支持查询挂载服务挂载客户端列表

2026-06-02 Version: v0.0.57
- 涉及产品: BLB，支持集群的创建、升级与续费
- 涉及产品: BLB，支持查询集群列表及集群详情
- 涉及产品: BLB，支持查询集群下BLB实例及更新集群

2026-06-01 Version: v0.0.56
- 涉及产品: BLB，支持服务发布点的创建、更新与删除
- 涉及产品: BLB，支持服务发布点列表及详情查询
- 涉及产品: BLB，支持服务发布点鉴权信息的增改删
- 涉及产品: BLB，支持服务发布点与实例的绑定和解绑

2026-06-01 Version: v0.0.55
- 涉及产品: APM，新增服务配置查询与管理接口
- 涉及产品: APM，新增链路追踪与拓扑分析接口
- 涉及产品: APM，新增告警策略全生命周期管理接口
- 涉及产品: APM，新增LLM链路追踪与会话统计接口

2026-06-01 Version: v0.0.54
- 涉及产品: ET，支持物理专线的申请与更新
- 涉及产品: ET，支持专线通道的新建、更新、删除及重新提交
- 涉及产品: ET，支持专线通道BFD、路由规则及路由参数的增删改查
- 涉及产品: ET，支持专线通道关联、解关联、IPv6开关及用户对象管理

2026-05-30 Version: v0.0.53
- 涉及产品: PFS，支持创建、删除、查询实例及标签更新和列表获取

2026-05-29 Version: v0.0.52
- 涉及产品: CFS，updateFileSystemLabels接口删除query.tag字段

2026-05-29 Version: v0.0.51
- 涉及产品: CFS，新增文件系统创建、更新、查询、释放接口
- 涉及产品: CFS，新增权限组及权限组规则的增删改查接口
- 涉及产品: CFS，新增挂载点创建、删除、查询及权限组修改接口
- 涉及产品: CFS，新增查询挂载客户端接口

2026-05-29 Version: v0.0.50
- 涉及产品: BLS，SDK版本更新

2026-05-29 Version: v0.0.49
- 涉及产品: BLB，支持IP组的创建、查询、更新与删除
- 涉及产品: BLB，支持IP组成员的创建、查询、更新与删除
- 涉及产品: BLB，支持IP组协议的创建、更新与删除

2026-05-29 Version: v0.0.48
- 涉及产品: BLS，新增日志组创建、更新、删除、列表及详情接口
- 涉及产品: BLS，新增日志推送与获取接口
- 涉及产品: BLS，新增日志检索分析及直方图查询接口

2026-05-28 Version: v0.0.47
- 涉及产品: DNS,支持开启和关闭解析记录

2026-05-28 Version: v0.0.46
- 涉及产品: BLB，支持查询普通安全组和企业安全组
- 涉及产品: BLB，支持绑定与解绑普通安全组
- 涉及产品: BLB，支持绑定与解绑企业安全组

2026-05-28 Version: v0.0.45
- 涉及产品: BLB，支持查询后端服务器列表及健康状态
- 涉及产品: BLB，支持添加、更新、释放后端服务器

2026-05-28 Version: v0.0.44
- 涉及产品: BLB，支持创建TCP/UDP/HTTP/HTTPS/SSL监听器
- 涉及产品: BLB，支持查询各协议类型监听器及全量监听列表
- 涉及产品: BLB，支持更新TCP/UDP/HTTP/HTTPS/SSL监听器配置
- 涉及产品: BLB，支持释放监听器资源

2026-05-28 Version: v0.0.43
- 涉及产品: CSN，attachCsnInstance/detachCsnInstance/resizeCsnBp/unbindCsnBp/bindCsnBp删除query.action字段

2026-05-28 Version: v0.0.42
- 涉及产品: PRIVATEZONE，disableRecord接口删除query.action字段
- 涉及产品: PRIVATEZONE，bindVpc接口删除query.action字段

2026-05-27 Version: v0.0.41
- 涉及产品: CSN，新增云智能网创建、查询、更新、删除及实例挂载接口
- 涉及产品: CSN，新增带宽包创建、查询、更新、删除及绑定解绑接口
- 涉及产品: CSN，新增地域带宽及TGW查询管理接口
- 涉及产品: CSN，新增路由条目、路由表及学习关联关系管理接口

2026-05-27 Version: v0.0.40
- 涉及产品: PRIVATEZONE，新增disableRecord暂停解析记录能力

2026-05-27 Version: v0.0.39
- 涉及产品: BLB，新增服务器组的创建、查询、更新、删除接口
- 涉及产品: BLB，新增服务器组端口的创建、更新、删除接口
- 涉及产品: BLB，新增后端服务器的添加、修改、移除及查询接口
- 涉及产品: BLB，实例详情新增name、vpcId、ipv6、supportAcl字段

2026-05-27 Version: v0.0.38
- 涉及产品: VPC，SDK版本更新

2026-05-26 Version: v0.0.37
- 涉及产品: VPC，SDK版本更新

2026-05-26 Version: v0.0.36
- 涉及产品: VPC，SDK版本更新

2026-05-26 Version: v0.0.35
- 涉及产品: VPC，addAclRule新增ipVersion字段并删除原ipversion字段
- 涉及产品: VPC，queryNetworkDetectionList的destPort类型由integer改为string
- 涉及产品: VPC，queryNetworkDetectionDetails的destPort类型由integer改为string

2026-05-26 Version: v0.0.34
- 涉及产品: VPC，updatePeerConn及updatePeerConnDeleteProtect新增clientToken字段
- 涉及产品: VPC，listPeerConn的vpcId由路径参数改为查询参数
- 涉及产品: VPC，listPeerConn和getPeerConn删除多个连接状态枚举字段

2026-05-26 Version: v0.0.33
- 涉及产品: VPC，createNat新增预留时长及单位字段支持

2026-05-26 Version: v0.0.32
- 涉及产品: VPC，getNat新增sessionConfig及TCP/UDP/ICMP超时字段

2026-05-26 Version: v0.0.31
- 涉及产品: VPC，createNat新增sessionConfig会话配置及TCP/UDP/ICMP超时参数

2026-05-26 Version: v0.0.30
- 涉及产品: VPC，SDK版本更新

2026-05-25 Version: v0.0.29
- 涉及产品: VPC，listDnatRule接口publicPort和privatePort字段类型由string改为integer

2026-05-25 Version: v0.0.28
- 涉及产品: VPC，queryAcl和queryAclRules的position字段类型从string修改为integer

2026-05-25 Version: v0.0.27
- 涉及产品: VPC，新增NAT网关创建、查询、更新、释放、续费、变配接口
- 涉及产品: VPC，新增SNAT规则创建、更新、查询、批量创建、删除接口
- 涉及产品: VPC，新增DNAT规则创建、更新、查询、批量创建、删除接口
- 涉及产品: VPC，新增NAT网关绑定/解绑EIP及释放保护开关接口

2026-05-25 Version: v0.0.26
- 涉及产品: BLB，更新实例新增query.clientToken字段，支持幂等控制
- 涉及产品: BLB，涉及创建、释放、查询、ACL更新及保护开关等实例接口变更

2026-05-25 Version: v0.0.25
- 涉及产品: VPC，新增弹性网卡创建、挂载、绑定EIP及内网IP批量管理接口
- 涉及产品: VPC，新增高可用虚拟IP创建、绑定解绑实例及EIP全套接口
- 涉及产品: VPC，新增IPv6网关创建、限速策略及只出不进策略管理接口
- 涉及产品: VPC，新增ACL规则、IP地址组族及网络探测全套管理接口

2026-05-22 Version: v0.0.24
- 涉及产品: BLB，支持创建、更新、查询及释放实例
- 涉及产品: BLB，支持查询实例列表及详情
- 涉及产品: BLB，支持更新实例ACL功能
- 涉及产品: BLB，支持修改实例保护开关

2026-05-22 Version: v0.0.23
- 涉及产品: BLB，新增应用型实例创建、查询、更新、释放接口
- 涉及产品: BLB，新增TCP/UDP/HTTP/HTTPS/SSL监听器增删改查接口
- 涉及产品: BLB，新增转发策略创建、修改、查询、删除接口

2026-05-22 Version: v0.0.22
- 涉及产品: CFW，支持创建、查询、修改、删除CFW策略
- 涉及产品: CFW，支持批量创建、修改、删除CFW规则
- 涉及产品: CFW，支持实例绑定、解绑、开启、关闭CFW防护
- 涉及产品: CFW，支持查询防护边界实例列表

2026-05-22 Version: v0.0.21
- 涉及产品: VPC，新增普通安全组的增删查及规则授权/撤销/更新
- 涉及产品: VPC，新增企业安全组的创建/删除及规则授权/删除/更新
- 涉及产品: VPC，新增普通与企业安全组列表查询及详情查看接口

2026-05-21 Version: v0.0.20
- 涉及产品: VPC，新增创建、查询、更新、删除路由规则接口
- 涉及产品: VPC，新增查询路由表接口
- 涉及产品: VPC，新增主备切换接口

2026-05-20 Version: v0.0.19
- 涉及产品: VPC，支持创建、更新、释放专线网关
- 涉及产品: VPC，支持查询专线网关列表及详情
- 涉及产品: VPC，支持绑定与解绑物理专线
- 涉及产品: VPC，支持创建专线网关健康检查

2026-05-18 Version: v0.0.18
- 涉及产品: EIP，新增共享带宽的创建、释放、续费、退订、更新及扩容
- 涉及产品: EIP，新增共享带宽移入移出EIP及IP数量升级
- 涉及产品: EIP，新增带宽包创建、查询、释放及自动释放时间更新
- 涉及产品: EIP，新增共享流量包管理及DDoS基础防护阈值与攻击记录查询

2026-05-15 Version: v0.0.17
- 涉及产品: SNIC，SDK版本更新

2026-05-15 Version: v0.0.16
- 涉及产品: CFW，支持查询CFW策略列表

2026-05-14 Version: v0.0.15
- 涉及产品: VPC，SDK版本更新

2026-05-14 Version: v0.0.14
- 涉及产品: SNIC，支持查询服务网卡详情及列表
- 涉及产品: SNIC，支持创建、更新、删除服务网卡
- 涉及产品: SNIC，支持更新服务网卡企业安全组和普通安全组
- 涉及产品: SNIC，支持查询可挂载的公共服务

2026-05-12 Version: v0.0.13
- 涉及产品: Core工具包，支持文件流

2026-04-29 Version: v0.0.12
- 涉及产品: PRIVATEZONE，支持PrivateZone管理功能
- 涉及产品: PRIVATEZONE，支持解析记录的增删改查操作
- 涉及产品: PRIVATEZONE，支持VPC关联与解关联功能
- 涉及产品: PRIVATEZONE，支持解析记录状态设置

2026-04-02 Version: v0.0.11
- 涉及产品: VPC，更新接口定义
2026-03-30 Version: v0.0.10
- 涉及产品:VPC 增加网关限速规则
- 涉及产品:BLB 询价接口添加响应

2026-03-20 Version: v0.0.9
- 涉及产品:BLB 增加BLB SDK

2026-03-16 Version: v0.0.8
- 涉及产品:VPN 增加VPN SDK

2026-03-02 Version: v0.0.7
- 涉及产品:EIP 创建资源转移返回转移ID&资源ID

2026-02-26 Version: v0.0.6
- 方法名称修正

2026-02-26 Version: v0.0.5
- 支持EIP资源转移

2026-02-12 Version: v0.0.4
- 代码格式化

2026-02-10 Version: v0.0.3
- 支持EIP接口SDK

2026-02-10 Version: v0.0.2
- Init

2026-02-07 Version: v0.0.1
- Init