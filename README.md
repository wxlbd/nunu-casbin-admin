# Gin-Casbin-Admin

一个基于 Go 语言开发的现代化后台管理系统，集成了 RBAC 权限管理、JWT 认证等功能。

## 核心特性

- **权限管理**
  - 基于 Casbin 的 RBAC 权限控制
  - 动态权限分配
  - 细粒度的 API 权限管理
  - 按钮级别的权限控制

- **用户认证**
  - JWT Token 认证
  - Token 自动续期
  - 多端登录控制
  - 密码加密存储

- **系统功能**
  - RESTful API 设计
  - 统一的错误处理
  - 请求响应日志记录
  - 数据库事务支持

- **项目特点**
  - 清晰的项目结构
  - 完善的日志系统
  - 统一的响应格式
  - 多数据库支持

## 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin) - 高性能 Web 框架
- **ORM**: [GORM](https://gorm.io/) - 优秀的 ORM 库
- **权限**: [Casbin](https://casbin.org/) - 灵活的访问控制框架
- **认证**: [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token
- **缓存**: [Redis](https://github.com/redis/go-redis) - 高性能缓存
- **配置**: [Viper](https://github.com/spf13/viper) - 配置管理
- **日志**: [Zap](https://github.com/uber-go/zap) - 高性能日志库
- **依赖注入**: [Wire](https://github.com/google/wire) - 编译时依赖注入

## 项目结构

```plaintext
.
├── cmd/                    # 应用程序入口
│   └── server/             # HTTP 服务器启动
├── configs/                # 配置文件
│   ├── config.yaml         # 主配置文件
│   └── casbin/             # Casbin 规则配置
├── internal/               # 内部代码
│   ├── dto/                # 数据传输对象
│   ├── handler/            # HTTP 处理器
│   ├── middleware/         # 中间件
│   ├── model/              # 数据模型
│   ├── repository/         # 数据访问层
│   ├── server/             # 服务器配置
│   └── service/            # 业务逻辑层
└── pkg/                    # 公共工具包
    ├── config/             # 配置管理
    ├── errors/             # 错误处理
    ├── ginx/               # Gin 扩展
    ├── jwtx/               # JWT 工具
    ├── log/                # 日志工具
    └── utils/              # 通用工具
```

## API 文档

### 认证接口
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/refresh-token` - 刷新令牌

### 个人中心
- `GET /api/profile` - 获取当前用户信息
- `GET /api/profile/menus` - 获取用户菜单
- `GET /api/profile/roles` - 获取当前用户角色

### 用户管理
- `GET /api/permission/user` - 获取用户列表
- `POST /api/permission/user` - 创建用户
- `PUT /api/permission/user/:id` - 更新用户
- `DELETE /api/permission/user/:ids` - 删除用户
- `GET /api/permission/user/:id` - 获取用户详情
- `GET /api/permission/user/:id/roles` - 获取用户角色
- `PATCH /api/permission/user/:id/password` - 修改用户密码
- `PUT /api/permission/user/:id/roles` - 分配用户角色

### 角色管理
- `GET /api/permission/role` - 获取角色列表
- `POST /api/permission/role` - 创建角色
- `PUT /api/permission/role/:id` - 更新角色
- `DELETE /api/permission/role/:ids` - 删除角色
- `GET /api/permission/role/:id` - 获取角色详情
- `GET /api/permission/role/:id/menus` - 获取角色菜单
- `PUT /api/permission/role/:id/menus` - 分配角色菜单

### 菜单管理
- `POST /api/permission/menu` - 创建菜单
- `PUT /api/permission/menu/:id` - 更新菜单
- `DELETE /api/permission/menu/:ids` - 删除菜单
- `GET /api/permission/menu/tree` - 获取菜单树

## 快速开始

### 环境要求
- Go 1.20+
- MySQL 5.7+ / PostgreSQL 10+ / SQLite 3
- Redis 6.0+

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/wxlbd/nunu-casbin-admin.git
cd nunu-casbin-admin
```

2. 安装依赖
```bash
go mod download
```

3. 配置环境
```bash
cp configs/config.yaml.example configs/config.yaml
# 修改配置文件中的数据库和Redis连接信息
```

4. 初始化数据库
```bash
# 导入数据库结构
mysql -u root -p your_database < scripts/schema.sql
# 导入初始数据
mysql -u root -p your_database < scripts/data.sql
```

5. 运行项目
```bash
go run cmd/server/main.go
```

## 开发指南

### 错误处理
使用统一的错误处理包 `pkg/errors`：
```go
if err != nil {
    return errors.WithMsg(errors.InvalidParam, "参数错误")
}
```

### 响应格式
使用 `pkg/ginx` 包处理响应：
```go
ginx.Success(c, data)
ginx.Error(c, code, message)
```

### 日志记录
使用 `pkg/log` 包记录日志：
```go
logger.Info("操作成功", zap.String("user", username))
logger.Error("操作失败", zap.Error(err))
```

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 联系方式

- 作者：[wxl]
- 邮箱：[gopher095@gmail.com]
- 项目地址：[https://github.com/wxlbd/nunu-casbin-admin]