#!/usr/bin/env bash
echo ">> ./build.sh"
./build.sh
echo

echo ">> ./bin/tool-ssl-renewal test  -c data/run.toml  -k ProjectId"
./bin/tool-ssl-renewal test  -c data/run.toml  -k ProjectId