# Todo 待办事项 边开发边思考 开发规划
- [ ] 中间件
- [x] validtaion 翻译
- [x] logger 日志处理
- [x] PostgreSQL 错误处理
- [x] 模拟短信验证码
- [x] JWT Token
- [ ] HTTPS
- [ ] 发邮件
- [ ] 登录功能
- [ ] 刷新Token
- [ ] 用户模块
- [ ] 电影院大厅和座位接口开发
- [ ] 电影和showtimes接口开发
- [ ] 电影票模块
- [ ] 。。。


### 后台管理API ｜ 用户模块
- 功能规划 -> 用户模块
1. 用户不需要注册，默认添加几个用户，并设置为管理员
1. 管理员登录后台暂定使用验证码登录，登录后可设置密码
1. 管理员登录后，生成 jwt token 和刷新 token 保持一定时间内免登录
1. 管理员登录后可以修改自己的用户信息
1. 管理员查看用户列表

### 创建基础数据
限于篇幅，对于一些基础表直接上sql添加基础数据，tb_theaters、tb_halls, 因此迁移一下基础数据；
初始化了 tb_theaters、tb_halls 表基础数据。

### tb_seats 座位表业务实现


### 记一个问题
因为数据库小部分字段设置可以为NULL的，因此在执行 sqlc generate 命令生成的代码带有sql.NullXXX类型，
然后格式化为JSON格式传输给前端，格式化后的结构并不友好，如：
``` json
"location": {
    "String": "广州市番禺区祈福缤纷世界",
    "Valid": true
},
```
参考：https://github.com/kyleconroy/sqlc/pull/1571

这是是一个功能性缺陷问题，还有待sqlc解决，已有PR提交。。。


### 构思接口
基础表的CRUD

### 影院模块构思

电影院有大厅，每个大厅有对应的座位，先假设每一个大厅都是
10 x 10 的座位，那么即得100个位置。

因此创建座位的接口入参
``` json
[
    {"row_number": 1, "col_number": 1, "status": "A", "hall_id": 1},
    {"row_number": 11, "col_number": 11, "status": "A", "hall_id": 1},
]
```
sqlc 如何实现批量插入
参考
https://medium.com/@amoghagarwal/insert-optimisations-in-golang-26884b183b35


### 
