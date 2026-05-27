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