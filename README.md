# tiktokProject

基于postgres+redis+gorm+gin+jwt搭建的简易tiktok抖声服务端

# 技术选型

![选型](https://internal-api-drive-stream.feishu.cn/space/api/box/stream/download/v2/cover/boxcn7LA3yDTtT5z10yRb7RlWpb/?fallback_source=1&height=1280&mount_node_token=LYGSdQweMoaCWqxYv2xc43e8nZc&mount_point=docx_image&policy=equal&width=1280)

# 架构

![架构](https://internal-api-drive-stream.feishu.cn/space/api/box/stream/download/v2/cover/boxcn5wOvB0XCL2kjSYqRgxQpUf/?fallback_source=1&height=1280&mount_node_token=JQkWdAKAEooSccx5slKc3S4fnqr&mount_point=docx_image&policy=equal&width=1280)

按照上方架构图进行设计。暂时只考虑了平流情况。对于突增流量可以在redis中添加对应的短期缓存。写入请求可以使用redis作为消息队列进行缓存，从而缓解数据库写入压力。

# 数据库

![ER](blob:https://yrswyur8nn.feishu.cn/3bd32913-d1fe-44cd-af9e-5cabd1f61e64)

# Redis

![结构](blob:https://yrswyur8nn.feishu.cn/d2d6ec6e-c3b5-4141-b006-251d08e84ea4)

0号数据库用于存储user信息和video信息。1号数据存储和video有关系的数据。2号数据库存储和user有关系的数据。

![索引结构关系](blob:https://yrswyur8nn.feishu.cn/13225017-8a15-4613-a0d2-89bd0cd024b8)

类似的`user_favorite_{user_id}, user_follow_{user_id}. user_follower_{user_id}`也是如此索引方法。他们不保存具体的对象信息而是保存对象的key。通过这个key到指定的数据库再去获取对象的json。可以提高系统的灵活性，降低操作的次数。

# 代码结构

```
.
├── controller
│   ├── comment.go // comment api handle                
│   ├── common.go // 保存基础结构，以及基础结构和数据库结构之间的转换方法                 
│   ├── demo_data.go
│   ├── favorite.go // favorite api handle
│   ├── feed.go // video stream api handle
│   ├── message.go // message api handle
│   ├── publish.go // pulish api and user publish video api handle
│   ├── relation.go // relation api handle
│   └── user.go // user api handle
├── docker-compose.yml // docker-compose配置文件
├── Dockerfile_air // 开发环境构建使用的Dockerfile
├── Dockerfile_postgre // 构建Postgres with rw-redis_fdw使用的Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── log.txt // 系统运行产生的日志文件
├── main.go // 程序入口
├── Makefile
├── middleware
│   ├── jwt
│   │   └── jwt.go // jwt中间件实现
│   └── redis
│       └── redis.go // redis中间件实现
├── public // 用于保存上传的视频和生成的封面图片
├── README.md
├── repository
│   ├── chat.go // message表与DAO
│   ├── comment.go // comment表与DAO
│   ├── db_init.go // 数据库初始化流程操作
│   ├── favorite.go // favorite表与DAO
│   ├── relation.go // relation表与DAO
│   ├── sql
│   │   └── redis_fdw_init.sql // postgres 初始化redis_fdw，外表建立以及对应触发器建立
│   ├── transcation.go // 多表事务
│   ├── user.go // user表与DAO
│   └── video.go // video表与DAO
├── router.go // 路由配置
├── service
│   ├── chat
│   │   └── chat.go // 向上提供的关于message的服务
│   ├── comment
│   │   └── comment.go // 向上提供的关于comment的服务
│   ├── favorite
│   │   └── favorite.go // 向上提供的关于favorite的服务
│   ├── message.go
│   ├── relation
│   │   └── relation.go // 向上提供的关于relation的服务
│   ├── user
│   │   └── user.go // 向上提供的关于user的服务
│   └── video
│       └── video.go // 向上提供的关于video的服务
├── test
│   ├── base_api_test.go
│   ├── common.go
│   ├── interact_api_test.go
│   ├── message_server_test.go
│   ├── social_api_test.go
│   └── test.md
└── utils
    ├── cover.go // 提供生成视频封面的功能
    └── log.go // 提供日志flag初始化的功能
```

# 测试

测试结果见`./test/test.md`

api测试全部通过。