package lambda

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"gopkg.in/yaml.v2"
)

type StageConfig struct {
	Profile    string `yaml:"profile"`
	Region     string `yaml:"region"`
	LambdaName string `yaml:"lambda"`
}

type TransportPair struct {
	Source StageConfig `yaml:"source"`
	Target StageConfig `yaml:"target"`
}

type Config struct {
	Cfg map[string]TransportPair `yaml:"config"`
}

var Cfg *Config

// Use one init fucntion to initialize all the configurations
// Using several init functions would lead to non-deteministic behavior
func init() {

	// LOGGER
	handler := slog.NewTextHandler(os.Stdout,
		&slog.HandlerOptions{Level: LevelDebug})
	Logger = slog.New(handler)

	//Clients
	sourceCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(sourceProfile))
	if err != nil {
		log.Fatalf("source configuration error %v ", err.Error())
	}
	targetCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(targetProfile))
	if err != nil {
		log.Fatalf("target configuration error %v ", err.Error())

	}
	// Create a new Lambda client with the source profile
	SourceClient = lambda.NewFromConfig(sourceCfg)
	// Create a new Lambda client with the target profile
	TargetClient = lambda.NewFromConfig(targetCfg)
}

func ReadConfig(filename string) (*Config, error) {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		Logger.Error("Error reading YAML file, here is why ", "error", err)
		return nil, err
	}
	var cfg Config
	stages := make(map[string]TransportPair)
	cfg.Cfg = stages
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		Logger.Error("Error parsing YAML file, here is why", "error", err)
	}
	return &cfg, nil
}
