package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"charm.land/huh/v2"
)

type sessionAction string

const (
	actionNext   sessionAction = "next"
	actionRepeat sessionAction = "repeat"
	actionChoose sessionAction = "choose"
	actionQuit   sessionAction = "quit"
)

func choosePractice(saved progress) (int, error) {
	options := make([]huh.Option[int], 0, len(practiceCatalog))
	for _, practice := range practiceCatalog {
		marker := "○"
		if saved.isCompleted(practice.number) {
			marker = "✓"
		}

		label := fmt.Sprintf("%s %02d: %s", marker, practice.number, practice.title)
		options = append(options, huh.NewOption(label, practice.number))
	}

	var selected int
	field := huh.NewSelect[int]().
		Title("Choose a practice").
		Options(options...).
		Value(&selected)

	return selected, runSelect(field)
}

func chooseNextAction(number int) (sessionAction, error) {
	options := make([]huh.Option[sessionAction], 0, 4)
	if next, found := followingPractice(number); found {
		label := fmt.Sprintf("Next: %02d — %s", next.number, next.title)
		options = append(options, huh.NewOption(label, actionNext))
	}
	options = append(
		options,
		huh.NewOption("Repeat this practice", actionRepeat),
		huh.NewOption("Choose another practice", actionChoose),
		huh.NewOption("Quit", actionQuit),
	)

	var selected sessionAction
	field := huh.NewSelect[sessionAction]().
		Title("What next?").
		Options(options...).
		Value(&selected)

	return selected, runSelect(field)
}

func runSelect[T comparable](field *huh.Select[T]) error {
	form := huh.NewForm(huh.NewGroup(field)).WithShowHelp(false)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	result := make(chan error, 1)
	go func() {
		result <- form.RunWithContext(ctx)
	}()

	select {
	case <-ctx.Done():
		return huh.ErrUserAborted
	case err := <-result:
		if ctx.Err() != nil {
			return huh.ErrUserAborted
		}
		return err
	}
}
