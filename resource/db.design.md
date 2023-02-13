# 数据库设计

这里是电影购票系统数据库设计环节，现在本地设计系统需要的字段，最后在 [dbdiagram](https://dbdiagram.io/d/63e90c55296d97641d804b10) 网站完成设计并生成SQL文件。

按照日常看电影的流程出发，假如想看一场电影，首先了解的是电影，进而选择合适的电影院，然后开始购票
流程，购票就需要选择电影播放时间短，也就是选择场次。

一切顺利后，购票也成功了，欢喜地来到电影院，取票后找到对应的放映大厅，最后找到属于的座位。

从上面的流程来看，必须三个硬性条件，电影院/剧院->大厅->座位，因此就有剧院表、大厅表、座位表。

从买车票的角度想，买的车票就是一个坐/睡的位置，因此看电影，买的也是一个座位罢了。

所以有了一个电影表，电影再根据大厅和座位编排场次，因此就有场次表。

具体表设计如下：

### 用户表设计（tb_users）
- id           (主键)
- role_id      (外键，角色ID)
- phone_number (非空、如果微信授权了，则用微信手机号)
- password    （可空，密码）
- name        （非空、随机生成默认值，如果微信授权了，则用微信昵称）
- avatar       (可空、随机生成默认值，如果微信授权了，则用微信昵称)
- openid       (可空、微信openid)
- unionid      (可空、微信unionid，同一主单体unionid相同)
- created_at
- updated_at
- deleted_at

### 角色表设计 (tb_roles)
- id            (主键)
- name          (非空、角色名)
- code          (非空、角色唯一code)
- description   (可空、角色描述)
- created_at
- updated_at
- deleted_at

<p style="font-size: 15px; color: #2563eb;">角色默认先设置两个吧</p>

| Name        | Code        | Description | 
| ----------- | ----------- | ----------- |
| Admin       | ADMIN       |  管理员，管理系统
| Consumer    | CONSUMER    |  消费者，可订票

### 电影院表/剧院（tb_theaters）
- id            (主键)
- name          (非空、角色名)
- location      (可空、位置，JSON 结构{description:"", lat: "", long:""})
- created_at
- updated_at
- deleted_at

### 电影院大厅（tb_halls）
- id            (主键)
- theater_id    (非空、剧院ID)
- name          (非空、大厅名；如：3号厅)
- total_seats   (可空、改大厅里有几个座位)
- created_at
- updated_at
- deleted_at

### 座位表（tb_seats）
- id            (主键)
- hall_id       (非空、外键，大厅ID)
- col_number    (非空、列号，排，如：6排)
- row_number    (非空、行号，座，如：6座)
- status        (非空、状态: A、B、N, 分别是: 可用、被预订、损坏)

### 电影 （tb_movies）
- id            (主键)
- title         (非空、电影名)
- release_date  (非空、上映/发布日期)
- director      (非空、导演)
- poster        (非空、 海报)
- duration      (非空、时长，单位: 分钟)
- genre         (可空、类型；如: 喜剧)
- star          (可空、主演)
- description   (可空、描述)
- created_at
- updated_at
- deleted_at

### 场次表（tb_showtimes）
- id            (主键)
- movie_id      (非空、外键，电影ID)
- hall_id       (非空、外键，大厅ID)
- start_time    (非空、放映日期)
- end_time      (非空、结束日期)
- created_at
- updated_at
- deleted_at

### 电影票 （tb_tickets）
- id            (主键)
- user_id       (可空、外键，用户ID，如果存在user_id，说明该票已被预订)
- showtime_id   (非空、外键，场次ID)
- seat_id       (非空、外键，座位ID)
- price         (非空、单价，单位: 分)
- booking_date  (可空、下订日期，被预订时设置时间)
- payment_status(可空、支付状态:Y/N)
- created_at
- updated_at
- deleted_at

### 付款表（tb_payments）
- id             (主键)
- user_id        (非空、外键，用户ID)
- ticket_id      (非空、外键，电影票ID)
- NumberOfSeats  (非空、预订了几个位置)
- payment_date   (非空、支付日期)
- payment_method (非空、支付方式)
- payment_amount (非空、支付总额)
- created_at     
- updated_at
- deleted_at

Payments 表没有设置子表，是因为可以通过user_id和ticket_id查出该预订的明细。


