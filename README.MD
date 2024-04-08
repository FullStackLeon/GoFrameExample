##### 框架信息

- 框架名称: GoFrame
- 框架文档: https://goframe.org/pages/viewpage.action?pageId=1114399

##### 实现功能

- 简易的注册和登录
- 登录成功后生成JWT Token并缓存在Redis
- 多角色路由控制，普通用户、Auth用户和管理员用户访问不同的API接口
- 用户数据表设计和读写

##### 数据表说明

- 数据库名: goframe
- 用户表名: users
- 存储位置: manifest/sql/users.sql

##### Postman测试

- 测试集合存储位置: manifest/postman/Goframe 示例API.postman_collection.json

##### 登录和权限校验实现思路

- 用户成功登录后生成JWT Token存储于Redis，键名为SessionId的值,并设置SessionId Cookie到客户端,SessionId在Redis中存储示例:

  ```markdown
  get SessionKeybv6ZNPFeK+/drS3oU8j5og==
  {\"JwtToken\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X3R5cGUiOiJhZG1pbiIsImVtYWlsIjoiMTAyQGdtYWlsLmNvbSIsImV4cCI6MTcxMjYzMzQ2NywidXNlcm5hbWUiOiJqYWNrMiJ9.zw8BOfIyFaVrkbsZ22HLCume6CA7G_yINsKcnkV0HXY\"}¡
  ```
- 多角色路由控制通过Map配置，配置存储位置：internal/consts/common.go

  > 路由控制的配置方案是临时的，生产可根据权限系统自动生成存储于缓存中
  >
- 设计两个中间件LoginCheck和RouterCheck，存储位置：internal/logic/middleware/middleware.go

  > LoginCheck 负责JWT Token解析，解析成功后将email和account_type设置到req参数方便Controller层使用
  > RouterCheck 负责多角色路由权限检查¡，普通用户只能访问普通用户的API，访问其他角色API会提示Forbidden

##### 项目本地运行步骤
- git clone https://github.com/FullStackLeon/GoFrameExample.git
- 创建数据库和数据表，参考上面**数据表说明**部分
- 在项目根路径GoFrameExample下，执行 go run main.go
