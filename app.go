package app

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
)

const (
  BUFSIZE = 1024
  GIT_FILE_PATH = "./.git"
)

func main() {
  app := cli.NewApp()
  app.Name = "gogit"
  app.Usage = "Awesome git client made by go."
  app.Version = "0.0.1"

  app.Action = func(ctx *cli.Context) {
    if ctx.Bool("c") {
      printConfig()
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

func printConfig() {
  file, err := os.Open(GIT_FILE_PATH + "/config")
  if err != nil {
    fmt.Println("[ERROR]: config file doesn't exists.")
  } else {
    buf := make([]byte, BUFSIZE)
    for {
      n, err := file.Read(buf)
      if n == 0 {
        break
      }
      if err != nil {
	fmt.Println("[ERROR]: An error occured while reading confg file.")
      }

      fmt.Print(string(buf[:n]))
    }

  }

}
