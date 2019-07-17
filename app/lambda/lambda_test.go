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

		Convey("Given a payment transactions CSV is constructed successfully", func() {

			var paymentTransactionsCSV models.CSV
			mockService.EXPECT().GetPaymentTransactionsCSV(&reconciliationMetaData).Return(paymentTransactionsCSV, nil).Times(1)

			Convey("And a payment products CSV is constructed successfully", func() {

				var paymentProductsCSV models.CSV
				mockService.EXPECT().GetPaymentProductsCSV(&reconciliationMetaData).Return(paymentProductsCSV, nil).Times(1)

				Convey("And the CSV's are uploaded successfully", func() {

					csvs := []models.CSV{paymentTransactionsCSV, paymentProductsCSV}
					mockFileTransfer.EXPECT().UploadCSVFiles(csvs).Return(nil).Times(1)

					Convey("Then the request is successful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Subject: Failure to construct payment transactions CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a failure when constructing a payment transactions CSV", func() {

			var paymentTransactionsCSV models.CSV
			mockService.EXPECT().GetPaymentTransactionsCSV(&reconciliationMetaData).Return(paymentTransactionsCSV, errors.New("Failed to construct payment transactions CSV")).Times(1)

			Convey("Then there is never an attempt to construct a payment products CSV", func() {

				mockService.EXPECT().GetPaymentProductsCSV(&reconciliationMetaData).Times(0)

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

	Convey("Subject: Failure to construct payment products CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a payment transactions CSV is constructed successfully", func() {

			var paymentTransactionsCSV models.CSV
			mockService.EXPECT().GetPaymentTransactionsCSV(&reconciliationMetaData).Return(paymentTransactionsCSV, nil).Times(1)

			Convey("But there's a failure when constructing a payment products CSV", func() {

				var paymentProductsCSV models.CSV
				mockService.EXPECT().GetPaymentProductsCSV(&reconciliationMetaData).Return(paymentProductsCSV, errors.New("Failed to construct payment products CSV")).Times(1)

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

		Convey("Given a payment transactions CSV is constructed successfully", func() {

			var paymentTransactionsCSV models.CSV
			mockService.EXPECT().GetPaymentTransactionsCSV(&reconciliationMetaData).Return(paymentTransactionsCSV, nil).Times(1)

			Convey("And a payment products CSV is constructed successfully", func() {

				var paymentProductsCSV models.CSV
				mockService.EXPECT().GetPaymentProductsCSV(&reconciliationMetaData).Return(paymentProductsCSV, nil).Times(1)

				Convey("But the CSV's are not uploaded successfully", func() {

					csvs := []models.CSV{paymentTransactionsCSV, paymentProductsCSV}
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
