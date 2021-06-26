#!/usr/bin/env bash
# This script builds the bot's Docker image
set -e

cd "$(dirname "${BASH_SOURCE[0]}")/.." || exit
docker build --tag prot:0.0.1 --file scripts/docker/Dockerfile .
