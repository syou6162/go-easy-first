package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func shuffle(data []*Sentence) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

var commandTrain = cli.Command{
	Name:  "train",
	Usage: "Train a parsing model by easy-first algorithm",
	Description: `
Train a parsing model by easy-first algorithm.
`,
	Action: doTrain,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "train-filename"},
		cli.StringFlag{Name: "dev-filename"},
		cli.StringFlag{Name: "model-filename"},
		cli.IntFlag{Name: "max-iter", Value: 10},
	},
}

var commandEval = cli.Command{
	Name:  "eval",
	Usage: "Evaluate a parsing model by easy-first algorithm",
	Description: `
Evaluate a parsing model by easy-first algorithm.
`,
	Action: doEval,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "test-filename"},
		cli.StringFlag{Name: "model-filename"},
	},
}

var Commands = []cli.Command{
	commandTrain,
	commandEval,
}

func doTrain(c *cli.Context) error {
	trainFilename := c.String("train-filename")
	devFilename := c.String("dev-filename")
	modelFilename := c.String("model-filename")
	maxIter := c.Int("max-iter")

	if trainFilename == "" {
		_ = cli.ShowCommandHelp(c, "train")
		return cli.NewExitError("`train-filename` is a required field to train a parser.", 1)
	}

	if devFilename == "" {
		_ = cli.ShowCommandHelp(c, "train")
		return cli.NewExitError("`dev-filename` is a required field to train a parser.", 1)
	}

	if modelFilename == "" {
		_ = cli.ShowCommandHelp(c, "train")
		return cli.NewExitError("`model-filename` is a required field to train a parser.", 1)
	}

	goldSents, _ := ReadData(trainFilename)
	devSents, _ := ReadData(devFilename)

	model := NewModel()
	for iter := 0; iter < maxIter; iter++ {
		shuffle(goldSents)
		for _, sent := range goldSents {
			model.Update(sent)
		}
		w := model.AveragedWeight()
		trainAccuracy := DependencyAccuracy(&w, goldSents)
		devAccuracy := DependencyAccuracy(&w, devSents)
		fmt.Println(fmt.Sprintf("%d, %0.03f, %0.03f", iter, trainAccuracy, devAccuracy))
	}

	w := model.AveragedWeight()
	SaveModel(&w, modelFilename)
	return nil
}

func doEval(c *cli.Context) error {
	testFilename := c.String("test-filename")
	modelFilename := c.String("model-filename")

	if testFilename == "" {
		_ = cli.ShowCommandHelp(c, "eval")
		return cli.NewExitError("`test-filename` is a required field to evaluate a parser.", 1)
	}

	if modelFilename == "" {
		_ = cli.ShowCommandHelp(c, "eval")
		return cli.NewExitError("`model-filename` is a required field to evaluate a parser.", 1)
	}

	goldSents, _ := ReadData(testFilename)
	weight, _ := LoadModel(modelFilename)
	start := time.Now()
	testAccuracy := DependencyAccuracy(weight, goldSents)
	end := time.Now().Sub(start).Seconds()

	data := [][]string{
		{fmt.Sprintf("%d", len(goldSents)), fmt.Sprintf("%0.02f", end), fmt.Sprintf("%0.03f", testAccuracy)},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sentences", "Seconds", "Accuracy"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "easy-first"
	app.Commands = Commands

	runtime.GOMAXPROCS(runtime.NumCPU())

	app.Run(os.Args)
}
