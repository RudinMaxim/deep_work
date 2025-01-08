package deepwork

import (
	"time"
)

// SessionState представляет состояние сессии глубокого труда
type SessionState int

const (
	Inactive SessionState = iota
	Active
	Completed
)

// DeepWorkSession представляет сессию глубокого труда
type DeepWorkSession struct {
	ID                  string
	State               SessionState
	StartTime           time.Time
	EndTime             time.Time
	FocusDuration       time.Duration
	DistractionAttempts int
	IsActive            bool
}

type Statistics struct {
	TotalSessions              int
	TotalFocusTime             time.Duration
	AverageFocusTime           time.Duration
	DistractionCount           int
	LongestSession             time.Duration
	Shortest                   time.Duration
	AverageDistraction         int
	LongestDistraction         time.Duration
	ShortestDistraction        time.Duration
	AverageDistractionAttempts int
}

// NewSession создает новую сессию глубокого труда
func NewSession(id string) *DeepWorkSession {
	return &DeepWorkSession{
		ID:                  id,
		State:               Inactive,
		StartTime:           time.Time{},
		EndTime:             time.Time{},
		FocusDuration:       0,
		DistractionAttempts: 0,
	}
}

// Start запускает сессию глубокого труда
func (s *DeepWorkSession) Start() {
	s.State = Active
	s.StartTime = time.Now()
}

// Stop останавливает сессию глубокого труда
func (s *DeepWorkSession) Stop() {
	s.State = Completed
	s.EndTime = time.Now()
	s.FocusDuration = s.EndTime.Sub(s.StartTime)
}

// AddDistraction добавляет попытку отвлечения
func (s *DeepWorkSession) AddDistraction() {
	s.DistractionAttempts++
}
