# tiktokProject

极简抖音app开发

```shell
go build && ./simple-demo
```

### 功能说明

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

- common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改
- 测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

Fork from [RaymondCode/simple-demo](https://tiktok-go)

### Android App

`public/app-release.apk` 设置服务端地址
为方便测试登录和注册，及修改网络请求的服务器地址，提供了退出登录和高级设置两个能力。

1. 点击退出登录会自动重启
2. 在高级设置中可以配置自己的服务端项目的前缀地址，如下配置的http://192.168.1.7:8080
   在app中访问上述某个接口时就会拼接该前缀地址，例如访问 http://192.168.1.7:8080/douyin/feed/ 拉取视频列表

在未登录情况下，双击右下角的 “我” 可以 打开高级设置

# Notic

在docker-compose.yml中的Host环境变量写入你的服务器地址。

# 评论功能

# 视频流功能

视频上传

公共视频流
