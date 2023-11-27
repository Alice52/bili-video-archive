package util

import (
	"os"
	"os/exec"
)

func Merge(video, audio, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-y",
		"-i", video,
		"-i", audio,
		"-c", "copy", // Just copy without re-encoding
		"-shortest", // Finish encoding when the shortest input stream ends
		outputFile,
		"-loglevel", "warning",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
