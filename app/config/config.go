package config

import (
	"github.com/companieshouse/gofigure"
)

// Config holds configuration details required to execute the lambda.
type Config struct {
	Collection   string `env:"MONGODB_COLLECTION"    flag:"mongodb-collection"    flagDesc:"MongoDB collection for data"`
	Database     string `env:"MONGODB_DATABASE"      flag:"mongodb-database"      flagDesc:"MongoDB database for data"`
	MongoDBURL   string `env:"MONGODB_URL"           flag:"mongodb-url"           flagDesc:"MongoDB server URL"`
	SFTPServer   string `env:"SFTP_SERVER"           flag:"sftp-server"           flagDesc:"Name of the SFTP server"`
	SFTPUserName string `env:"SFTP_USERNAME"         flag:"sftp-username"         flagDesc:"Username of SFTP server"`
	SFTPPassword string `env:"SFTP_PASSWORD"         flag:"sftp-password"         flagDesc:"Password of SFTP server"`
	SFTPFilePath string `env:"SFTP_FILE_PATH"        flag:"sftp-file-path"        flagDesc:"File path on the SFTP server"`
}

// Get returns configuration details marshalled into a Config struct
func Get() (*Config, error) {

	cfg := &Config{}

	err := gofigure.Gofigure(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
