# CLAUDE.md — vdl 技術脈絡 (Technical Context)

General-purpose video download CLI。業務定義見 [README.md](README.md)。
Module `github.com/bizshuk/video-downloader`，Go 1.26。TUI 採用 `bubbletea` 與 `lipgloss`。

## 結構與 ownership

```text
video-downloader/
├── main.go              # flag 解析 + subcommand 派工 + 逐 URL 派工 + platform auto-detect
└── pkg/
    ├── config/
    │   └── config.go    # config 與 data 目錄解析與建立 (~/.config/vdl)
    ├── download/
    │   ├── quality.go   # MediaType 與 quality tier → yt-dlp format selector
    │   └── download.go  # Request 驗證、runtime 檢查、exec yt-dlp (含 platform extra args)
    ├── platform/
    │   └── platform.go  # URL hostname → Platform 自動偵測、per-platform auth/args
    └── tui/
        ├── file.go      # 下載檔案讀取、格式化、系統預設開啟
        ├── model.go     # Bubble Tea Model、狀態轉換、事件迴圈
        ├── run.go       # TUI 程式進入點 (alt-screen)
        ├── view.go      # 視圖渲染（Header/Banner/Table/Footer 高度預算）
        └── width.go     # 字符寬度測量與截斷（runewidth 私有實例）
```

`pkg/download` 不認識 flag，`main.go` 不認識 yt-dlp 參數。
`pkg/platform` 只做 URL → Platform 映射與 per-platform yt-dlp args，不碰網路。

## 關鍵決策 (Key Decisions)

- CLI 用 **stdlib `flag`** 而非 cobra：使用者要的是 `-type` / `-qtype`
  單破折號長旗標，pflag 會把 `-type` 當成短旗標叢集解析而失敗。
- **Platform auto-detection** 從 URL hostname 推斷來源平台，自動套用
  per-platform yt-dlp 參數（如 `--cookies-from-browser`）。
- Instagram / Facebook 預設自動使用 `--cookies-from-browser chrome`，
  使用者可透過 `-cookies-from` 覆蓋或指定其他瀏覽器。
- 從 `ytdl` 演化而來，完全相容 ytdl 的 quality tier 系統、output contract
  與 TUI monitor。差異在於新增 platform detection 與 cookie 自動注入。
- 下載引擎是 **yt-dlp**，支援 1700+ 網站。ffmpeg 同時服務 mp4 合流與
  mp3 轉檔，兩者在下載開始前一次檢查 (`checkRuntime`)。
- 輸出契約：**stdout 只有檔案路徑**，進度與橫幅走 stderr。
- 輸出目錄固定為 `~/.config/vdl/data`（不存在則自動建立）。
- 預設 `--no-playlist`：playlist URL 只抓所指向的單支影片。
- `signal.NotifyContext` 綁 `exec.CommandContext`，Ctrl-C 會連帶終止
  yt-dlp 子行程，不留孤兒。

## 開發與驗證 (Development and Verification)

```bash
gofmt -l . && go vet ./... && go test ./... -count=1   # 全離線
go build -o bin/vdl . && ./bin/vdl -help
```

單元測試全離線（只驗參數映射與 platform detection，不打網路）。
真實下載屬手動 smoke test。
