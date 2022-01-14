package dao

import (
	"testing"

	"github.com/companieshouse/payment-reconciler/models"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/companieshouse/payment-reconciler/config"
	. "github.com/smartystreets/goconvey/convey"
)

var cfg *config.Config

func TestUnitGetTransactionsData(t *testing.T) {
	Convey("Get Transaction Data", t, func() {
		if cfg == nil {
			cfg, _ = config.Get()
		}

		client = &mongo.Client{}

		dao := New(cfg)
		reconciliationMetaData := models.ReconciliationMetaData{}

		transactionsList, err := dao.GetTransactionsData(&reconciliationMetaData)
		So(transactionsList, ShouldBeZeroValue)
		So(err.Error(), ShouldEqual, "error retrieving transactions data: the Find operation must have a Deployment set before Execute can be called")
	})
}

func TestUnitGetProductsData(t *testing.T) {
	Convey("Get Products Data", t, func() {
		if cfg == nil {
			cfg, _ = config.Get()
		}
		client = &mongo.Client{}

		dao := New(cfg)
		reconciliationMetaData := models.ReconciliationMetaData{}

		productsList, err := dao.GetProductsData(&reconciliationMetaData)
		So(productsList, ShouldBeZeroValue)
		So(err.Error(), ShouldEqual, "error retrieving products data: the Find operation must have a Deployment set before Execute can be called")
	})
}

func TestUnitGetRefundsData(t *testing.T) {
	Convey("Get Refunds Data", t, func() {
		if cfg == nil {
			cfg, _ = config.Get()
		}
		client = &mongo.Client{}

		dao := New(cfg)
		reconciliationMetaData := models.ReconciliationMetaData{}

		refundsList, err := dao.GetRefundsData(&reconciliationMetaData)
		So(refundsList, ShouldBeZeroValue)
		So(err.Error(), ShouldEqual, "error retrieving refunds data: the Find operation must have a Deployment set before Execute can be called")
	})
}
