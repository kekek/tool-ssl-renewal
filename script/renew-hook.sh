#!/usr/bin/env bash

## 执行绑定任务脚本
## 每个ucloud 需要更换证书的 ulb vserver 都需要执行绑定

set -e

echo "$0 is running..."

toolPath="./bin/tool-ssl-renewal"
configFilePath="data/run.toml"

echo ">> create ssl certificate"
echo $(pwd)

#sslId=`${toolPath} create -c ${configFilePath}`
sslId="ssl-vuvbzhwn"
echo $sslId

echo "www 绑定新的证书"
echo "${toolPath}  binding -c ${configFilePath} -s ${sslId} -u \"ulb-cqmpcf25\" -v \"vserver-1qgzajvb\""
#${toolPath}  binding -c ${configFilePath} -s ${sslId} -u "ulb-cqmpcf25" -v "vserver-1qgzajvb"

echo "comet 绑定新证书"
echo "${toolPath}  binding -c ${configFilePath} -s ${sslId} -u \"ulb-cwwmvypr\" -v \"vserver-2mrxtihd\""
#${toolPath}  binding -c ${configFilePath} -s ${sslId} -u "ulb-cwwmvypr" -v "vserver-2mrxtihd"

echo "$0 job done"