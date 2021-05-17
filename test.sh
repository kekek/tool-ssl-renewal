#!/usr/bin/env bash

set +e

echo ">> ./build.sh"
./build.sh
echo

echo ">> ./bin/tool-ssl-renewal -c ./bin/test.conf"
./bin/tool-ssl-renewal -c ./data/test.conf