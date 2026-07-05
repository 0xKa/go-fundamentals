package cli

import (
	"errors"
	"fmt"
	"os"

	"charm.land/huh/v2"
)

type application struct {
	progress     progress
	progressPath string
}

// Start runs the command-line application.
func Start() {
	if err := run(os.Args[1:]); err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return
		}

		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	options, err := parseOptions(args, os.Stderr)
	if errors.Is(err, errHelpRequested) {
		return nil
	}
	if err != nil {
		return err
	}

	progressPath, err := progressFilePath()
	if err != nil {
		return err
	}

	if options.resetProgress {
		if err := resetProgress(progressPath); err != nil {
			return err
		}
		fmt.Println("Progress reset.")
		return nil
	}

	savedProgress, err := loadProgress(progressPath)
	if err != nil {
		return err
	}

	if options.list {
		printPracticeList(os.Stdout, savedProgress)
		return nil
	}

	app := application{
		progress:     savedProgress,
		progressPath: progressPath,
	}

	if options.hasPracticeNumber {
		return app.runAndTrack(options.practiceNumber)
	}

	var number int
	if options.resume {
		var found bool
		number, found = firstIncomplete(app.progress)
		if !found {
			fmt.Println("All practices are complete. Nice work!")
			return nil
		}
	} else {
		number, err = choosePractice(app.progress)
		if err != nil {
			return err
		}
	}

	return app.runSession(number)
}

func (app *application) runSession(number int) error {
	for {
		if err := app.runAndTrack(number); err != nil {
			return err
		}

		action, err := chooseNextAction(number)
		if err != nil {
			return err
		}

		switch action {
		case actionNext:
			next, _ := followingPractice(number)
			number = next.number
		case actionRepeat:
			continue
		case actionChoose:
			number, err = choosePractice(app.progress)
			if err != nil {
				return err
			}
		case actionQuit:
			return nil
		default:
			return huh.ErrUserAborted
		}
	}
}

func (app *application) runAndTrack(number int) error {
	if err := executePractice(number); err != nil {
		return err
	}

	if !app.progress.markCompleted(number) {
		return nil
	}

	if err := saveProgress(app.progressPath, app.progress); err != nil {
		return err
	}

	fmt.Printf(
		"\nProgress: %d/%d practices complete.\n",
		completedPracticeCount(app.progress),
		len(practiceCatalog),
	)

	return nil
}
