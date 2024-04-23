package ddb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DDB struct {
	DB *dynamodb.Client
}

var DB *DDB = &DDB{}

func (c *DDB) Connect() {
	cre := config.WithCredentialsProvider(
		aws.CredentialsProviderFunc(
			func(_ context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     ddb.Credentials.AccessKeyID,
					SecretAccessKey: ddb.Credentials.SecretAccessKey,
					SessionToken:    ddb.Credentials.SessionToken,
					Source:          ddb.Credentials.Source,
				}, nil
			},
		),
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{URL: "http://localhost:8000"}, nil
		})))

	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	c.DB = client
	c.CreateTable()
}

func (c *DDB) CreateTable() {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("MyTable"),
	}

	_, err := c.DB.CreateTable(context.TODO(), input)
	if err != nil {
		panic(err)
	}
}
