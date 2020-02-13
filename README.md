# 会员服务

只用于存放会员信息， uid 用于所有服务的用户标识。

个人用户与企业用户都作为会员，可以实现企业认证、个人认证

每个会员在不同机构下，登录账号密码不同

# 数据结构

会员账号集合(ms_member_info)
会员编号_id、所属机构会员ID、登录账号、登录密码(md5值)、手机号、邮箱、昵称、头像、用户签名、注册时间、更新时间、最后登录时间、状态（1=正常 2=禁用 9=注销）、账号类型（1=个人 2=机构）、实名认证(1=是 0=否)、推荐人推荐码、当前用户推荐码

扩展字段用户绑定微信
微信appid、微信openid


扩展字段个人会员信息(user)
会员编号_id、真实姓名、证件类型、证件号码、银行卡数组、证件照正面、证件照反面、联系地址

扩展字段机构会员信息(org)
会员编号_id、证件号码、证件照片、机构名称、负责人姓名、负责人电话、联系地址

会员所属机构账户(ms_member_account)
会员编号_id、服务机构会员ID、积分账户ID、现金账户ID、彩票账户ID、短信账户ID、电子券账户ID、产品账户ID、商城ID 、推荐人推荐码、当前用户推荐码

实名认证审核记录(ms_member_audit)，存储审核信息，审核通过之后更新到会员信息集合
审核编号、会员编号_id、发行机构会员编号、需要审核的字段、审核状态(10=待审核 11=处理中  12=通过  13=拒绝 )、审核说明、提交审核时间、处理时间、审核完成时间

操作日志记录(ms_member_log)


# 接口定义

- 会员注册(参数：手机号、密码、短信验证码)
- 会员登录(参数：账号密码  响应：所属机构ID、token )
- 重置密码
- 个人会员列表
- 机构会员列表
- 批量导入会员用户，用户注册时，如果用户存在，则是重置密码





##部署说明

编译镜像：docker pull golang:1.13-alpine
运行环境：docker pull mszs/alpine:3.10

docker build -t registry.cn-hangzhou.aliyuncs.com/mszs/member:latest ./
docker push registry.cn-hangzhou.aliyuncs.com/mszs/member:latest
docker pull registry.cn-hangzhou.aliyuncs.com/mszs/member:latest

docker service update --force --update-parallelism 1 --update-delay 3s member
docker service update  --replicas 3  member

Docker service 
```
// docker镜像源
docker service create --name member --network cluster -p 39701:80 --replicas 1  -d mszs/member:latest
// 阿里云镜像源
docker service create --name member --network cluster --replicas 1 -d registry.cn-hangzhou.aliyuncs.com/mszs/member:latest
```

测试环境接口地址：http://211.152.57.29:39701/api/v2/graphql  


Docker service 
```
docker service create --name member --network cluster --replicas 1  -d hub.unionlive.com/member:latest
docker service update --force --update-parallelism 1 --update-delay 3s --image hub.unionlive.com/member:latest member

```



# Change Log 


## v1.0.0
    [Release 2019-09-11 ]
- [Feature] 创建项目 














