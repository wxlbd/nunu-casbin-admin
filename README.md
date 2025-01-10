# Nunu-Casbin-Admin

一个基于 Go 语言开发的现代化后台管理系统，集成了 RBAC 权限管理、JWT 认证等功能。

## 特性

- 基于 Casbin 的 RBAC 权限管理
- JWT Token 认证和自动续期
- 基于 Gin 的 RESTful API
- 统一的错误处理和响应格式
- 支持多种数据库（MySQL、PostgreSQL、SQLite）
- 完整的用户、角色、菜单管理
- 优雅的项目结构和代码组织

## 技术栈

- **框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **权限**: [Casbin](https://casbin.org/)
- **认证**: [JWT](https://github.com/golang-jwt/jwt)
- **缓存**: [Redis](https://github.com/redis/go-redis)
- **配置**: [Viper](https://github.com/spf13/viper)
- **日志**: [Zap](https://github.com/uber-go/zap)
- **依赖注入**: [Wire](https://github.com/google/wire)

## 项目结构

```plaintext
.
├── cmd/                    # 应用程序入口
│   └── server/             # HTTP 服务器
├── configs/                # 配置文件目录
│   ├── config.yaml         # 应用配置
│   └── casbin/             # Casbin 配置
├── internal/               # 内部代码
│   ├── dto/                # 数据传输对象
│   ├── handler/            # HTTP 处理器
│   │   ├── request/        # 请求结构
│   │   └── response/       # 响应结构
│   ├── middleware/         # 中间件
│   ├── model/              # 数据模型
│   ├── repository/         # 数据访问层
│   └── service/            # 业务逻辑层
├── pkg/                    # 公共包
│   ├── config/             # 配置管理
│   ├── helper/             # 辅助工具
│   ├── http/               # HTTP 客户端
│   ├── jwtx/               # JWT 工具
│   ├── log/                # 日志工具
│   └── utils/              # 通用工具
└── storage/                # 存储目录
    └── logs/               # 日志文件
```

## 核心功能

### 用户管理
- 用户 CRUD
- 密码加密存储
- 登录历史记录
- 用户状态管理

### 角色管理
- 角色 CRUD
- 角色-菜单分配
- 角色-API权限控制

### 菜单管理
- 菜单 CRUD
- 菜单树形结构
- 按钮级权限控制

### 权限控制
- 基于 Casbin 的 RBAC
- 细粒度的 API 权限控制
- 动态权限分配

## 快速开始

### 环境要求

- Go 1.20+
- MySQL 5.7+ / PostgreSQL 10+ / SQLite 3
- Redis 6.0+

### 安装

1. 克隆项目
```bash
git clone https://github.com/yourusername/nunu-casbin-admin.git
cd nunu-casbin-admin
```

2. 安装依赖
```bash
go mod download
```

3. 配置数据库
```bash
# 编辑 configs/config.yaml
cp configs/config.yaml.example configs/config.yaml
```

4. 初始化数据库
```bash
# 导入数据库结构
mysql -u root -p your_database < mineadmin.sql
```

5. 运行项目
```bash
go run cmd/server/main.go
```

## API 文档

### 认证相关
- POST /admin/v1/login - 用户登录
- POST /admin/v1/refresh-token - 刷新令牌

### 用户管理
- GET /admin/v1/user - 获取用户列表
- POST /admin/v1/user - 创建用户
- PUT /admin/v1/user/:id - 更新用户
- DELETE /admin/v1/user/:id - 删除用户

### 角色管理
- GET /admin/v1/role - 获取角色列表
- POST /admin/v1/role - 创建角色
- PUT /admin/v1/role/:id - 更新角色
- DELETE /admin/v1/role/:id - 删除角色

### 菜单管理
- GET /admin/v1/menu/tree - 获取菜单树
- POST /admin/v1/menu - 创建菜单
- PUT /admin/v1/menu/:id - 更新菜单
- DELETE /admin/v1/menu/:id - 删除菜单

## 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 提交 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 联系方式

- 作者：[wxl]
- 邮箱：[gopher095@gmail.com]
- 项目地址：[https://github.com/wxlbd/nunu-casbin-admin]