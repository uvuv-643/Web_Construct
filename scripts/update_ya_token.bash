#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

grep -v '^YA_TOKEN=' "$SCRIPT_DIR/../ai-proxy/.env" > "$SCRIPT_DIR/../ai-proxy/.env.tmp" && mv "$SCRIPT_DIR/../ai-proxy/.env.tmp" "$SCRIPT_DIR/../ai-proxy/.env"
echo -n "YA_TOKEN=" >> "$SCRIPT_DIR/../ai-proxy/.env"
yc iam create-token | head -1 >> "$SCRIPT_DIR/../ai-proxy/.env"