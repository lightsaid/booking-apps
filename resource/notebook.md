# Notebook 笔记

好记性不如烂笔头～

记录开发过程遇到新的知识点、疑点、难点、随手记一记。

### Go Workspaces (go 工作区)

第一次是用 Go Workspaces 模式记录一下～

项目的想法是这样的，booking-sys 目录是后台API服务，在这里实现数据的CRUD和迁移；
在 booking-gapi 实现 gRPC 服务，为小程序提供接口服务，数据的CRUD方法依赖于 booking-sys；
因此尝试使用  Go Workspaces 模式。

``` bash
# 初始化 booking-sys 模块
cd booking-sys 
go mod init github.com/lightsaid/booking-sys

# 创建 booking-gapi 项目并初始化模块
mkdir booking-gapi && cd booking-gapi
go mod init github.com/lightsaid/booking-gapi

# 将 booking-apps 初始化工作空间
cd booking-apps
# 将 booking-sys booking-gapi 两个文件夹定义为work的module
go work init booking-sys booking-gapi

# 如果后面添加其他模块，通过以下命令添加
go work use ./xxx
```
