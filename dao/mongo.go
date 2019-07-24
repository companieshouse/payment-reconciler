package dao

import (
	"fmt"

	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Mongo provides a MongoDB implementation of the DAO
type Mongo struct {
	Config *config.Config
}

// New returns a new Mongo struct using the provided config
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

// GetTransactionsData fetches transactions data
func (m *Mongo) GetTransactionsData() (models.TransactionsList, error) {

	var transactions []models.Transaction

	var transactionsData models.TransactionsList

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return transactionsData, fmt.Errorf("error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(m.Config.Database).C(m.Config.TransactionsCollection).Find(bson.M{}).All(&transactions)
	if err != nil {
		return transactionsData, fmt.Errorf("error retrieving transactions data: %s", err)
	}

	transactionsData = models.TransactionsList{
		Transactions: transactions,
	}

	return transactionsData, err
}

// GetProductsData fetches products data
func (m *Mongo) GetProductsData() (models.ProductsList, error) {

	var products []models.Product

	var productsData models.ProductsList

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return productsData, fmt.Errorf("error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(m.Config.Database).C(m.Config.ProductsCollection).Find(bson.M{}).All(&products)
	if err != nil {
		return productsData, fmt.Errorf("error retrieving products data: %s", err)
	}

	productsData = models.ProductsList{
		Products: products,
	}

	return productsData, err
}
