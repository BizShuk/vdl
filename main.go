// Command vdl downloads a video URL from any supported platform as mp3 or mp4
// at a numbered quality tier.
package main

import (
	"github.com/bizshuk/vdl/cmd"
	_ "github.com/bizshuk/vdl/cmd/bilibili"
	_ "github.com/bizshuk/vdl/cmd/facebook"
	_ "github.com/bizshuk/vdl/cmd/instagram"
	_ "github.com/bizshuk/vdl/cmd/tiktok"
	_ "github.com/bizshuk/vdl/cmd/twitter"
	_ "github.com/bizshuk/vdl/cmd/youtube"

	_ "github.com/bizshuk/vdl/svc/bilibili"
	_ "github.com/bizshuk/vdl/svc/facebook"
	_ "github.com/bizshuk/vdl/svc/generic"
	_ "github.com/bizshuk/vdl/svc/instagram"
	_ "github.com/bizshuk/vdl/svc/tiktok"
	_ "github.com/bizshuk/vdl/svc/twitter"
	_ "github.com/bizshuk/vdl/svc/youtube"
)

func main() {
	cmd.Execute()
}
