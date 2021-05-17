#!/usr/bin/env bash

set +e

./build.sh

./bin/tool-ssl-renewal -c ./bin/test.conf