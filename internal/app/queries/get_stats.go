package queries

import (
	"encoding/json"
	"os"

	"deep-work-app/internal/domain/deepwork"
)

type GetStatsQuery struct {
	Statistics deepwork.Statistics `json:"statistics"`
}

func NewGetStatsQuery(repo deepwork.Repository) (*GetStatsQuery, error) {
	stats, err := repo.GetStatistics()
	if err != nil {
		return nil, err
	}

	return &GetStatsQuery{
		Statistics: stats,
	}, nil
}

func (q *GetStatsQuery) Execute() error {
	data, err := json.MarshalIndent(q.Statistics, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("stats.json", data, 0644)
}

func GetStats(repo deepwork.Repository) (string, error) {

}
