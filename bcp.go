package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/urfave/cli.v1"
)

func file_existence(path string) (bool, error) {
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

		dir, err := os.Getwd()
		if err != nil {
			cli.NewExitError("Unable to get current location", 1)
		}
		olddir := dir + "/old"

		dirExists, err := file_existence(olddir)
		if err != nil {
			cli.NewExitError("Unable to get destination dir location", 1)
		}
		now := time.Now()
		date := now.Format("20060102")
		destinationFile := ""
		if dirExists == true {
			for prefix := 1; prefix < 100; prefix++ {
				destinationFile = olddir + "/" + sourceFile + "." + date + "-" + fmt.Sprintf("%02d", prefix)
				fileExists, err := file_existence(destinationFile)
				if err != nil {
					cli.NewExitError("Unable to get destination file location", 1)
				}
				if fileExists == false {
					break
				}
			}
		} else {
			os.Mkdir(olddir, 0755)
			destinationFile = olddir + "/" + sourceFile + "." + date + "-01"
		}

		b, err := ioutil.ReadFile(sourceFile)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(destinationFile, b, 0644)
		if err != nil {
			panic(err)
		}
		return nil
	}

	app.Run(os.Args)
}
