#!/usr/bin/env bash

echo ">> build linux && scp to server"

workhost="106.75.72.70"
workdir="/root/.acme.sh/self-hook/"

export GOOS="linux"
./build.sh

moduleName=$(go list -m)

program=$(basename ${moduleName})

suffix="$(go env GOOS)-$(go env GOARCH)"

target="${program}-${suffix}"

scp bin/${program} root@${workhost}:"${workdir}${target}"

echo ">> publish on server:"
#
echo "cd ${workdir} && mv ${target}  $program"