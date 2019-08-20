package service

import (
	"errors"
	"testing"

	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/dao"
	"github.com/companieshouse/payment-reconciler/models"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

const expectedTransactionsFileNamePrefix string = "CHS_PaymentTransactions_"
const expectedProductsFileNamePrefix string = "CHS_PaymentProducts_"
const expectedCSVFileSuffix = ".csv"
const reconciliationDate string = "2019-01-01"

func createMockService(cfg *config.Config, mockDao *dao.MockDAO) *ServiceImpl {

	return &ServiceImpl{
		Config: cfg,
		DAO:    mockDao,
	}
}

func TestUnitGetTransactionsCSV(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.Config{}
	reconciliationMetaData := models.ReconciliationMetaData{
		ReconciliationDate: reconciliationDate,
	}

	Convey("Subject: Success", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given transactions data is successfully fetched", func() {

			var transactions models.TransactionsList
			mockDao.EXPECT().GetTransactionsData(&reconciliationMetaData).Return(transactions, nil).Times(1)

			Convey("Then no errors are returned", func() {

				transactionsCSV, err := svc.GetTransactionsCSV(&reconciliationMetaData)
				So(err, ShouldBeNil)

				Convey("And a CSV is successfully constructed", func() {

					So(transactionsCSV, ShouldNotBeNil)
					So(transactionsCSV.Data, ShouldResemble, transactions)
					So(transactionsCSV.FileName, ShouldEqual, expectedTransactionsFileNamePrefix+reconciliationMetaData.ReconciliationDate+expectedCSVFileSuffix)
				})
			})
		})
	})

	Convey("Subject: Failure to retrieve transactions data", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given an error when fetching transactions data", func() {

			var transactions models.TransactionsList
			mockDao.EXPECT().GetTransactionsData(&reconciliationMetaData).Return(transactions, errors.New("failure to fetch transactions data")).Times(1)

			Convey("Then errors are returned", func() {

				transactionsCSV, err := svc.GetTransactionsCSV(&reconciliationMetaData)
				So(err, ShouldNotBeNil)

				Convey("And no CSV is constructed", func() {

					So(transactionsCSV.Data, ShouldBeNil)
				})
			})
		})
	})
}

func TestUnitGetProductsCSV(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg := config.Config{}
	reconciliationMetaData := models.ReconciliationMetaData{
		ReconciliationDate: reconciliationDate,
	}

	Convey("Subject: Success", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given products data is successfully fetched", func() {

			var products models.ProductsList
			mockDao.EXPECT().GetProductsData(&reconciliationMetaData).Return(products, nil).Times(1)

			Convey("Then no errors are returned", func() {

				productsCSV, err := svc.GetProductsCSV(&reconciliationMetaData)
				So(err, ShouldBeNil)

				Convey("And a CSV is successfully constructed", func() {

					So(productsCSV, ShouldNotBeNil)
					So(productsCSV.Data, ShouldResemble, products)
					So(productsCSV.FileName, ShouldEqual, expectedProductsFileNamePrefix+reconciliationMetaData.ReconciliationDate+expectedCSVFileSuffix)
				})
			})
		})
	})

	Convey("Subject: Failure to retrieve products data", t, func() {

		mockDao := dao.NewMockDAO(mockCtrl)

		svc := createMockService(&cfg, mockDao)

		Convey("Given an error when fetching products data", func() {

			var products models.ProductsList
			mockDao.EXPECT().GetProductsData(&reconciliationMetaData).Return(products, errors.New("failure to fetch transactions data")).Times(1)

			Convey("Then errors are returned", func() {

				productsCSV, err := svc.GetProductsCSV(&reconciliationMetaData)
				So(err, ShouldNotBeNil)

				Convey("And no CSV is constructed", func() {

					So(productsCSV.Data, ShouldBeNil)
				})
			})
		})
	})
}
