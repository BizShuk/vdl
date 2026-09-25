# vdl (video-downloader)

General-purpose CLI that downloads a video URL from **any supported platform**
as **mp3** or **mp4** at a numbered quality tier. It is a thin wrapper: `yt-dlp`
fetches the streams and `ffmpeg` performs muxing and transcoding. The platform
is auto-detected from the URL, and platform-specific authentication (cookies)
is applied automatically when needed.

## Supported Platforms

| Platform    | URL Examples                               | Auth           |
| ----------- | ------------------------------------------ | -------------- |
| YouTube     | `youtube.com`, `youtu.be`                  | —              |
| Instagram   | `instagram.com/reel/...`                   | cookies (auto) |
| TikTok      | `tiktok.com/@user/video/...`               | —              |
| Twitter / X | `twitter.com`, `x.com`                     | —              |
| Facebook    | `facebook.com/watch/...`, `fb.watch/...`   | cookies (auto) |
| Reddit      | `reddit.com`, `v.redd.it`                  | —              |
| Vimeo       | `vimeo.com/...`                            | —              |
| Bilibili    | `bilibili.com/video/...`, `b23.tv/...`     | —              |
| Generic     | any other yt-dlp-supported site            | —              |

Instagram and Facebook default to `--cookies-from-browser chrome`. Override
with `-cookies-from <browser>` or set to empty to skip.

## Usage

```bash
vdl [flags] <url>...
vdl config
vdl monitor    # alias: vdl m
```

### Commands

| Command               | Meaning                                      |
| --------------------- | -------------------------------------------- |
| `config`              | show config folder path                      |
| `monitor` (`m`)       | interactive TUI to browse downloaded files   |

### Flags

| Flag             | Default | Meaning                                              |
| ---------------- | ------- | ---------------------------------------------------- |
| `-type`          | `mp3`   | output format: `mp3` or `mp4`                        |
| `-qtype`         | `3`     | quality tier `1` (lowest) to `5` (highest)           |
| `-cookies-from`  | —       | browser to read cookies from (chrome, firefox, etc.) |

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

# Instagram Reel
vdl -type mp4 https://instagram.com/reel/ABC123/           # auto-uses Chrome cookies

# TikTok
vdl -type mp4 https://tiktok.com/@user/video/123

# Twitter/X
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

Download only material you have the right to download, and respect the source
site's terms of service.
