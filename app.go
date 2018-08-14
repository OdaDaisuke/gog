package main

import (
    "os"
    "github.com/urfave/cli"
    "github.com/OdaDaisuke/gogit/commands"
)

func main() {
  app := genCLIApp()

  app.Action = func(ctx *cli.Context) {
    if ctx.Bool("c") {
      commands.PrintConfig()
    }

  }

  app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "c",
      Usage: "Print user configs.",
    },
  }

  app.Run(os.Args)
}
