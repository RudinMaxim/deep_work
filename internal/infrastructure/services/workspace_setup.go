package services

import (
	"fmt"
	"os/exec"
	"runtime"

	"deep-work-app/internal/config"
)

// WorkspaceSetupService отвечает за настройку рабочего пространства
type WorkspaceSetupService struct {
	Settings config.DeepWorkSettings
}

// NewWorkspaceSetupService создает новый экземпляр WorkspaceSetupService
func NewWorkspaceSetupService(settings config.DeepWorkSettings) *WorkspaceSetupService {
	return &WorkspaceSetupService{Settings: settings}
}

// SetDisplaySettings настраивает яркость и цветовую температуру экрана
func (w *WorkspaceSetupService) SetDisplaySettings() error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf(
			"$brightness = %d; $monitor = Get-WmiObject -Namespace root/WMI -Class WmiMonitorBrightnessMethods; $monitor.WmiSetBrightness(1, $brightness)",
			w.Settings.Brightness,
		))
		return cmd.Run()
	case "darwin":
		cmd := exec.Command("brightness", fmt.Sprintf("%d", w.Settings.Brightness))
		return cmd.Run()
	default:
		return fmt.Errorf("настройка экрана не поддерживается для ОС %s", runtime.GOOS)
	}
}

// SetWindowPosition устанавливает позицию окна программы
func (w *WorkspaceSetupService) SetWindowPosition(program string, pos config.Position) error {
	switch runtime.GOOS {
	case "windows":
		script := fmt.Sprintf(`
			$proc = Get-Process "%s" -ErrorAction SilentlyContinue
			if ($proc) {
				$wshell = New-Object -ComObject wscript.shell
				$wshell.AppActivate($proc.MainWindowTitle)
				$window = $proc.MainWindowHandle
				[Windows.Win32.User32]::MoveWindow($window, %d, %d, %d, %d, $true)
			}
		`, program, pos.X, pos.Y, pos.Width, pos.Height)
		cmd := exec.Command("powershell", "-Command", script)
		return cmd.Run()
	default:
		return fmt.Errorf("позиционирование окон не поддерживается для ОС %s", runtime.GOOS)
	}
}
