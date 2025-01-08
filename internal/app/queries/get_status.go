package queries

import (
	"fmt"
	"time"

	"deep-work-app/internal/domain/deepwork"
)

type GetStatusQuery struct {
	DeepWorkSession deepwork.DeepWorkSession
}

func NewGetStatusQuery(session deepwork.DeepWorkSession) *GetStatusQuery {
	return &GetStatusQuery{DeepWorkSession: session}
}

func (q *GetStatusQuery) Execute() string {
	if !q.DeepWorkSession.IsActive {
		return "Статус: неактивен"
	}

	duration := time.Since(q.DeepWorkSession.StartTime)
	return fmt.Sprintf("Статус: активен\nТекущая сессия: %s\nПопытки отвлечься: %d",
		formatDuration(duration), q.DeepWorkSession.DistractionAttempts)
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	if h > 0 {
		return fmt.Sprintf("%dч %dмин", h, m)
	}
	return fmt.Sprintf("%dмин", m)
}
