package commands

import (
	"fmt"
	"os"
	"time"

	"deep-work-app/internal/app"
	"deep-work-app/internal/infrastructure/services"
)

type StartDeepWorkCommand struct {
	application *app.Application
}

func NewStartDeepWorkCommand(application *app.Application) *StartDeepWorkCommand {
	return &StartDeepWorkCommand{application: application}
}

func (cmd *StartDeepWorkCommand) Execute() error {
	if cmd.application.IsDeepWorkActive() {
		return fmt.Errorf("режим 'Deep Work' уже запущен")
	}

	if err := cmd.application.LoadConfig(); err != nil {
		return fmt.Errorf("ошибка загрузки конфигурации: %v", err)
	}

	if err := cmd.application.SetupWorkspace(); err != nil {
		return fmt.Errorf("ошибка настройки рабочего пространства: %v", err)
	}

	services.PlayMusic(cmd.application.Config.MusicFile)

	cmd.application.StartMonitoringDistractions()

	cmd.application.Statistics.TotalSessions++
	cmd.application.Statistics.LastSession = time.Now()
	if err := cmd.application.SaveStatistics(); err != nil {
		return fmt.Errorf("ошибка сохранения статистики: %v", err)
	}

	if err := os.WriteFile(app.StatusFile, []byte("active"), 0644); err != nil {
		return fmt.Errorf("ошибка создания файла-флага: %v", err)
	}

	fmt.Println("Режим 'Deep Work' запущен. Успехов в работе!")
	return nil
}
