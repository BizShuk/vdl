# vdl (video-downloader)

General-purpose CLI that downloads a video URL from **any supported platform**
as **mp3** or **mp4** at a numbered quality tier. It is a thin wrapper: `yt-dlp`
fetches the streams and `ffmpeg` performs muxing and transcoding. The platform
is auto-detected from the URL, or can be explicitly dispatched via platform subcommands.

> **遷移指南**：若由 `ytdl` 升級而來，可參考 [skills/migrate-from-ytdl/SKILL.md](skills/migrate-from-ytdl/SKILL.md)。

## Supported Platforms

| Platform    | Subcommand   | URL Examples                               | Auth           |
| ----------- | ------------ | ------------------------------------------ | -------------- |
| YouTube     | `youtube`    | `youtube.com`, `youtu.be`                  | —              |
| Instagram   | `instagram`  | `instagram.com/reel/...`, `.../p/...`      | cookies (advisory) |
| TikTok      | `tiktok`     | `tiktok.com/@user/video/...`               | —              |
| Twitter / X | `twitter`    | `twitter.com`, `x.com`                     | —              |
| Facebook    | `facebook`   | `facebook.com/watch/...`, `fb.watch/...`   | cookies (advisory) |
| Reddit      | (generic)    | `reddit.com`, `v.redd.it`                  | —              |
| Vimeo       | (generic)    | `vimeo.com/...`                            | —              |
| Bilibili    | `bilibili`   | `bilibili.com/video/...`, `b23.tv/...`     | —              |
| Generic     | (fallback)   | any other yt-dlp-supported site            | —              |

## Usage

```bash
# 直接下載（自動辨識來源平台）
vdl [flags] <url>...

# 或透過平台子命令
vdl <platform> [flags] <url>...

# 工具命令
vdl config
vdl monitor    # alias: vdl m
```

### Commands

| Command               | Meaning                                      |
| --------------------- | -------------------------------------------- |
| `config`              | show config folder path                      |
| `monitor` (`m`)       | interactive TUI to browse downloaded files   |
| `youtube` (`yt`)      | download from YouTube                        |
| `instagram` (`ig`)    | download from Instagram                      |
| `tiktok`              | download from TikTok                         |
| `twitter` (`x`)       | download from Twitter/X                      |
| `facebook` (`fb`)     | download from Facebook                       |
| `bilibili` (`bili`)   | download from Bilibili                       |

### Flags

| Flag (POSIX / Long / Short)   | Default | Meaning                                              |
| ----------------------------- | ------- | ---------------------------------------------------- |
| `-type` / `--type` / `-t`     | `mp3`   | output format: `mp3` or `mp4`                        |
| `-qtype` / `--qtype` / `-q`   | `3`     | quality tier `1` (lowest) to `5` (highest)           |
| `-cookies-from` / `--cookies-from` | —  | browser to read cookies from (chrome, firefox, etc.) |

Downloads land in `~/.config/vdl/data` (fixed).

Multiple URLs may be passed; each is downloaded with the same settings.
Playlist URLs download only the single referenced video.

## Quality Tiers

| `-qtype` | `-type mp4` (max resolution) | `-type mp3` (audio bitrate) |
| -------- | ---------------------------- | --------------------------- |
| 1        | 360p                         | 64 kbps                     |
| 2        | 480p                         | 128 kbps                    |
| 3        | 720p                         | 192 kbps                    |
| 4        | 1080p                        | 256 kbps                    |
| 5        | best available (uncapped)    | 320 kbps                    |

## Examples

```bash
# YouTube
vdl https://youtu.be/VIDEO_ID                              # 192 kbps mp3
vdl -type mp4 -qtype 5 https://youtu.be/VIDEO_ID           # best quality mp4
vdl youtube -t mp4 -q 5 https://youtu.be/VIDEO_ID          # via platform subcommand

# Instagram Reel
vdl -type mp4 https://instagram.com/reel/ABC123/           # auto-detected
vdl instagram -t mp4 -q 5 https://instagram.com/reel/ABC/  # via instagram subcommand

# TikTok & Twitter
vdl -type mp4 https://tiktok.com/@user/video/123
vdl -type mp4 https://x.com/user/status/123

# Explicit cookie browser override
vdl -cookies-from firefox https://instagram.com/reel/ABC123/

# Config and Monitor
vdl config                                                  # prints ~/.config/vdl
vdl m                                                       # interactive TUI
```

## Output Contract

On success **stdout carries only the resulting file path**, one line per URL,
so it composes with other tools:

```bash
open "$(vdl -type mp4 https://instagram.com/reel/ABC123/)"
```

Download progress, the per-URL banner, and all errors go to stderr.

## Requirements

`yt-dlp` and `ffmpeg` must be on `PATH`; both are checked before any download
starts.

```bash
brew install yt-dlp ffmpeg
```
