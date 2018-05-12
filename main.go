package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mathematician/bifurcate/version"
	"github.com/sirupsen/logrus"
)

const (
	BANNER = `
 ______  _____ _______ _     _  ______ _______ _______ _______ _______
 |_____]   |   |______ |     | |_____/ |       |_____|    |    |______
 |_____] __|__ |       |_____| |    \_ |_____  |     |    |    |______
                                                                      
Tool to generate bifurcations between aws account and terraform state
Version: %s
`
)

var (
	debug bool
	vrsn  bool
)

func init() {
	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&debug, "debug", false, "run in debug")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("bifurcate version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		printUsageAndExit("No arguments supplied.", 1)
	}

	arg := flag.Args()[0]

	if arg == "help" {
		printUsageAndExit("", 0)
	}

	if arg == "version" {
		fmt.Printf("bifurcate version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	s3Bucket := flag.Args()[0]
	fmt.Printf("Bucket where state files are stored: %s", s3Bucket)
}

func printUsageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
