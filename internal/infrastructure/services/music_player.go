package services

import (
	"os/exec"
	"runtime"
)

type MusicPlayer struct {
	file string
}

func NewMusicPlayer(file string) *MusicPlayer {
	return &MusicPlayer{file: file}
}

func (mp *MusicPlayer) Play() error {
	if mp.file == "" {
		return nil
	}

	switch runtime.GOOS {
	case "windows":
		return exec.Command("powershell", "-Command", "(New-Object Media.SoundPlayer '"+mp.file+"').PlayLooping()").Start()
	case "darwin":
		return exec.Command("afplay", mp.file).Start()
	case "linux":
		return exec.Command("mpg123", mp.file).Start()
	default:
		return nil
	}
}

func (mp *MusicPlayer) Stop() error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("taskkill", "/F", "/IM", "wmplayer.exe").Run()
	case "darwin":
		return exec.Command("killall", "afplay").Run()
	case "linux":
		return exec.Command("killall", "mpg123").Run()
	default:
		return nil
	}
}