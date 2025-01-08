package deepwork

import "time"

type DistractionCounter struct {
	ProgramAttempts map[string]int
	LastAttemptTime time.Time
}

func NewDistractionCounter() *DistractionCounter {
	return &DistractionCounter{
		ProgramAttempts: make(map[string]int),
	}
}

func (dc *DistractionCounter) RecordAttempt(program string) {
	dc.ProgramAttempts[program]++
	dc.LastAttemptTime = time.Now()
}

func (dc *DistractionCounter) ResetIfExpired(expiration time.Duration) {
	if time.Since(dc.LastAttemptTime) > expiration {
		dc.ProgramAttempts = make(map[string]int)
	}
}

func (dc *DistractionCounter) TotalAttempts() int {
	total := 0
	for _, attempts := range dc.ProgramAttempts {
		total += attempts
	}
	return total
}