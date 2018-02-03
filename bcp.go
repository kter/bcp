package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	// "time"

	"gopkg.in/urfave/cli.v1"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	app := cli.NewApp()
	app.Name = "bcp"
	app.Usage = "bcp backup-file-name"
	app.Action = func(c *cli.Context) error {
		sourceFile := c.Args().Get(0)
		fmt.Printf("file %q", sourceFile)

		dir, err := os.Getwd()
		if err != nil {
			cli.NewExitError("Unable to get current location", 1)
		}
		fmt.Printf(dir)
		// olddir := dir + "/old"

		// if dir != 0 {
		// 	fmt.Printf("determine prefix number")
		// } else {
		// 	fmt.Printf("create new old dir")
		// }

		// prefix := "-01"
		// if prefix == nil {

		// }

		// now := time.Now()
		// date :=  now.Format("20060102")

		// srcFileName := c.Args()
		// dstFileName := srcFileName +  +

		// b, err := ioutil.ReadFile("input.txt")
		// if err != nil {
		//     panic(err)
		// }

		// // write the whole body at once
		// err = ioutil.WriteFile("output.txt", b, 0644)
		// if err != nil {
		//     panic(err)
		// }
		return nil
	}

	app.Run(os.Args)
}
