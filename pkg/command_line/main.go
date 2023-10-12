package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli demo"
	app.Author = "author"
	app.Commands = []cli.Command{
		cli.Command{
			Name:         "update",
			ShortName:    "",
			Aliases:      nil,
			Usage:        "update --config=file --coustom=true",
			UsageText:    "",
			Description:  "",
			ArgsUsage:    "",
			Category:     "",
			BashComplete: nil,
			Before:       nil,
			After:        nil,
			Action:       func(ctx *cli.Context) { fmt.Println("call update", ctx.String("config"), ctx.Bool("custom")) },
			OnUsageError: nil,
			Subcommands:  nil,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "config",
					Usage:    "--config",
					Required: true,
				},
				cli.BoolFlag{
					Name:     "custom",
					Usage:    "--custom",
					Required: true,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

// go build ./main.go
// ./main update --config=file1 --custom=true
