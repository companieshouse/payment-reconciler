package filetransfer

import (
	"encoding/csv"
	"fmt"
	"net"
	"path/filepath"

	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/models"
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

// UploadCSVFiles uploads an array of CSV's to an STFP server
func (t *SFTP) UploadCSVFiles(csvs []models.CSV) error {

	log.Info("Starting upload of CSV's. Initiating SSH connection to " + t.Config.SFTPServer)

	client, err := ssh.Dial("tcp", t.Config.SFTPServer+":22", t.SSHClientConfig)
	if err != nil {
		return fmt.Errorf("failed to establish connection: %s", err)
	}
	defer client.Close()

	sftpSession, err := sftp.NewClient(client)
	if err != nil {
		return fmt.Errorf("error creating SFTP session: %s", err)
	}
	defer sftpSession.Close()

	log.Info("Connection established. Writing CSV's")

	for i := 0; i < len(csvs); i++ {

		file, err := sftpSession.Create(filepath.Join(t.Config.SFTPFilePath, filepath.Base(csvs[i].FileName)))
		if err != nil {
			return fmt.Errorf("failed to create CSV: %s", err)
		}

		w := csv.NewWriter(file)

		if err := w.WriteAll(csvs[i].Data.ToCSV()); err != nil {
			return fmt.Errorf("error writing CSV data: %s", err)
		}

		file.Close()
	}

	return nil
}
