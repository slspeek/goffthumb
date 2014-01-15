package ffmpeg

import (
  "fmt"
  "os/exec"
)

const (
  ffmpegthumbnailer = "/usr/bin/ffmpegthumbnailer"
)

func CreateThumb(infile string , outfile string, timeindication string, size int) (err error) {
  cmd := exec.Command(ffmpegthumbnailer, "-i", infile, "-o", outfile, 
    "-s", fmt.Sprint("%v",size) , "-t", timeindication)
  err = cmd.Run()
  return
}
