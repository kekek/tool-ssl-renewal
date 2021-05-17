#!/usr/bin/env bash

echo ">> build linux && scp to server"

export GOOS="linux"
./build.sh

moduleName=$(go list -m)

program=$(basename ${moduleName})

scp bin/${program} root@106.75.72.70:/root/.acme.sh/self-hook/

echo ">> publish on server:"

echo "cd /root/.acme.sh/self-hook && mv ${program}  tool-update-ssl-ucloud"