package service

import (
	"errors"
	"testing"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/dao"
	"github.com/companieshouse/payment-reconciler/app/filetransfer"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func createMockService(mockDao *dao.MockDAO, cfg *config.Config, mockFileTransfer *filetransfer.MockFileTransfer) *Service {

	return &Service{
		DAO:          mockDao,
		Config:       cfg,
		FileTransfer: mockFileTransfer,
	}
}

func TestUnitHandleRequest(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := &config.Config{}
	reconciliation := &models.Reconciliation{}

	Convey("Success path", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		mockService := createMockService(mockDao, cfg, mockFileTransfer)

		Convey("Given reconciliation data is retrieved successfully", func() {

			mockDao.EXPECT().GetReconciliationData().Return(&models.ReconciliationData{}, nil).Times(1)

			Convey("And a CSV is uploaded successfully", func() {

				mockFileTransfer.EXPECT().UploadCSV(gomock.Any()).Return(nil).Times(1)

				Convey("Then the request is successful", func() {

					err := mockService.HandleRequest(reconciliation)
					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Error fetching reconciliation data", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		mockService := createMockService(mockDao, cfg, mockFileTransfer)

		Convey("Given there's an error when fetching reconciliation data", func() {

			mockDao.EXPECT().GetReconciliationData().Return(nil, errors.New("Error retrieving reconciliation data")).Times(1)

			Convey("Then the call to upload a CSV is never made", func() {

				mockFileTransfer.EXPECT().UploadCSV(gomock.Any()).Times(0)

				Convey("And the request is unsuccessful", func() {

					err := mockService.HandleRequest(reconciliation)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})

	Convey("Error uploading CSV", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		mockService := createMockService(mockDao, cfg, mockFileTransfer)

		Convey("Given reconciliation data is retrieved successfully", func() {

			mockDao.EXPECT().GetReconciliationData().Return(&models.ReconciliationData{}, nil).Times(1)

			Convey("But there's an error when uploading a CSV", func() {

				mockFileTransfer.EXPECT().UploadCSV(gomock.Any()).Return(errors.New("Error uploading CSV")).Times(1)

				Convey("Then the request is unsuccessful", func() {

					err := mockService.HandleRequest(reconciliation)
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
