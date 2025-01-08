package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DeepWorkSettings struct {
	SessionDurationMinutes      int      `json:"session_duration_minutes"`
	ShortBreakMinutes           int      `json:"short_break_minutes"`
	LongBreakMinutes            int      `json:"long_break_minutes"`
	SessionsUntilLongBreak      int      `json:"sessions_until_long_break"`
	NotificationIntervalMinutes int      `json:"notification_interval_minutes"`
	AutoStartTime               string   `json:"auto_start_time"`
	AutoEndTime                 string   `json:"auto_end_time"`
	WorkingDays                 []string `json:"working_days"`
	AutomaticBackup             bool     `json:"automatic_backup"`
	BackupInterval              int      `json:"backup_interval_minutes"`
	Brightness                  int      `json:"brightness"`
}

type Position struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Config struct {
	MusicFile        string           `json:"music_file"`
	MusicPlaylists   []string         `json:"music_playlists"`
	DeepWorkSettings DeepWorkSettings `json:"deep_work_settings"`
}

var config Config

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения конфигурационного файла: %v", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("ошибка разбора конфигурационного файла: %v", err)
	}

	config = cfg
	return &cfg, nil
}
