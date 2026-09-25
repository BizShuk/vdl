#!/usr/bin/env bash
# run:setup — 冪等 (idempotent) 的設定前置作業，不安裝任何軟體，不啟動任何服務。
# 設定的唯一位置是 ~/.config/vdl，repo 內的 tmp/config 只是指過去的 symlink。

set -euo pipefail

PROJECT_ROOT="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
CONFIG_DIR="${XDG_CONFIG_HOME:-${HOME}/.config}/vdl"
WORKSPACE_LINK="$PROJECT_ROOT/tmp/config"

mkdir -p "$CONFIG_DIR/data" "$CONFIG_DIR/logs" "$PROJECT_ROOT/tmp"

ln -sfn "$CONFIG_DIR" "$WORKSPACE_LINK"
echo "已備妥設定目錄與工作區軟連結：$WORKSPACE_LINK -> $CONFIG_DIR"
