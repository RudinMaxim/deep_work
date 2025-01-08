package app

import (
	"fmt"
	"os"
	"time"

	"deep-work-app/internal/app/queries"
	"deep-work-app/internal/config"
	"deep-work-app/internal/domain/deepwork"
	"deep-work-app/internal/infrastructure/persistence"
	"deep-work-app/internal/infrastructure/services"
	"deep-work-app/internal/infrastructure/services/commands"
)

type Application struct {
	config          *config.Config
	deepWorkService *services.WorkspaceSetupService
}

func NewApplication(config *config.Config) *Application {
	repo := persistence.NewFileRepository()
	deepWorkService := deepwork.NewService(repo)

	return &Application{
		config:          config,
		deepWorkService: deepWorkService,
	}
}

func (app *Application) StartDeepWork() error {
	if err := commands.StartDeepWork(app.deepWorkService); err != nil {
		return fmt.Errorf("не удалось запустить режим глубокого труда: %w", err)
	}
	return nil
}

func (app *Application) StopDeepWork() error {
	if err := commands.StopDeepWork(app.deepWorkService); err != nil {
		return fmt.Errorf("не удалось остановить режим глубокого труда: %w", err)
	}
	return nil
}

func (app *Application) GetStatus() (string, error) {
	status, err := queries.GetStatus(app.deepWorkService)
	if err != nil {
		return "", fmt.Errorf("не удалось получить статус: %w", err)
	}
	return status, nil
}

func (app *Application) GetStats() (string, error) {
	stats, err := queries.GetStats(app.deepWorkService)
	if err != nil {
		return "", fmt.Errorf("не удалось получить статистику: %w", err)
	}
	return stats, nil
}

func (app *Application) Run() {
	if err := app.StartDeepWork(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := app.StopDeepWork(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
