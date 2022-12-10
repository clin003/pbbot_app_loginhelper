pbbot_app_loginhelper

# pbbot_app_loginhelper
## 源码托管
[gitee baicailin/pbbot_app_loginhelper](https://gitee.com/lyhuilin/pbbot_app_loginhelper)
[github baicailin/pbbot_app_loginhelper](https://github.com/clin003/pbbot_app_loginhelper)

## 镜像地址
[docker baicailin/pbbot_app_loginhelper](https://hub.docker.com/r/baicailin/pbbot_app_loginhelper)
```
docker pull baicailin/pbbot_app_loginhelper
```
## 使用说明

pbbot_app_loginhelper 可用于 [GMC](https://hub.docker.com/r/baicailin/gmc) 的多账户快速登录，需配合 GMC 使用 （该程序用于GMC非首次登录）

## 使用方法

1、首次启动一闪而过生成配置文件

2、添加服务配置模版(配置文件存储在 ./conf/ 目录中)

```yaml
runmode: debug
# 服务绑定端口
addr: :8081
# 服务器的ip:port
url: http://127.0.0.1:8081
# 自检服务重试的次数
max_ping_count: 10
#TLS服务域名(可选配置)
autotls_domain:  api.lyhuilin.com
#打开Tls配置，true开启TLS服务，false会关闭TLS服务。 (可选配置)
autotls_enable:  false
# tls证书key文件路径
tls_key_file: "./cert/apiclient.key"
# tls证书pem文件路径
tls_pem_file: "./cert/apiclient.pem"
#
## openapi的日志相关配置说明。
# writers: 输出位置，有2个可选项：file,stdout。选择file会将日志记录到logger_file指定的日志文件中，选择stdout会将日志输出到标准输出，当然也可以两者同时选择
# logger_level: 日志级别，DEBUG, INFO, WARN, ERROR, FATAL
# logger_file: 日志文件
# log_format_text: 日志的输出格式，json或者plaintext，true会输出成json格式，false会输出成非json格式
# rollingPolicy: rotate依据，可选的有：daily, size。如果选daily则根据天进行转存，如果是size则根据大小进行转存
# log_rotate_date: rotate转存时间，配合rollingPolicy: daily使用
# log_rotate_size: rotate转存大小，配合rollingPolicy: size使用
# log_backup_count:当日志文件达到转存标准时，log系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数。
# openapi的日志相关配置。
log:
  writers: file,stdout
  logger_level: INFO
  logger_file: log/openapi.log
  log_format_text: false
  rollingPolicy: size
  log_rotate_date: 1
  log_rotate_size: 1
  log_backup_count: 7
```

3、修改配置文件 loginhelper.json
```json
{
	"debug": false,
	"gmc_server_url": "http://127.0.0.1:9000",
	"CheckSleep": 600000000000,
	"Logins": [
		{
			"bot_id": 1234567890,
			"password": "123456ab",
			"device_seed": 1234567890,
			"client_protocol": 1
		},
		{
			"bot_id": 1234567890,
			"password": "123456ab",
			"device_seed": 1234567890,
			"client_protocol": 1
		}
	]
}
```

4、重新启动 pbbot_app_loginhelper

5、在GMC中添加 pbbot_app_loginhelper 的插件服务地址（默认配置为 ws://127.0.0.1:8081/ws/rq/ ）

##	Docker 使用方法
[baicailin/pbbot_app_loginhelper](https://hub.docker.com/r/baicailin/pbbot_app_loginhelper)
```
docker pull baicailin/pbbot_app_loginhelper
```
配置方法参考上面的。

### 独立插件模式
在app/docker-compose.yml中写入以下内容：
```yaml
version: "3"
services:
  pbbot_app_loginhelper:
    image:  baicailin/pbbot_app_loginhelper
    restart: always
    # build: ./..
    command: /app/pbbot_app_loginhelper -c=/app/conf/config.yaml
    volumes:
      - ./conf:/app/conf
      - ./autocert:/app/autocert
      - ./log:/app/log
      # - /tmp:/app/log
    # environment:
    #   TZ: Asia/Shanghai
    ports:
      - "8081:8081"
```

###	GMC+pbbot_app_loginhelper插件模式

```yaml
version: "3"
services:
  pbbot_app_loginhelper:
    image:  baicailin/pbbot_app_loginhelper
    restart: always
    # build: ./..
    command: /app/pbbot_app_loginhelper -c=/app/conf/config.yaml
    volumes:
      - ./conf:/app/conf
      - ./autocert:/app/autocert
      #- ./log:/app/log
      - /tmp:/app/log
    environment:
      TZ: Asia/Shanghai
    networks:
      - gmc_net
      # - gmc_default
    # ports:
    #   - "7005:8081"
  gmc:
    image:  baicailin/gmc
    restart: always
    # build: ./..
    command: gmc
    volumes:
      - ./device:/data/device
      - ./plugins:/data/plugins
    environment:
      TZ: Asia/Shanghai
    networks:
      - gmc_net      
    ports:
      - "7000:9000"

# docker network create gmc_net
networks:
  gmc_net:
    driver: bridge
```

使用容器共享网络后的GMC插件地址为：
```
 ws://pbbot_app_loginhelper:8081
```
在loginhelper.json中GMC的地址为：
```
http://gmc:9000
```

### 相关项目

[GMC](https://github.com/2mf8/Go-Mirai-Client/releases)
[GMCLoginHelper](https://github.com/2mf8/GMCLoginHelper)