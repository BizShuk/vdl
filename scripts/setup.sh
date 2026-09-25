#!/usr/bin/env bash
# run:install — 冪等 (idempotent) 的主機依賴軟體檢查與安裝，不碰設定。

set -euo pipefail

MISSING_DEPS=()

for bin in yt-dlp ffmpeg; do
    if ! command -v "$bin" >/dev/null 2>&1; then
        MISSING_DEPS+=("$bin")
    fi
done

if [ ${#MISSING_DEPS[@]} -eq 0 ]; then
    echo "系統已安裝所需外部工具：yt-dlp, ffmpeg"
    exit 0
fi

echo "缺少以下必要外部工具: ${MISSING_DEPS[*]}"
if command -v brew >/dev/null 2>&1; then
    echo "正在透過 Homebrew 進行安裝: brew install ${MISSING_DEPS[*]}"
    brew install "${MISSING_DEPS[@]}"
else
    echo "錯誤：未找到 Homebrew，請手動安裝 ${MISSING_DEPS[*]} 後重試。" >&2
    exit 1
fi
