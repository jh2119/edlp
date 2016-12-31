// EDLP - https://github.com/jh2119/edlp
//
package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"time"
)

// Build and Version should be assigned the software's version information
// based off of GIT data by the build environment
var Version = "n/a"
var Build = "n/a"

func main() {

	var fileHandleInput *os.File
	var fileHandleOutput *os.File
	var pathDestination string
	var err error

	// process commandline arguments
	inputFilename, outputFilename, pathDestination := commandline(os.Args)

	// check input file conditions
	if len(inputFilename) == 0 {
		fmt.Println("An input file must be provided.")
		return
	}

	// open input file
	fileHandleInput, err = os.Open(inputFilename)

	if err != nil {
		fmt.Println("Error opening source file: " + inputFilename)
		return
	}

	// open output file
	if len(outputFilename) == 0 {		
		outputFilename = generateFilenameOutput(inputFilename, pathDestination)
	}

	fileHandleOutput, err = os.Open(outputFilename)

	if err != nil {
		fmt.Println("Error opening destination file: " + outputFilename)
		return
	}

	_ = fileHandleInput
	_ = fileHandleOutput

	// TODO: open target output file
	// TODO: parse
	// TODO: save data tree
}

// commandLine uses urfave's cli to get input and output file information cli args
//
func commandline(args []string) (string, string, string) {
	var filenameInput string
	var filenameOutput string
	var pathOutput string

	app := cli.NewApp()
	app.Name = "EDLP"
	app.Version = Build //Version
	app.Compiled = time.Now()
	app.Usage = "the insatiable design document processor"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "John Hall",
			//Email: "...",
		},
	}
	app.Copyright = "(c) 2016 John Hall"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "destination, d",
			Usage:       "Specify the destination path",
			Destination: &pathOutput,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Specify an output filename",
			Destination: &filenameOutput,
		},
	}

	app.Action = func(c *cli.Context) error {
		filenameInput = c.Args().Get(0)
		return nil
	}

	app.Run(args)

	return filenameInput, filenameOutput, pathOutput
}
