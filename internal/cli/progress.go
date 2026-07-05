package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type progress struct {
	Completed []int `json:"completed"`
}

func progressFilePath() (string, error) {
	configDirectory, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find user config directory: %w", err)
	}

	return filepath.Join(configDirectory, "go-fundamentals", "progress.json"), nil
}

func loadProgress(path string) (progress, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return progress{}, nil
	}
	if err != nil {
		return progress{}, fmt.Errorf("read progress: %w", err)
	}

	var saved progress
	if err := json.Unmarshal(data, &saved); err != nil {
		return progress{}, fmt.Errorf("decode progress file %q: %w", path, err)
	}

	sort.Ints(saved.Completed)
	return saved, nil
}

func saveProgress(path string, saved progress) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create progress directory: %w", err)
	}

	data, err := json.MarshalIndent(saved, "", "  ")
	if err != nil {
		return fmt.Errorf("encode progress: %w", err)
	}
	data = append(data, '\n')

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("save progress: %w", err)
	}

	return nil
}

func resetProgress(path string) error {
	err := os.Remove(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return fmt.Errorf("reset progress: %w", err)
	}
	return nil
}

func (saved progress) isCompleted(number int) bool {
	for _, completed := range saved.Completed {
		if completed == number {
			return true
		}
	}
	return false
}

func (saved *progress) markCompleted(number int) bool {
	if saved.isCompleted(number) {
		return false
	}

	saved.Completed = append(saved.Completed, number)
	sort.Ints(saved.Completed)
	return true
}
