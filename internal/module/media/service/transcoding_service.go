package service

import (
	"log"
	"os/exec"

	"chillhub/internal/module/media/model"
)

type TranscodingService struct {}

func NewTranscodingService() *TranscodingService {
	return &TranscodingService{}
}

func (t *TranscodingService) Transcode(media *model.Media) {
	log.Println("Start transcoding:", media.ID.Hex())

	cmd := exec.Command(
		"ffmpeg",
		"-i", "/tmp/input.mp4",
		"-hls_time", "4",
		"-hls_playlist_type", "vod",
		"/tmp/output.m3u8",
	)

	if err := cmd.Run(); err != nil {
		log.Println("Transcode failed:", err)
	}
}
