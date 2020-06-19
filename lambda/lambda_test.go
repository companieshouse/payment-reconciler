package lambda

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	l "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/companieshouse/payment-reconciler/config"
	"github.com/companieshouse/payment-reconciler/filetransfer"
	"github.com/companieshouse/payment-reconciler/models"
	"github.com/companieshouse/payment-reconciler/service"
	"github.com/golang/mock/gomock"
	"github.com/pkg/sftp"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/ssh"
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

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("And a products CSV is constructed successfully", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, nil).Times(1)

				Convey("And the CSV's are uploaded successfully", func() {

					csvs := []models.CSV{transactionsCSV, productsCSV}
					mockFileTransfer.EXPECT().UploadCSVFiles(csvs).Return(nil).Times(1)

					Convey("Then the request is successful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Subject: Failure to construct transactions CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a failure when constructing a transactions CSV", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, errors.New("failed to construct transactions CSV")).Times(1)

			Convey("Then there is never an attempt to construct a products CSV", func() {

				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Times(0)

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

	Convey("Subject: Failure to construct products CSV", t, func() {

		mockService := service.NewMockService(mockCtrl)
		mockFileTransfer := filetransfer.NewMockFileTransfer(mockCtrl)

		lambda := createMockLambda(&cfg, mockService, mockFileTransfer)

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("But there's a failure when constructing a products CSV", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, errors.New("failed to construct products CSV")).Times(1)

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

		Convey("Given a transactions CSV is constructed successfully", func() {

			var transactionsCSV models.CSV
			mockService.EXPECT().GetTransactionsCSV(&reconciliationMetaData).Return(transactionsCSV, nil).Times(1)

			Convey("And a products CSV is constructed successfully", func() {

				var productsCSV models.CSV
				mockService.EXPECT().GetProductsCSV(&reconciliationMetaData).Return(productsCSV, nil).Times(1)

				Convey("But the CSV's are not uploaded successfully", func() {

					csvs := []models.CSV{transactionsCSV, productsCSV}
					mockFileTransfer.EXPECT().UploadCSVFiles(csvs).Return(errors.New("failure to upload CSV's")).Times(1)

					Convey("Then the request is unsuccessful", func() {

						err := lambda.Execute(&reconciliationMetaData)
						So(err, ShouldNotBeNil)
					})
				})
			})
		})
	})
}

func TestIntegrationExecute(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfg, err := config.GetTestConfig()
	if err != nil {
		t.Errorf("Failed to fetch test config: %s", err)
		t.FailNow()
	}

	if cfg.NotPopulated() {
		t.Log("Some config required for integration tests is not populated. Skipping test")
		t.SkipNow()
	}

	Convey("Given I connect to the SFTP server", t, func() {

		client, err := ssh.Dial("tcp", cfg.SFTPServer+":"+cfg.SFTPPort, filetransfer.NewSSHConfig(&cfg.SFTPConfig))
		if err != nil {
			t.Errorf("Failed to establish connection to the SFTP server: %s", err)
			t.FailNow()
		}
		defer func() {
			err = client.Close()
			if err != nil {
				t.Errorf("Failed to close connection to the SFTP server: %s", err)
				t.FailNow()
			}
		}()

		sftpSession, err := sftp.NewClient(client)
		if err != nil {
			t.Errorf("Error creating SFTP session: %s", err)
			t.FailNow()
		}
		defer func() {
			err = sftpSession.Close()
			if err != nil {
				t.Errorf("Failed to close SFTP session: %s", err)
				t.FailNow()
			}
		}()

		Convey("And I remove all existing CSV's generated by the lambda", func() {

			files, err := sftpSession.ReadDir(cfg.SFTPFilePath)
			if err != nil {
				t.Errorf("Error reading SFTP directory contents: %s", err)
				t.FailNow()
			}

			for _, file := range files {
				if strings.HasPrefix(file.Name(), service.TransactionsFileNamePrefix) ||
					strings.HasPrefix(file.Name(), service.ProductsFileNamePrefix) {

					err := sftpSession.Remove(filepath.Join(cfg.SFTPFilePath, file.Name()))
					if err != nil {
						t.Errorf("Failed to remove file from SFTP server: %s", err)
						t.FailNow()
					}
				}
			}

			Convey("When I invoke the lambda to reconcile payments", func() {

				sess := session.Must(session.NewSessionWithOptions(session.Options{
					SharedConfigState: session.SharedConfigEnable,
				}))

				awsLambdaClient := l.New(sess, &aws.Config{Region: aws.String(cfg.EnvRegionAWS)})

				yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
				metaData := models.ReconciliationMetaData{ReconciliationDate: yesterday}

				payload, err := json.Marshal(metaData)
				if err != nil {
					t.Errorf("Error marshalling ReconciliationMetaData: %s", err)
					t.FailNow()
				}

				_, err = awsLambdaClient.Invoke(&l.InvokeInput{FunctionName: aws.String(cfg.LambdaName), Payload: payload})
				if err != nil {
					t.Errorf("Failed to invoke lambda: %s", err)
					t.FailNow()
				}

				Convey("Then new csv's are uploaded to the SFTP server", func() {

					transactionsFileName := filepath.Join(cfg.SFTPFilePath, service.TransactionsFileNamePrefix+yesterday+service.CsvFileSuffix)
					productsFileName := filepath.Join(cfg.SFTPFilePath, service.ProductsFileNamePrefix+yesterday+service.CsvFileSuffix)

					// Fetch details of the transactions file
					transactionsFile, transactionsFileErr := sftpSession.Stat(transactionsFileName)

					// Fetch details of the products file
					productsFile, productsFileErr := sftpSession.Stat(productsFileName)

					// Assert the files exist and no errors are present
					So(transactionsFileErr, ShouldBeNil)
					So(productsFileErr, ShouldBeNil)
					So(transactionsFile, ShouldNotBeNil)
					So(productsFile, ShouldNotBeNil)
				})
			})
		})
	})
}
