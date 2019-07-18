package dao

import (
	"fmt"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
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
func (m *Mongo) GetTransactionsData() (models.TransactionsData, error) {

	var transactions []models.Transaction

	var transactionsData models.TransactionsData

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return transactionsData, fmt.Errorf("error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(m.Config.Database).C(m.Config.TransactionsCollection).Find(bson.M{}).All(&transactions)
	if err != nil {
		return transactionsData, fmt.Errorf("error retrieving transactions data: %s", err)
	}

	transactionsData = models.TransactionsData{
		Transactions: transactions,
	}

	return transactionsData, err
}

// GetProductsData fetches products data
func (m *Mongo) GetProductsData() (models.ProductsData, error) {

	var products []models.Product

	var productsData models.ProductsData

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return productsData, fmt.Errorf("error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	err = mongoSession.DB(m.Config.Database).C(m.Config.ProductsCollection).Find(bson.M{}).All(&products)
	if err != nil {
		return productsData, fmt.Errorf("error retrieving products data: %s", err)
	}

	productsData = models.ProductsData{
		Products: products,
	}

	return productsData, err
}
