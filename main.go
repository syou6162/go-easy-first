package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"math/rand"
	"os"
	"runtime"
)

func shuffle(data []*Sentence) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

var commandTrain = cli.Command{
	Name:        "train",
	Usage:       "Train a parsing model by easy-first algorithm",
	Description: `
Train a parsing model by easy-first algorithm.
`,
	Action:      doTrain,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "train-filename"},
		cli.StringFlag{Name: "dev-filename"},
		cli.IntFlag{Name: "max-iter", Value: 10},
	},
}

var Commands = []cli.Command{
	commandTrain,
}

func doTrain(c *cli.Context) error {
	trainFilename := c.String("train-filename")
	devFilename := c.String("dev-filename")
	maxIter := c.Int("max-iter")

	if trainFilename == "" {
		_ = cli.ShowCommandHelp(c, "train")
		return cli.NewExitError("`train-filename` is a required field to train a parser.", 1)
	}

	if devFilename == "" {
		_ = cli.ShowCommandHelp(c, "train")
		return cli.NewExitError("`dev-filename` is a required field to train a parser.", 1)
	}

	goldSents, _ := ReadData(trainFilename)
	devSents, _ := ReadData(devFilename)

	model := NewModel()
	for iter := 0; iter < maxIter; iter++ {
		shuffle(goldSents)
		for _, sent := range goldSents {
			model.Update(sent)
		}

		trainAccuracy := DependencyAccuracy(&model, goldSents)
		devAccuracy := DependencyAccuracy(&model, devSents)
		fmt.Println(fmt.Sprintf("%d, %0.03f, %0.03f", iter, trainAccuracy, devAccuracy))
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "easy-first"
	app.Commands = Commands

	runtime.GOMAXPROCS(runtime.NumCPU())

	app.Run(os.Args)
}
