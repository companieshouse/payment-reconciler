package dao

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/models"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

var client *mongo.Client

const deadline = 5 * time.Second

// MongoDatabaseInterface is an interface that describes the mongodb driver
type MongoDatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

func getDB(mongoDBURL, databaseName string) MongoDatabaseInterface {
	return getMongoClient(mongoDBURL).Database(databaseName)
}

func getMongoClient(mongoDBURL string) *mongo.Client {
	if client != nil {
		return client
	}

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(ctx, clientOptions)

	// Assume the caller of this func cannot handle the case where there is no database connection
	// so the service must crash here as it cannot continue.
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Check we can connect to the mongodb instance. Failure here should result in a crash.
	pingContext, cancel := context.WithDeadline(ctx, time.Now().Add(deadline))

	err = client.Ping(pingContext, nil)
	if err != nil {
		log.Error(errors.New("ping to mongodb timed out. please check the connection to mongodb and that it is running"))
		os.Exit(1)
	}

	defer cancel()

	log.Info("connected to mongodb successfully")

	return client
}

// GetTransactionsData fetches transactions data
func (m *Mongo) GetTransactionsData(reconciliationMetaData *models.ReconciliationMetaData) (models.TransactionsList, error) {

	var transactions []models.Transaction
	var transactionsData models.TransactionsList

	collection := getMongoClient(m.Config.Database).Database(m.Config.Database).Collection(m.Config.TransactionsCollection)

	cur, err := collection.Find(context.Background(), bson.M{"transaction_date": bson.M{
		"$gt": reconciliationMetaData.StartTime,
		"$lt": reconciliationMetaData.EndTime,
	}})
	if err != nil {
		return transactionsData, fmt.Errorf("error retrieving transactions data: %s", err)
	}

	defer closeCursor(cur)

	if err = cur.All(context.Background(), &transactions); err != nil {
		return transactionsData, fmt.Errorf("error retrieving transactions data: %s", err)
	}

	transactionsData = models.TransactionsList{
		Transactions: transactions,
	}

	return transactionsData, err
}

// GetProductsData fetches products data
func (m *Mongo) GetProductsData(reconciliationMetaData *models.ReconciliationMetaData) (models.ProductsList, error) {

	var products []models.Product
	var productsData models.ProductsList

	collection := getMongoClient(m.Config.Database).Database(m.Config.Database).Collection(m.Config.ProductsCollection)

	cur, err := collection.Find(context.Background(), bson.M{"transaction_date": bson.M{
		"$gt": reconciliationMetaData.StartTime,
		"$lt": reconciliationMetaData.EndTime,
	}})
	if err != nil {
		return productsData, fmt.Errorf("error retrieving products data: %s", err)
	}

	defer closeCursor(cur)

	if err = cur.All(context.Background(), &products); err != nil {
		return productsData, fmt.Errorf("error retrieving products data: %s", err)
	}

	productsData = models.ProductsList{
		Products: products,
	}

	return productsData, err
}

// GetRefundsData fetches refunds data
func (m *Mongo) GetRefundsData(reconciliationMetaData *models.ReconciliationMetaData) (models.RefundsList, error) {

	var refunds []models.Refund
	var refundsData models.RefundsList

	collection := getMongoClient(m.Config.Database).Database(m.Config.Database).Collection(m.Config.RefundsCollection)

	cur, err := collection.Find(context.Background(), bson.M{"transaction_date": bson.M{
		"$gt": reconciliationMetaData.StartTime,
		"$lt": reconciliationMetaData.EndTime,
	}})
	if err != nil {
		return refundsData, fmt.Errorf("error retrieving refunds data: %s", err)
	}

	defer closeCursor(cur)

	if err = cur.All(context.Background(), &refunds); err != nil {
		return refundsData, fmt.Errorf("error retrieving refunds data: %s", err)
	}

	refundsData = models.RefundsList{
		Refunds: refunds,
	}

	return refundsData, err
}

func closeCursor(cur *mongo.Cursor) {
	err := cur.Close(context.Background())
	if err != nil {
		log.Error(fmt.Errorf("error closing mongo cursor: %s", err))
	}
}
