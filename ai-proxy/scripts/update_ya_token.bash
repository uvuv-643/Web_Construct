#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

grep -v '^YA_TOKEN=' "$SCRIPT_DIR/../.env" > "$SCRIPT_DIR/../.env.tmp" && mv "$SCRIPT_DIR/../.env.tmp" "$SCRIPT_DIR/../.env"
echo -n "YA_TOKEN=" >> "$SCRIPT_DIR/../.env"
yc iam create-token | head -1 >> "$SCRIPT_DIR/../.env"