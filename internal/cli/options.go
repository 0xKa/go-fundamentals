package cli

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"
)

type options struct {
	list              bool
	resume            bool
	resetProgress     bool
	hasPracticeNumber bool
	practiceNumber    int
}

var errHelpRequested = errors.New("help requested")

func parseOptions(args []string, output io.Writer) (options, error) {
	var parsed options

	flags := flag.NewFlagSet("go-fundamentals", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.BoolVar(&parsed.list, "list", false, "list practices and completion status")
	flags.BoolVar(&parsed.resume, "resume", false, "run the first incomplete practice")
	flags.BoolVar(&parsed.resetProgress, "reset-progress", false, "clear saved progress")
	flags.Usage = func() {
		printUsage(output)
	}

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return options{}, errHelpRequested
		}
		return options{}, err
	}

	actionCount := countTrue(parsed.list, parsed.resume, parsed.resetProgress)
	if actionCount > 1 {
		return options{}, errors.New("--list, --resume, and --reset-progress cannot be combined")
	}

	positional := flags.Args()
	if len(positional) > 1 {
		return options{}, errors.New("provide at most one practice number")
	}
	if len(positional) == 0 {
		return parsed, nil
	}
	if actionCount > 0 {
		return options{}, errors.New("a practice number cannot be combined with an option")
	}

	number, err := strconv.Atoi(positional[0])
	if err != nil {
		return options{}, fmt.Errorf("invalid practice number %q", positional[0])
	}

	parsed.hasPracticeNumber = true
	parsed.practiceNumber = number
	return parsed, nil
}

func printUsage(output io.Writer) {
	fmt.Fprintln(output, "Usage: go-fundamentals [options] [practice-number]")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "Options:")
	fmt.Fprintln(output, "  --list            list practices and completion status")
	fmt.Fprintln(output, "  --resume          run the first incomplete practice")
	fmt.Fprintln(output, "  --reset-progress  clear saved progress")
	fmt.Fprintln(output, "  --help            show this help")
}

func countTrue(values ...bool) int {
	count := 0
	for _, value := range values {
		if value {
			count++
		}
	}
	return count
}
