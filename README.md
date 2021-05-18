## 说明
    
let's encrypt 证书只有 3个月的期限， 部署需要每三个月更新一次。 好在let'enctypt 可以通过API，保证证书自动续期，证书续期后，还需要安装部署到其他的平台， 本项目即是一个更新证书到ucloud云的工具

将编译好的二进制文件和配置文件发送到acme 根目录后， 使用 acme.sh 自带的 renew-hook， 每次证书续期后， 同时自动更新证书到平台


## 程序使用

两个功能 : 创建和绑定ssl证书到ucloud的负载均衡

- 创建ssl证书
- 绑定新到ulb的虚拟服务上


```bash 
Usage:
  tool-ssl-renewal [command]

Available Commands:
  binding     给ucloud负载均衡下的虚拟服务，绑定证书
  create      create a ssl certificate to ULB
  help        Help about any command
  test        测试命令的运行情况

Flags:
  -c, --config string   config file (default is $HOME/.tool-test.yaml)
  -h, --help            help for tool-ssl-renewal

Use "tool-ssl-renewal [command] --help" for more information about a command.
```

## 配置文件格式 
```bash
ProjectId = "porject id"

SSLConfigPath = "acme 生成相关域名的路径"

[UCloudCredential]
PublicKey = "ucloud PublicKey"
PrivateKey = "ucloud PrivateKey"

```

## hook 自动更新

let's encrypt 颁发证书时会接收 renew-hook 参数，假如第一次颁发时没有假如该参数， 可以重新颁发一次证书，只要域名参数完全一致即可 

我们用shell script [script/renew-hook.sh] 来作为 renew-hook, 以便后续修改

```bash 
acme.sh --issue -d ilaiyang.com.cn --renew-hook "path/to/renew-hook.sh &>> path/to/hook.log "
```


## v2

- create  在ucloud生成一个新的证书,返回新创建的sslID值
- bind 绑定命令， 接收新的sslID， v-server-ID， ulb-ID， 创建新的绑定


## 参考文档

- https://github.com/acmesh-official/acme.sh/wiki/Using-pre-hook-post-hook-renew-hook-reloadcmd
- 