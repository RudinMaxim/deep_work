package persistence

import (
	"encoding/json"
	"os"

	"deep-work-app/internal/domain/deepwork"
)

type FileRepository struct {
	filePath string
}

func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{filePath: filePath}
}

func (r *FileRepository) Save(session deepwork.DeepWorkSession) error {
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}

func (r *FileRepository) Load() (deepwork.DeepWorkSession, error) {
	var session deepwork.DeepWorkSession
	data, err := os.ReadFile(r.filePath)
	if os.IsNotExist(err) {
		return session, nil
	}
	if err != nil {
		return session, err
	}
	err = json.Unmarshal(data, &session)
	return session, err
}

func (r *FileRepository) Delete() error {
	return os.Remove(r.filePath)
}

func (r *FileRepository) Exists() bool {
	_, err := os.Stat(r.filePath)
	return !os.IsNotExist(err)
}
