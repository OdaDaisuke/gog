package main

import "github.com/urfave/cli"

func genCLIApp() *cli.App {
	app := cli.NewApp()
	app.Name = "gogit"
	app.Usage = "Awesome git client made by go."
	app.Version = "0.0.1"
	return app
}