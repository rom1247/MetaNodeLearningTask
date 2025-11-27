
## 项目介绍
基于Gin 框架和 GORM 库的博客系统的后端

## 使用说明
```
- golang版本 >= 1.25.4
- IDE推荐：Goland
```
初始化
```bash

# 克隆项目
git clone git@github.com:rom1247/MetaNodeLearningTask.git

# 使用 go mod 并安装go依赖包
go mod midy

# 使用wire 生成依赖注入
cd backend\phase_two\task_four\internal\bootstrap
wire
# 使用swag 生成文档
cd /backend/phase_two/task_four/
swag init -g ./cmd/server/main.go -o ./docs
# 运行
cd /backend/phase_two/task_four/cmd/server
go run

```
## 技术选型

- 后端：用 `gin` 搭建基础restful风格的API
- 数据库：采用 `gorm` 实现对数据库的基本操作。
- API文档：使用`Swagger`构建自动化文档。
- 依赖注入：使用`wire` 生成依赖注入
- 日志：使用 `zap`实现日志记录。

## 目录结构
```shell
├── cmd
│   └── server                  主程序人口
├── cofnigs                     配置目录
├── docs                        swagger文档目录
├── internal
│   ├── app                     应用服务
│   ├── bootstrap               初始化
│   └── domain                  领域层
│       ├── model               领域模型
│       ├── repository          仓储接口
│       └── service             领域服务
│   └── infrastructure          基础设置
│       ├── cache               缓存
│       ├── middeware           中间件
│       └── persistence         持久层实现
│   └── interfaces
│       └── http
│           ├── controller      控制器
│           └── routes          路由
├── pkg                         工具

```
## 主要功能
- 用户注册、登录和JWT认证
- 文章的CRUD操作
- 发表文章评论
- 统一的错误处理
- 完整的日志记录

