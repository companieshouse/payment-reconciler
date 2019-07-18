package lambda

import (
	"errors"
	"testing"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/filetransfer"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/companieshouse/payment-reconciler/app/service"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func createMockLambda(cfg *config.Config, mockService *service.MockService, mockFileTransfer *filetransfer.MockFileTransfer) *Lambda {

	return &Lambda{
		Config:       cfg,
		Service:      mockService,
		FileTransfer: mockFileTransfer,
	}
}

func TestUnitExecute(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.Config{}
	reconciliationMetaData := models.ReconciliationMetaData{}

	Convey("Subject: Success", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("And a products CSV is constructed successfully", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, nil).Times(1)

				Convey("And the CSV's are uploaded successfully", func() {

					csvs := []models.CSV{transactionsCSV, productsCSV}
					mockFileTransfer.EXPECT().UploadCSVFiles(csvs).Return(nil).Times(1)

					Convey("Then the request is successful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Subject: Failure to construct transactions CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a failure when constructing a transactions CSV", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, errors.New("Failed to construct transactions CSV")).Times(1)

			Convey("Then there is never an attempt to construct a products CSV", func() {

				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Times(0)

				Convey("And no CSV's are uploaded", func() {

					mockFileTransfer.EXPECT().UploadCSVFiles(gomock.Any()).Times(0)

					Convey("And the request is unsuccessful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})

	Convey("Subject: Failure to construct products CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("But there's a failure when constructing a products CSV", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, errors.New("Failed to construct products CSV")).Times(1)

				Convey("Then no CSV's are uploaded", func() {

					mockFileTransfer.EXPECT().UploadCSVFiles(gomock.Any()).Times(0)

					Convey("And the request is unsuccessful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})

	Convey("Subject: Failure to upload CSV's", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("And a products CSV is constructed successfully", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, nil).Times(1)

				Convey("But the CSV's are not uploaded successfully", func() {

					csvs := []models.CSV{transactionsCSV, productsCSV}
					mockFileTransfer.EXPECT().UploadCSVFiles(csvs).Return(errors.New("Failure to upload CSV's")).Times(1)

					Convey("Then the request is unsuccessful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})
}
