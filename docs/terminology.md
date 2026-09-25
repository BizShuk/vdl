# vdl 術語表 (Terminology)

| 術語 (Term)        | 定義 (Definition)                                                        |
| ------------------ | ----------------------------------------------------------------------- |
| Platform           | 影片來源網站（YouTube、Instagram、TikTok 等），由 URL hostname 自動偵測 |
| MediaType          | 輸出格式：`mp3`（音訊）或 `mp4`（影片）                                 |
| Quality Tier       | 品質等級 1–5，mp4 對應解析度上限，mp3 對應 bitrate                       |
| Cookie Browser     | 用於 `--cookies-from-browser` 的瀏覽器名稱（chrome, firefox 等）         |
| Format Selector    | yt-dlp 的 `--format` 參數值，決定要下載的串流組合                        |
| Output Contract    | stdout 只輸出檔案路徑，所有其他資訊走 stderr                            |
| TUI Monitor        | Bubble Tea 打造的互動式下載清單瀏覽器                                   |
| Data Dir           | 固定的下載輸出目錄 `~/.config/vdl/data`                                  |
| Config Dir         | 應用設定根目錄 `~/.config/vdl`，支援 XDG_CONFIG_HOME                     |
| Runtime Check      | 下載前先驗 yt-dlp 與 ffmpeg 是否存在於 PATH                             |
