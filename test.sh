#!/usr/bin/env bash

set +e

echo ">> ./build.sh"
./build.sh
echo

echo ">> ./bin/tool-ssl-renewal -c ./data/run.toml"
./bin/tool-ssl-renewal -c ./data/run.toml