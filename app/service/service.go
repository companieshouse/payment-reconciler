package service

import (
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/dao"
	"github.com/companieshouse/payment-reconciler/app/filetransfer"
	"github.com/companieshouse/payment-reconciler/app/models"
)

type Service struct {
	Config       *config.Config
	DAO          dao.DAO
	FileTransfer filetransfer.FileTransfer
}

func New(cfg *config.Config) *Service {

	return &Service{
		Config:       cfg,
		DAO:          dao.New(cfg),
		FileTransfer: filetransfer.New(cfg),
	}
}

func (svc *Service) HandleRequest(reconciliation *models.Reconciliation) error {

	log.Info("Payment reconciliation lambda executing. Fetching reconciliation data.")

	reconciliationData, err := svc.DAO.GetReconciliationData()
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Successfully retrieved reconciliation data. Preparing to upload CSV.")
	log.Trace("Reconciliation data", log.Data{"Reconciliation data": reconciliationData})

	err = svc.FileTransfer.UploadCSV(reconciliationData.ToCSV())
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("Successfully uploaded CSV. Execution complete.")
	return nil
}
