package main

import (
	"fmt"

	. "github.com/aws/aws-lambda-go/lambda"
	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/payment-reconciler/app/config"
	"github.com/companieshouse/payment-reconciler/app/lambda"
)

func main() {

	log.Namespace = "payment-reconciler"

	cfg, err := config.Get()
	if err != nil {
		log.Error(fmt.Errorf("error configuring service: %s. - exiting", err), nil)
		return
	}

	log.Trace("Config", log.Data{"Config": cfg})
	log.Info("Payment reconciliation lambda started")

	reconciliationLambda := lambda.New(cfg)

	Start(reconciliationLambda.Execute)
}
