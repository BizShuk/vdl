---
name: migrate-from-ytdl
description: Use when migrating from ytdl to video-downloader (vdl), replacing YouTube-only downloads with multi-platform video workflows, migrating download history, configuring backward-compatible aliases, or troubleshooting platform authentication.
---

# Migrate from ytdl to video-downloader (vdl)

## Overview

`vdl` (`video-downloader`) 是 `ytdl` 的功能超集 (superset)。它保留了 `ytdl` 的完整命令列旗標介面、品質分級契約 (quality tiers 1~5) 與 stdout 唯檔案路徑的管線組合特性，同時將支援範圍由單一 YouTube 擴展至 Instagram、TikTok、Twitter/X、Facebook、Reddit、Bilibili 等平台。

## When to Use

- 需要將現有 `ytdl` 下載工作流、Shell 腳本或排程任務升級至支援多平台時
- 需要下載 Instagram Reels/Posts、TikTok 短影音、X/Twitter 影片等非 YouTube 媒體時
- 需要將既有 `~/.config/ytdl/data` 歷史檔案整併入 `~/.config/vdl/data` 時
- 希望透過 Shell Alias 或 Symlink 達成零停機平滑升級時

## Quick Reference (對照表)

| 項目 | `ytdl` (舊) | `vdl` (新) | 相容性 |
| :--- | :--- | :--- | :--- |
| **主要指令** | `ytdl <url>...` | `vdl <url>...` | 完全相容（支援 URL 直接傳入） |
| **音訊格式** | `-type mp3` | `-type mp3` / `-t mp3` / `--type mp3` | 100% 向後相容 |
| **影片格式** | `-type mp4` | `-type mp4` / `-t mp4` / `--type mp4` | 100% 向後相容 |
| **品質分級** | `-qtype 1~5` | `-qtype 1~5` / `-q 1~5` / `--qtype 1~5` | 100% 向後相容 |
| **子命令** | `ytdl config`, `ytdl m` | `vdl config`, `vdl monitor` (`vdl m`) | 完全相容 |
| **平台專屬命令** | 無 | `vdl youtube`, `vdl instagram` 等 | 新增功能 |
| **認證旗標** | 無 | `--cookies-from <browser>` | 新增功能 |
| **輸出契約** | stdout 僅檔案路徑 | stdout 僅檔案路徑 | 100% 向後相容 |
| **資料目錄** | `~/.config/ytdl/data` | `~/.config/vdl/data` | 需遷移或連結 |

## Step-by-Step Migration Guide

### 1. 安裝與二進位編譯

在 `~/projects/tools/vdl` 執行安裝：

```bash
cd ~/projects/tools/vdl
pnpm run install
```

這將編譯並安裝 `vdl` 二進位檔至 `$GOPATH/bin`（或 `~/.local/go/bin`）。

### 2. 歷史下載檔案遷移

將原先在 `~/.config/ytdl/data` 的檔案移動或建立軟連結至 `~/.config/vdl/data`：

```bash
# 確保 vdl 設定樹已備妥
vdl config

# 方法 A：直接複製/搬移既有歷史檔案
mv ~/.config/ytdl/data/* ~/.config/vdl/data/ 2>/dev/null || true

# 方法 B：若舊腳本仍依賴舊路徑，可保留舊目錄並以軟連結導向新目錄
rmdir ~/.config/ytdl/data 2>/dev/null || true
ln -sfn ~/.config/vdl/data ~/.config/ytdl/data
```

### 3. Shell Alias 平滑過渡

若要在不更動現有肌肉記憶與 Shell 腳本的前提下立即套用 `vdl`，可在 Shell Profile (`~/.zshrc` 或 `~/.bashrc`) 設定：

```bash
alias ytdl="vdl"
```

此時所有既有的 `ytdl -type mp4 -qtype 5 <url>` 將自動經由 `vdl` 處理，並獲得多平台解析能力。

### 4. 自動化腳本與管線遷移

由於 `vdl` 嚴格維持與 `ytdl` 相同的輸出契約（**stdout 唯檔案絕對路徑，進度與診斷走 stderr**），所有既有的管線組合腳本可直接無痛替換 binary 名稱：

```bash
# 原 ytdl 寫法
open "$(ytdl -type mp4 -qtype 5 https://youtu.be/xxx)"

# 改用 vdl 寫法（直接替換指令即可）
open "$(vdl -type mp4 -qtype 5 https://youtu.be/xxx)"
```

## 多平台新功能運用

### 下載 Instagram Reel
Instagram 公開影片可直接下載；若遇到私人帳號或登入牆，可帶入瀏覽器 Cookie：

```bash
# 公開 Reel 直接下載 (1080p MP4)
vdl -type mp4 -qtype 5 "https://www.instagram.com/reels/DdrPiPKku2A/"

# 使用 Chrome 已登入之 Session
vdl -type mp4 -cookies-from chrome "https://www.instagram.com/reel/CODE/"
```

### 下載 TikTok / Twitter / Facebook
```bash
# TikTok
vdl -type mp4 "https://www.tiktok.com/@user/video/123"

# Twitter / X
vdl -type mp4 "https://x.com/user/status/123"

# Facebook 影片
vdl -type mp4 "https://fb.watch/abc123/"
```

## Common Mistakes & Troubleshooting

| 狀況 | 原因 | 處置方式 |
| :--- | :--- | :--- |
| `vdl: command not found` | Go bin 目錄未加入 `$PATH` | 確保 `~/.local/go/bin` 或 `~/go/bin` 在系統 `$PATH` 中 |
| `yt-dlp not found in PATH` | 缺少外部核心依賴 | 執行 `brew install yt-dlp ffmpeg` |
| Instagram 提示需要 Cookies | 影片為私密貼文或受登入牆限制 | 加入 `--cookies-from chrome`（或 `firefox`、`safari`）旗標 |
| 舊檔案在 `vdl monitor` 看不到 | 舊下載檔仍停留在 `~/.config/ytdl/data` | 將舊檔案移動或軟連結至 `~/.config/vdl/data` |
