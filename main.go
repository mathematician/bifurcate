package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mathematician/bifurcate/tfstate"
	"github.com/mathematician/bifurcate/version"

	"github.com/sirupsen/logrus"
)

const (
	// BANNER is what is printed for help/info output
	BANNER = `
 ______  _____ _______ _     _  ______ _______ _______ _______ _______
 |_____]   |   |______ |     | |_____/ |       |_____|    |    |______
 |_____] __|__ |       |_____| |    \_ |_____  |     |    |    |______
                                                                      
Tool to generate bifurcations between aws account and terraform state
Version: %s
`
)

var (
	region string

	debug bool
	vrsn  bool
)

func init() {
	flag.StringVar(&region, "region", "", "aws region")

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

	if region != "" {
		os.Setenv("AWS_REGION", region)
	}
}

func main() {
	s3Bucket := flag.Args()[0]
	fmt.Printf("Bucket where state files are stored: %s\n", s3Bucket)

	key := "operations/bastion/terraform.tfstate"

	tfstateResources, err := tfstate.GetResources(s3Bucket, key)
	if err != nil {
		panic("Error, " + err.Error())
	}

	fmt.Printf("Resources: \n")
	for _, resource := range tfstateResources {
		fmt.Printf("%+v\n", resource)
	}
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
