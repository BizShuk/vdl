// Command vdl downloads a video URL from any supported platform as mp3 or mp4
// at a numbered quality tier.
package main

import (
	"github.com/bizshuk/video-downloader/cmd"
	_ "github.com/bizshuk/video-downloader/cmd/bilibili"
	_ "github.com/bizshuk/video-downloader/cmd/facebook"
	_ "github.com/bizshuk/video-downloader/cmd/instagram"
	_ "github.com/bizshuk/video-downloader/cmd/tiktok"
	_ "github.com/bizshuk/video-downloader/cmd/twitter"
	_ "github.com/bizshuk/video-downloader/cmd/youtube"

	_ "github.com/bizshuk/video-downloader/svc/bilibili"
	_ "github.com/bizshuk/video-downloader/svc/facebook"
	_ "github.com/bizshuk/video-downloader/svc/generic"
	_ "github.com/bizshuk/video-downloader/svc/instagram"
	_ "github.com/bizshuk/video-downloader/svc/tiktok"
	_ "github.com/bizshuk/video-downloader/svc/twitter"
	_ "github.com/bizshuk/video-downloader/svc/youtube"
)

func main() {
	cmd.Execute()
}
