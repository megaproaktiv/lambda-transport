package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	mylambda "github.com/tecracer/lambda-transport/lambda"
)

var Logger *slog.Logger

const (
	LevelDebug = slog.Level(-4)
	LevelInfo  = slog.Level(0)
	LevelWarn  = slog.Level(4)
	LevelError = slog.Level(8)
)

const NO_STAGE = 1

func main() {
	stagep := flag.String("stage", "", "define the application stage")
	helpFlag := flag.Bool("help", false, "Display help message")
	verboseFlag := flag.Bool("verbose", false, "Display more messages")
	// Parse the flags.
	flag.Parse()

	if *helpFlag || len(os.Args) == 1 {
		displayHelpMessage()
		return
	}

	if len(*stagep) == 0 {
		flag.PrintDefaults()
		os.Exit(NO_STAGE)
	}
	stage := *stagep
	// // Configuration

	configfile := ".transport/config.yml"
	if *verboseFlag {
		fmt.Println("Using config file: ", configfile)
	}
	cfg, err := mylambda.ReadConfig(configfile)
	if err != nil {
		Logger.Error("Error reading configuration file, here is why", "error", err, "filename", configfile)
	}

	// download source lambda code
	sourceLambdaName := cfg.Cfg[stage].Source.LambdaName
	if *verboseFlag {
		fmt.Println("Source Lambda Name: ", sourceLambdaName)
	}
	mylambda.DownLoadSourceLambdaCode(mylambda.SourceClient, sourceLambdaName)
	// upload target lambda code
	targetLambdaName := cfg.Cfg[stage].Target.LambdaName
	if len(targetLambdaName) == 0 {
		Logger.Error("Target Lambda Name is empty")
	}
	if *verboseFlag {
		fmt.Println("Target Lambda Name: ", targetLambdaName)
	}
	mylambda.UploadTargetLambdaCode(mylambda.TargetClient, targetLambdaName)
}

func displayHelpMessage() {
	fmt.Println(`App to transport Lambda code from one account to another.
Config file is .transport/config.yml in cuurent directory.
See config-example.yml.
Usage:
  -help: Display this help message.
  -stage: Define the application stage.
  -verbose: Display more messages.
example: ./lambda-transport -stage dev
  this would use the configuration for the dev stage.
  `)
}
