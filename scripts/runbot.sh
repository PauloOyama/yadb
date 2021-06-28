#!/usr/bin/env bash
# This is a helper script which automatically exports the variables set in the .env file at the root of this project
# to the application
set -e
set -a

cd "$(dirname "${BASH_SOURCE[0]}")/.." || exit
source .env
exec go run cmd/*.go
