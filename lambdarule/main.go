package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/eventbridge#PutRuleInput

func HandleRequest() {
	// Load the AWS SDK configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(fmt.Sprintf("failed to load configuration, %v", err))
	}
	// Create an EventBridge client
	client := eventbridge.NewFromConfig(cfg)
	// Set the input for the PutRule API call
	ruleName := "my-eventbridge-rule"
	eventPattern := `{
                        "source": [
                         "aws.ec2"
                        ],
                        "detail-type": [
                         "EC2 Instance State-change Notification"
                        ],
                        "detail": {
                         "state": [
                                     "terminated"
                         ]
                        }
            }`
	ruleInput := &eventbridge.PutRuleInput{
		Name:         &ruleName,
		EventPattern: &eventPattern,
		State:        types.RuleStateEnabled,
	}
	// Call the PutRule API
	result, err := client.PutRule(context.Background(), ruleInput)

	if err != nil {
		panic(fmt.Sprintf("failed to create eventbridge rule, %v", err))
	}
	fmt.Printf("Successfully created eventbridge rule with ARN: %s\n", *result.RuleArn)
}

func main() {
	lambda.Start(HandleRequest)
}
