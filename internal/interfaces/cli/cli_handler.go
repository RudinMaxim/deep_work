package cli

import (
	"fmt"
	"os"

	"deep-work-app/internal/app"

	"github.com/spf13/cobra"
)

type CLIHandler struct {
	App *app.Application
}

func NewCLIHandler(app *app.Application) *CLIHandler {
	return &CLIHandler{App: app}
}

func (h *CLIHandler) Execute() {
	rootCmd := &cobra.Command{
		Use:   "deep_work",
		Short: "Deep Work CLI application",
	}

	rootCmd.AddCommand(h.startCommand())
	rootCmd.AddCommand(h.stopCommand())
	rootCmd.AddCommand(h.statusCommand())
	rootCmd.AddCommand(h.statsCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (h *CLIHandler) startCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Запуск режима Deep Work",
		Run: func(cmd *cobra.Command, args []string) {
			h.App.StartDeepWork()
		},
	}
}

func (h *CLIHandler) stopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Остановка режима Deep Work",
		Run: func(cmd *cobra.Command, args []string) {
			h.App.StopDeepWork()
		},
	}
}

func (h *CLIHandler) statusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Показать текущий статус",
		Run: func(cmd *cobra.Command, args []string) {
			// h.App.ShowStatus()
		},
	}
}

func (h *CLIHandler) statsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Показать статистику",
		Run: func(cmd *cobra.Command, args []string) {
			// h.App.ShowStats()
		},
	}
}
