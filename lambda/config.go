package lambda

import (
	"context"
	"errors"
	"log"
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
const configPath = ".transport/config.yml"

func Configure(stage string) error {
	Logger.Debug("InitConfig", "stage", stage)
	ReadConfig(configPath)
	sourceProfile := Cfg.Cfg[stage].Source.Profile
	// test if exits and more then 3 chars
	if len(sourceProfile) < 3 {
		log.Fatalf("source profile name is too short %v ", sourceProfile)
		return errors.New("source profile name is too short")
	}
	Logger.Debug("Source Profile ", "sourceprofile", sourceProfile)
	//Clients
	sourceCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(sourceProfile))
	if err != nil {
		log.Fatalf("source configuration error %v ", err.Error())
		return err
	}
	targetProfile := Cfg.Cfg[stage].Target.Profile
	if len(targetProfile) < 3 {
		log.Fatalf("source profile name is too short %v ", targetProfile)
		return errors.New("source profile name is too short")
	}
	Logger.Debug("Target Profile ", "targetprofile", targetProfile)
	targetCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(targetProfile))
	if err != nil {
		log.Fatalf("target configuration error %v ", err.Error())
		return err
	}
	// Create a new Lambda client with the source profile
	SourceClient = lambda.NewFromConfig(sourceCfg)
	// Create a new Lambda client with the target profile
	TargetClient = lambda.NewFromConfig(targetCfg)
	return nil
}

func ReadConfig(filename string) (*Config, error) {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Logger.Error("File does not exist", "filename", filename)
	}
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
	Logger.Debug("Config file read successfully", "filename", filename)
	return &cfg, nil
}
