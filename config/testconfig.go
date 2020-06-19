package config

import (
	"github.com/companieshouse/gofigure"
)

// TestConfig holds configuration details required to test the lambda.
type TestConfig struct {
	EnvRegionAWS string `env:"ENV_REGION_AWS"   flag:"env-region-aws"   flagDesc:"The AWS region code for the lambda"`
	LambdaName   string `env:"LAMBDA_NAME"      flag:"lambda-name"      flagDesc:"The name of the lambda"`
	SFTPConfig
}

// GetTestConfig returns configuration details marshalled into a Config struct
func GetTestConfig() (*TestConfig, error) {

	cfg := &TestConfig{}

	err := gofigure.Gofigure(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// NotPopulated returns a boolean indicating whether any test configs are not populated
func (cfg *TestConfig) NotPopulated() bool {

	return cfg.LambdaName == "" ||
		cfg.EnvRegionAWS == "" ||
		cfg.SFTPFilePath == "" ||
		cfg.SFTPPort == "" ||
		cfg.SFTPPassword == "" ||
		cfg.SFTPServer == "" ||
		cfg.SFTPUserName == ""
}
