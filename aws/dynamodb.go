package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *dynamodb.Client
var table string

func init() {
	table = os.Getenv("DYNAMO_TABLE_NAME")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(getEndpoint)))
	if err != nil {
		panic(err)
	}
	client = dynamodb.NewFromConfig(cfg)
}

func getEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	endpoint := aws.Endpoint{}

	if url, exists := os.LookupEnv("DYNAMO_ENDPOINT"); exists {
		endpoint.URL = url
		return endpoint, nil
	}
	return endpoint, &aws.EndpointNotFoundError{}
}

func Scan() ([]Screenshot, error) {
	output, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &table,
	})
	if err != nil {
		return nil, err
	}

	screenshots := []Screenshot{}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &screenshots)
	if err != nil {
		return nil, err
	}

	return screenshots, nil
}
