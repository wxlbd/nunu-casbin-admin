<div align="center">
    <h1>🚀 Gin-Casbin-Admin</h1>
    <p>一个基于 Go 语言开发的现代化后台管理系统</p>
    <p>
        <a href="https://golang.org/">
            <img src="https://img.shields.io/badge/Go-1.20%2B-blue" alt="Go version">
        </a>
        <a href="https://github.com/gin-gonic/gin">
            <img src="https://img.shields.io/badge/Gin-1.10.0-brightgreen" alt="Gin version">
        </a>
        <a href="https://gorm.io">
            <img src="https://img.shields.io/badge/GORM-1.25.12-red" alt="GORM version">
        </a>
        <a href="https://github.com/casbin/casbin">
            <img src="https://img.shields.io/badge/Casbin-2.103.0-orange" alt="Casbin version">
        </a>
        <a href="https://github.com/wxlbd/gin-casbin-admin/blob/main/LICENSE">
            <img src="https://img.shields.io/github/license/wxlbd/gin-casbin-admin" alt="License">
        </a>
    </p>
</div>

## ✨ 特性

- 🔐 **RBAC 权限管理**: 基于 Casbin 的细粒度权限控制
- 🔑 **JWT 认证**: 支持 Token 自动续期
- 🎯 **RESTful API**: 规范的接口设计
- 📝 **日志系统**: 基于 Zap 的高性能日志
- 🔄 **事务支持**: 数据库操作的完整性保证
- 🛡️ **统一错误处理**: 规范的错误响应
- 📦 **模块化设计**: 清晰的项目结构

## 🚀 快速开始

### 环境要求

- Go 1.20+
- MySQL 5.7+ / PostgreSQL 10+
- Redis 6.0+

### 安装

```bash
# 克隆项目
git clone https://github.com/wxlbd/gin-casbin-admin.git

# 安装依赖
go mod download

# 配置环境
cp configs/config.yaml.example configs/config.yaml

# 运行项目
go run cmd/server/main.go
```

## 📚 文档

详细文档请查看 [Wiki](https://github.com/wxlbd/gin-casbin-admin/wiki)

## 🔨 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **权限**: [Casbin](https://casbin.org/)
- **缓存**: [Redis](https://redis.io/)
- **配置**: [Viper](https://github.com/spf13/viper)
- **日志**: [Zap](https://github.com/uber-go/zap)

## 📦 项目结构

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

## 🤝 贡献

欢迎提交 PR 和 Issue！

## 📄 许可证

[MIT License](LICENSE)

## 📧 联系

- 作者：wxl
- 邮箱：gopher095@gmail.com

---
<p align="center">如果这个项目对你有帮助，请给一个 ⭐️</p>