#minipro
config.yaml example
```yaml
name: "minipro"
mode: "dev"
port: 8080
version: "v0.1.3"
start_time: "2022-05-11"
machine_id: 1


auth:
  jwt_expire: 8760

log:
  level: "debug"
  filename: "web_log"
  max_size: 200
  max_age: 30
  max_backups: 7
mysql:
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "xxx"
  dbname: "xxx"
  max_conns: 200
  max_idle_conns: 50
redis:
  host: "127.0.0.1"
  port: 6379
  db: 0
  password: ""
  poolsize: 1

```
