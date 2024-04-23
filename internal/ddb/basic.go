package ddb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DDB struct {
	DB *dynamodb.DynamoDB
}

var DB *DDB = &DDB{}

func (c *DDB) Connect() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-1"),
		Endpoint: aws.String("http://localhost:8000"),
	})

	if err != nil {
		panic(err)
	}

	c.DB = dynamodb.New(sess)
	c.CreateTable()
}

func (c *DDB) CreateTable() {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("MyTable"),
	}

	_, err := c.DB.CreateTable(input)
	if err != nil {
		panic(err)
	}
}
