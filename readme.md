
### 启动项目
```go
# 运行项目
go run main.go
# gin框架在debug模式下的输出
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/acs                  --> GinCasbin/routers.api.func1 (3 handlers)
[GIN-debug] DELETE /api/acs/:id              --> GinCasbin/routers.api.func2 (3 handlers)
[GIN-debug] POST   /api/routers              --> GinCasbin/routers.api.func3 (4 handlers)
[GIN-debug] POST   /api/v1/user              --> GinCasbin/routers.api.func4 (4 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> GinCasbin/routers.api.func5 (4 handlers)
[GIN-debug] PUT    /api/v1/user/:id          --> GinCasbin/routers.api.func6 (4 handlers)
[GIN-debug] GET    /api/v1/user/:id          --> GinCasbin/routers.api.func7 (4 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

### 测试
```bash
# 访问接口
# 参数缺失
curl -X POST http://127.0.0.1:8080/api/routers
{"code":-1,"message":"header miss userName","data":null}


# 无访问权限
curl -X POST -H "userName:tom" http://127.0.0.1:8080/api/routers
{"code":-1,"message":"access denied","data":null}


# 添加一条规则(代码中是模拟数据)
curl -X POST http://127.0.0.1:8080/api/acs
{"code":200,"message":"success","data":"add success"}

# 再次访问(有访问权限,可以访问)
curl -X POST -H "userName:tom" http://127.0.0.1:8080/api/routers
{
    "code":200,
    "message":"success",
    "data":[
        {
            "method":"POST",
            "path":"/api/acs"
        },
        {
            "method":"POST",
            "path":"/api/routers"
        },
        {
            "method":"POST",
            "path":"/api/v1/user"
        },
        {
            "method":"DELETE",
            "path":"/api/acs/:id"
        },
        {
            "method":"DELETE",
            "path":"/api/v1/user/:id"
        },
        {
            "method":"PUT",
            "path":"/api/v1/user/:id"
        },
        {
            "method":"GET",
            "path":"/api/v1/user/:id"
        }
    ]
}

# 直接向数据库添加几条Policy策略
INSERT INTO `ginCasbin`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user', 'POST', NULL, NULL, NULL);
INSERT INTO `ginCasbin`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/:id', 'GET', NULL, NULL, NULL);
INSERT INTO `ginCasbin`.`casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'admin', '/api/v1/user/:id', 'PUT', NULL, NULL, NULL);

#再测试
## 添加接口
curl -X POST -H "userName:admin" http://127.0.0.1:8080/api/v1/user
{"code":200,"message":"user add success"}
## 查询接口
curl -X GET -H "userName:admin" http://127.0.0.1:8080/api/v1/user/99
{"code":200,"message":"user Get success 99"}
## 更新接口
curl -X PUT -H "userName:admin" http://127.0.0.1:8080/api/v1/user/199
{"code":200,"message":"user update success 199"}
## 删除接口(没有分配访问权限)
curl -X DELETE -H "userName:admin" http://127.0.0.1:8080/api/v1/user/299
{"code":-1,"message":"access denied","data":null}

```

### casbin_rule.sql
上述 Demo 的SQL文件如下(该表是gorm适配器自动创建的)
```sql
-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', 'zhangsan', '/api/v1/ping', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'coder', '/api/v2/user/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'coder', '/api/v2/routers', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user', 'POST', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/:id', 'GET', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'admin', '/api/v1/user/:id', 'PUT', null, null, null);
INSERT INTO `casbin_rule` VALUES ('p', 'tom', '/api/routers', 'POST', '', '', '');

```