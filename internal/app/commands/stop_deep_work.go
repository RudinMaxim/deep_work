package commands

import (
	"fmt"
	"os"
	"time"

	"deep-work-app/internal/app"
	"deep-work-app/pkg/utils"
)

type StopDeepWorkCommand struct {
	application *app.Application
}

func NewStopDeepWorkCommand(application *app.Application) *StopDeepWorkCommand {
	return &StopDeepWorkCommand{application: application}
}

func (cmd *StopDeepWorkCommand) Execute() error {
	if !cmd.application.IsDeepWorkActive() {
		fmt.Println("Режим 'Deep Work' не запущен.")
		return nil
	}

	// Сохранение статистики
	cmd.application.SaveStatistics()

	// Создание финального бэкапа
	cmd.application.AutoBackup()

	// Восстановление настроек дисплея
	if err := cmd.application.RestoreDisplaySettings(); err != nil {
		return fmt.Errorf("ошибка восстановления настроек дисплея: %v", err)
	}

	// Остановка фоновой музыки
	cmd.application.StopMusic()

	// Удаление файла-флага
	if err := os.Remove(app.StatusFile); err != nil {
		return fmt.Errorf("ошибка удаления файла-флага: %v", err)
	}

	// Показ итогов сессии
	cmd.showSessionSummary()

	fmt.Println("Режим 'Deep Work' остановлен. Хорошая работа!")
	return nil
}

func (cmd *StopDeepWorkCommand) showSessionSummary() {
	summary := fmt.Sprintf(
		"Итоги сессии:\n"+
			"Продолжительность: %s\n"+
			"Попытки отвлечься: %d\n"+
			"Продуктивность: %.1f%%\n"+
			"Всего сессий: %d",
		utils.FormatDuration(time.Since(cmd.application.LastSession)),
		cmd.application.DistractionAttempts,
		cmd.application.CalculateDailyProductivity(),
		cmd.application.TotalSessions,
	)

	fmt.Println(summary)
}
