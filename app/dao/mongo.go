package dao

import (
	"fmt"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

const paymentReconciliationDB string = "payment_reconciliation"
const paymentTransactionsCollection = "payment_transactions"
const paymentProductsCollection string = "payment_products"

type Mongo struct {
	Config *config.Config
}

func New(cfg *config.Config) *Mongo {

	return &Mongo{
		Config: cfg,
	}
}

// getMongoSession retrieves a fresh MongoDB session
func getMongoSession(cfg *config.Config) (*mgo.Session, error) {

	s, err := mgo.Dial(cfg.MongoDBURL)
	if err != nil {
		return nil, err
	}

	return s.Copy(), nil
}

// GetPaymentTransactionsData fetches payment transactions data
func (m *Mongo) GetPaymentTransactionsData() (models.PaymentTransactionsData, error) {

	var paymentTransactions []models.PaymentTransaction

	var paymentTransactionsData models.PaymentTransactionsData

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return paymentTransactionsData, fmt.Errorf("Error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(paymentReconciliationDB).C(paymentTransactionsCollection).Find(bson.M{}).All(&paymentTransactions)
	if err != nil {
		return paymentTransactionsData, fmt.Errorf("Error retrieving payment transactions data: %s", err)
	}

	paymentTransactionsData = models.PaymentTransactionsData{
		PaymentTransactions: paymentTransactions,
	}

	return paymentTransactionsData, err
}

// GetPaymentProductsData fetches payment transactions data
func (m *Mongo) GetPaymentProductsData() (models.PaymentProductsData, error) {

	var paymentProducts []models.PaymentProduct

	var paymentProductsData models.PaymentProductsData

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return paymentProductsData, fmt.Errorf("Error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(paymentReconciliationDB).C(paymentProductsCollection).Find(bson.M{}).All(&paymentProducts)
	if err != nil {
		return paymentProductsData, fmt.Errorf("Error retrieving payment products data: %s", err)
	}

	paymentProductsData = models.PaymentProductsData{
		PaymentProducts: paymentProducts,
	}

	return paymentProductsData, err
}
