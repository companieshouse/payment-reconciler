package filetransfer

import (
	"encoding/csv"
	"fmt"
	"net"
	"path/filepath"

	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTP struct {
	Config          *config.Config
	SSHClientConfig *ssh.ClientConfig
}

func New(cfg *config.Config) *SFTP {

	sshCfg := &ssh.ClientConfig{
		User: cfg.SFTPUserName,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Auth: []ssh.AuthMethod{
			ssh.Password(cfg.SFTPPassword),
		},
	}

	sshCfg.SetDefaults()

	return &SFTP{
		Config:          cfg,
		SSHClientConfig: sshCfg,
	}
}

func (t *SFTP) UploadCSV(csvData [][]string) error {

	log.Info("Starting CSV upload. Initiating SSH connection to " + t.Config.SFTPServer)

	client, err := ssh.Dial("tcp", t.Config.SFTPServer+":22", t.SSHClientConfig)
	if err != nil {
		return fmt.Errorf("Failed to establish connection: %s", err)
	}
	defer client.Close()

	sftpSession, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("Error creating SFTP session: %s", err)
	}
	defer sftpSession.Close()

	log.Info("Connection established. Writing CSV")

	file, err := sftpSession.Create(filepath.Join(t.Config.SFTPFilePath, filepath.Base("test.csv")))
	if err != nil {
		return fmt.Errorf("Failed to create CSV: %s", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)

	if err := w.WriteAll(csvData); err != nil {
		return fmt.Errorf("Error writing CSV data: %s", err)
	}

	return nil
}
