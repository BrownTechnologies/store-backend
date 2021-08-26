package dynamodb

import (
	"log"
	"store-backend/modals"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DB struct {
	DBSession *dynamodb.DynamoDB
}

const tableZipCodeJapan string = "zipcode-japan"

func New() DB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	return DB{DBSession: dynamodb.New(sess)}
}

func (svc *DB) InsertIntoZipcode(entry modals.ZipCodeEntry) {

	mEntery, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		log.Fatalf("Got error marshalling new zipcode entry item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      mEntery,
		TableName: aws.String(tableZipCodeJapan),
	}

	_, err = svc.DBSession.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

func (svc *DB) InsertBatchZipcodes(entry ...modals.ZipCodeEntry) {

	mEntery, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		log.Fatalf("Got error marshalling new zipcode entry item: %s", err)
	}

	input := &dynamodb.BatchWriteItemInput()

	_, err = svc.DBSession.BatchWriteItem()
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}
