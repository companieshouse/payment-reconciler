package config

import (
	"github.com/companieshouse/gofigure"
)

// Config holds configuration details required to execute the lambda.
type Config struct {
	TransactionsCollection string `env:"MONGODB_PAYMENT_REC_TRANSACTIONS_COLLECTION"  flag:"mongodb-payment-rec-transactions-collection"   flagDesc:"MongoDB collection for payment transactions data"`
	ProductsCollection     string `env:"MONGODB_PAYMENT_REC_PRODUCTS_COLLECTION"      flag:"mongodb-payment-rec-products-collection"       flagDesc:"MongoDB collection for payment products data"`
	RefundsCollection      string `env:"MONGODB_PAYMENT_REC_REFUNDS_COLLECTION"       flag:"mongodb-payment-rec-refunds-collection"        flagDesc:"MongoDB collection for payment refunds data"`
	Database               string `env:"MONGODB_PAYMENT_REC_DATABASE"                 flag:"mongodb-payment-rec-database"                  flagDesc:"MongoDB database for payment reconciliation data"`
	MongoDBURL             string `env:"MONGODB_URL"                                  flag:"mongodb-url"                                   flagDesc:"MongoDB server URL" json:"-"`
	SFTPServer             string `env:"SFTP_SERVER"                                  flag:"sftp-server"                                   flagDesc:"Name of the SFTP server" json:"-"`
	SFTPPort               string `env:"SFTP_PORT"                                    flag:"sftp-port"                                     flagDesc:"Port to connect to the SFTP server"`
	SFTPUserName           string `env:"SFTP_USERNAME"                                flag:"sftp-username"                                 flagDesc:"Username of SFTP server" json:"-"`
	SFTPPassword           string `env:"SFTP_PASSWORD"                                flag:"sftp-password"                                 flagDesc:"Password of SFTP server" json:"-"`
	SFTPFilePath           string `env:"SFTP_FILE_PATH"                               flag:"sftp-file-path"                                flagDesc:"File path on the SFTP server" json:"-"`
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
