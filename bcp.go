package main

import (
  "fmt"
  "os"

  "gopkg.in/urfave/cli.v1"
)

func main() {
  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func(c *cli.Context) error {
    fmt.Printf("file %q", c.Args().Get(0))
    return nil
  }

  app.Run(os.Args)
}
