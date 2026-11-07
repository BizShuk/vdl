# CLAUDE.md — vdl 技術脈絡 (Technical Context)

General-purpose video download CLI。業務定義見 [README.md](README.md)。
Module `github.com/bizshuk/vdl`，Go 1.26。TUI 採用 `bubbletea` 與 `lipgloss`，CLI 採用 `cobra` + Two layers 結構。

## 結構與 ownership (遵循 golang-dev)

```text
vdl/
├── main.go                     # cmd.Execute() only + blank-imports cmd/<platform> 與 svc/<platform>
├── cmd/
│   ├── root.go                 # RootCmd + Execute，支援直接傳 URL 抓取與旗標前置正規化
│   ├── monitor.go              # monitor (alias m) — interactive TUI
│   ├── config.go               # config
│   ├── youtube.go              # 第一層平台指令群組 (package cmd)
│   ├── youtube/                # 第二層平台指令 (package youtube)
│   │   └── download.go
│   ├── instagram.go            # 第一層平台指令群組
│   ├── instagram/
│   │   └── download.go
│   ├── tiktok.go / twitter.go / facebook.go / bilibili.go
│   └── tiktok/ / twitter/ / facebook/ / bilibili/
├── svc/
│   ├── platform.go             # Platform 介面定義、Options 結構與 Registry (Register/Detect/All)
│   ├── youtube/                # YouTube 平台服務與專屬設定
│   ├── instagram/              # Instagram 平台服務與認證/cookie 設定
│   ├── tiktok/ / twitter/ / facebook/ / bilibili/ / generic/
└── pkg/                        # 通用共用元件 (Common Components)
    ├── config/                 # config 與 data 目錄解析 (~/.config/vdl)
    ├── download/               # 通用 yt-dlp 執行、ffmpeg runtime 檢查、Quality Tiers
    └── tui/                    # Bubble Tea monitor TUI
```

- `main.go` 位於 repo 根目錄，只負責 `cmd.Execute()` 與 blank-imports。
- `cmd/<platform>/` 負責 CLI 子命令介面（第一層群組與第二層動詞）。
- `svc/<platform>/` 負責各平台的特定設定、URL 匹配與業務邏輯，各服務透過 `init()` 自我註冊到 `svc.Register()`。
- `pkg/` 僅包含與平台無關的共通元件（config, download runner, tui）。

## 關鍵決策 (Key Decisions)

- CLI 遵循 **golang-dev (cobra)** 規範：採用兩層子命令結構，同時在 `Execute()` 前置將使用者慣用的 `-type` / `-qtype` / `-cookies-from` 單破折號旗標無損轉為 POSIX 長旗標，達成與 `ytdl` 100% 向後相容。
- **透明 URL 派發**：若使用者直接給定 URL（`vdl [flags] <url>...`），自動路由至內建下載器，由 `svc.Detect(url)` 自動匹配所屬平台服務。
- **Platform 服務隔離**：各平台專屬邏輯（如 Facebook/Instagram 的 Cookies 需求與提取參數）完全封裝於 `svc/<platform>`，底層 `pkg/download` 不含任何平台特化邏輯。
- 下載引擎是 **yt-dlp**，支援 1700+ 網站。ffmpeg 同時服務 mp4 合流與 mp3 轉檔，下載前一次檢查 (`CheckRuntime`)。
- 輸出契約：**stdout 只有檔案路徑**，進度與橫幅走 stderr。
- 輸出目錄固定為 `~/.config/vdl/data`（不存在則自動建立）。
- 預設 `--no-playlist`：playlist URL 只抓所指向的單支影片。
- `signal.NotifyContext` 綁 `exec.CommandContext`，Ctrl-C 連帶終止 yt-dlp 子行程。

## 開發與驗證 (Development and Verification)

```bash
pnpm run lint && pnpm run test && pnpm run build
```

單元測試全離線（只驗參數映射與 platform detection，不打網路）。
真實下載屬手動 smoke test。
