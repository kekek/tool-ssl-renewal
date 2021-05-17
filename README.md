## 说明
    
let's encrypt 证书只有 3个月的期限， 部署需要每三个月更新一次。 好在let'enctypt 可以通过API，保证证书自动续期，证书续期后，还需要安装部署到其他的平台， 本项目即是一个更新证书到ucloud云的工具

将编译好的二进制文件和配置文件发送到acme 根目录后， 使用 acme.sh 自带的 renew-hook， 每次证书续期后， 同时自动更新证书到平台


## 运行

```bash 

./tool-ssl-renewal -c run.toml 

Usage of ./tool-ssl-renewal:
  -c string
    	path to conf that create by acme.sh (default "data/run.toml")
  -v	show build version for the program

```

## 配置文件
```bash
VServerId = "vserver id"
UlbId = "ulb id"
ProjectId = "porject id"

SSLConfigPath = "acme 生成相关域名的路径"

[UCloudCredential]
PublicKey = "ucloud PublicKey"
PrivateKey = "ucloud PrivateKey"

```

## 更新hook

let's encrypt 获取证书后， 会自动更新，自动更新

```bash 

acme.sh --deploy -d ilaiyang.com.cn --renew-hook "path/to/tool-ssl-renewal -c path/to/conf.toml"

```
