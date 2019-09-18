<h1 align="center">rest-api </h1>

> 基于golang 的 RESTful API 脚手架

## 实现功能

- 基于 [gin](https://github.com/gin-gonic/gin) 封装
- 数据库**ORM** 使用 [ gorm](https://github.com/jinzhu/gorm)
- 使用 [spf13/viper](https://github.com/spf13/viper) 读取**配置文件**`conf/config.yaml`
- 使用 [lexkong/log](https://github.com/lexkong/log) 记录和管理 **API 日志**
- 自定义错误码 `pkg/errno`
- **JWT** API 身份验证
- 支持 **HTTPS**
- [Swagger](https://github.com/swaggo/swag) 文档
- ...

## 使用

1. 安装与编译

```sh
git clone https://github.com/imyhui/rest-api.git
make
```

2. 修改配置

   `conf/config.yaml`

   ```yaml
   db:
     name: db_name
     addr: 127.0.0.1:3306
     username: db_username
     password: db_password
   ```

3. 数据库创建 `db.sql`

4. 运行

   ```sh
   ./admin.sh start
   ```
