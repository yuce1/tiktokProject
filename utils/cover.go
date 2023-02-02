package utils

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GenVideoCover(path string, desPath string) error {

	buf := bytes.NewBuffer(nil)
	// default get the 10st frame in video
	err := ffmpeg.Input(path).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 10)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg", "loglevel": "quiet"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}

	var img image.Image
	img, err = imaging.Decode(buf)
	if err != nil {
		return err
	}

	err = imaging.Save(img, desPath)
	if err != nil {
		return err
	}

	return nil
}
