package deepwork

import (
	"errors"
	"sync"
)

// Repository интерфейс для доступа к данным сессий глубокого труда.
type Repository interface {
	SaveSession(session DeepWorkSession) error
	GetSession(id string) (DeepWorkSession, error)
	GetAllSessions() ([]DeepWorkSession, error)
	DeleteSession(id string) error
	GetStatistics() (Statistics, error)
}

type MemoryRepository struct {
	sessions map[string]DeepWorkSession
	mutex    sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		sessions: make(map[string]DeepWorkSession),
	}
}

func (r *MemoryRepository) SaveSession(session DeepWorkSession) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.sessions[session.ID] = session
	return nil
}

func (r *MemoryRepository) GetSession(id string) (DeepWorkSession, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	session, exists := r.sessions[id]
	if !exists {
		return DeepWorkSession{}, errors.New("session not found")
	}
	return session, nil
}

func (r *MemoryRepository) GetAllSessions() ([]DeepWorkSession, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	sessions := make([]DeepWorkSession, 0, len(r.sessions))
	for _, session := range r.sessions {
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (r *MemoryRepository) DeleteSession(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.sessions[id]; !exists {
		return errors.New("session not found")
	}
	delete(r.sessions, id)
	return nil
}

func (r *MemoryRepository) GetStatistics() (Statistics, error) {
	return Statistics{}, nil
}
