package service

import (
	"errors"
	"testing"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/dao"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

const expectedPaymentTransactionsFileNamePrefix string = "CHS_PaymentTransactions_"
const expectedPaymentProductsFileNamePrefix string = "CHS_PaymentProducts_"
const expectedCSVFileSuffix = ".csv"
const reconciliationDate string = "2019-01-01"

func createMockService(cfg *config.Config, mockDao *dao.MockDAO) *ServiceImpl {

	return &ServiceImpl{
		Config: cfg,
		DAO:    mockDao,
	}
}

func TestUnitGetPaymentTransactionsCSV(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.Config{}
	reconciliationMetaData := models.ReconciliationMetaData{
		Date: reconciliationDate,
	}

	Convey("Subject: Success", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given payment transactions data is successfully fetched", func() {

			var paymentTransactions models.PaymentTransactionsData
			mockDao.EXPECT().GetPaymentTransactionsData().Return(paymentTransactions, nil).Times(1)

			Convey("Then a CSV is successfully constructed", func() {

				paymentTransactionsCSV, err := svc.GetPaymentTransactionsCSV(&reconciliationMetaData)
				So(paymentTransactionsCSV, ShouldNotBeNil)
				So(paymentTransactionsCSV.Data, ShouldResemble, paymentTransactions)
				So(paymentTransactionsCSV.FileName, ShouldEqual, expectedPaymentTransactionsFileNamePrefix+reconciliationMetaData.Date+expectedCSVFileSuffix)

				Convey("And no errors are returned", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Subject: Failure to retrieve payment transactions data", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given an error when fetching payment transactions data", func() {

			var paymentTransactions models.PaymentTransactionsData
			mockDao.EXPECT().GetPaymentTransactionsData().Return(paymentTransactions, errors.New("Failure to fetch payment transactions data")).Times(1)

			Convey("Then no CSV is constructed", func() {

				paymentTransactionsCSV, err := svc.GetPaymentTransactionsCSV(&reconciliationMetaData)
				So(paymentTransactionsCSV.Data, ShouldBeNil)

				Convey("And errors are returned", func() {

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestUnitGetPaymentProductsCSV(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.Config{}
	reconciliationMetaData := models.ReconciliationMetaData{
		Date: reconciliationDate,
	}

	Convey("Subject: Success", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given payment products data is successfully fetched", func() {

			var paymentProducts models.PaymentProductsData
			mockDao.EXPECT().GetPaymentProductsData().Return(paymentProducts, nil).Times(1)

			Convey("Then a CSV is successfully constructed", func() {

				paymentProductsCSV, err := svc.GetPaymentProductsCSV(&reconciliationMetaData)
				So(paymentProductsCSV, ShouldNotBeNil)
				So(paymentProductsCSV.Data, ShouldResemble, paymentProducts)
				So(paymentProductsCSV.FileName, ShouldEqual, expectedPaymentProductsFileNamePrefix+reconciliationMetaData.Date+expectedCSVFileSuffix)

				Convey("And no errors are returned", func() {

					So(err, ShouldBeNil)
				})
			})
		})
	})

	Convey("Subject: Failure to retrieve payment products data", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given an error when fetching payment products data", func() {

			var paymentProducts models.PaymentProductsData
			mockDao.EXPECT().GetPaymentProductsData().Return(paymentProducts, errors.New("Failure to fetch payment transactions data")).Times(1)

			Convey("Then no CSV is constructed", func() {

				paymentProductsCSV, err := svc.GetPaymentProductsCSV(&reconciliationMetaData)
				So(paymentProductsCSV.Data, ShouldBeNil)

				Convey("And errors are returned", func() {

					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
