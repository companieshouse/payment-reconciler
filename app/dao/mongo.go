package dao

import (
	"fmt"

	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/models"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Config *config.Config
}

func New(cfg *config.Config) *Mongo {

	return &Mongo{
		Config: cfg,
	}
}

func getMongoSession(cfg *config.Config) (*mgo.Session, error) {

	s, err := mgo.Dial(cfg.MongoDBURL)
	if err != nil {
		return nil, err
	}

	return s.Copy(), nil
}

func (m *Mongo) GetReconciliationData() (*models.ReconciliationData, error) {

	mongoSession, err := getMongoSession(m.Config)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %s", err)
	}
	defer mongoSession.Close()

	payments := []models.PaymentResourceData{}

	c := mongoSession.DB(m.Config.Database).C(m.Config.Collection)

	err = c.Find(bson.M{}).All(&payments)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving reconciliation data: %s", err)
	}

	reconciliationData := &models.ReconciliationData{
		Payments: payments,
	}

	return reconciliationData, err
}
