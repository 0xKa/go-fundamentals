package cli

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestProgressRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	saved := progress{}

	if !saved.markCompleted(2) {
		t.Fatal("first completion was not marked")
	}
	if !saved.markCompleted(1) {
		t.Fatal("second completion was not marked")
	}
	if saved.markCompleted(2) {
		t.Fatal("duplicate completion was marked")
	}

	if err := saveProgress(path, saved); err != nil {
		t.Fatal(err)
	}

	loaded, err := loadProgress(path)
	if err != nil {
		t.Fatal(err)
	}

	want := progress{Completed: []int{1, 2}}
	if !reflect.DeepEqual(loaded, want) {
		t.Fatalf("got %+v, want %+v", loaded, want)
	}
}

func TestResetProgress(t *testing.T) {
	path := filepath.Join(t.TempDir(), "progress.json")
	if err := saveProgress(path, progress{Completed: []int{1}}); err != nil {
		t.Fatal(err)
	}

	if err := resetProgress(path); err != nil {
		t.Fatal(err)
	}
	if err := resetProgress(path); err != nil {
		t.Fatalf("resetting missing progress returned an error: %v", err)
	}

	loaded, err := loadProgress(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Completed) != 0 {
		t.Fatalf("got completed practices %v, want none", loaded.Completed)
	}
}
