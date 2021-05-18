#!/usr/bin/env bash

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
echo "${toolPath}  binding -c ${configFilePath} -l ${sslId} -b \"ulb-cqmpcf25\" -s \"vserver-1qgzajvb\""
#${toolPath}  binding -c ${configFilePath} -l ${sslId} -b "ulb-cqmpcf25" -s "vserver-1qgzajvb"

echo "comet 绑定新证书"
echo "${toolPath}  binding -c ${configFilePath} -l ${sslId} -b \"ulb-cwwmvypr\" -s \"vserver-2mrxtihd\""
#${toolPath}  binding -c ${configFilePath} -l ${sslId} -b "ulb-cwwmvypr" -s "vserver-2mrxtihd"

echo "$0 job done"