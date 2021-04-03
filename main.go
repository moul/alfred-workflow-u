package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"go.deanishe.net/fuzzy"
	"go.uber.org/zap"
	"moul.io/srand"
	"moul.io/zapconfig"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

type workflow struct {
	logger *zap.Logger
	wf     *aw.Workflow
}

func run(_ []string) error {
	var wf workflow
	// nolint:gomnd
	wf.wf = aw.New(
		aw.MaxResults(100),
		aw.SortOptions(
			fuzzy.AdjacencyBonus(10.0),
			fuzzy.LeadingLetterPenalty(-0.1),
			fuzzy.MaxLeadingLetterPenalty(-3.0),
			fuzzy.UnmatchedLetterPenalty(-0.5),
		),
		update.GitHub("moul/alfred-workflow-u"),
		aw.HelpURL("moul/alfred-workflow-u/issues"),
	)
	rand.Seed(srand.Fast())
	logger, err := zapconfig.Configurator{}.Build()
	if err != nil {
		return err
	}
	wf.logger = logger

	wf.wf.Run(wf.run)
	return nil
}

func (wf *workflow) run() {
	wf.wf.Args()
	flag.Parse()

	query := ""
	if args := flag.Args(); len(args) > 0 {
		query = args[0]
	}

	wf.wf.NewItem("foo").
		Subtitle("bar").
		Arg("https://github.com/moul/alfred-workflow-u#foo").
		UID("foo").
		Valid(true)
	wf.wf.NewItem("bar").
		Subtitle("baz").
		Arg("https://github.com/moul/alfred-workflow-u#bar").
		UID("bar").
		Valid(true)
	wf.wf.Filter(query)
	wf.wf.WarnEmpty("No matching items", "Try a different query?")
	wf.wf.SendFeedback()
}
