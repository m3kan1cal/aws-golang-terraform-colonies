package libs

import "github.com/aws/aws-lambda-go/events"

// Response represents AWS Lambda Proxy Request/Response functionality (default behavior):
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Request represents AWS Lambda Proxy Request/Response functionality (default behavior):
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Request events.APIGatewayProxyRequest

// Colony represents metadata attached to colony.
type Colony struct {
	ColonyID    string `json:"colonyId"`
	Planet      string `json:"planet"`
	Corporation string `json:"corporation"`
	Coords      string `json:"coords"`
}

// ClientError represents custom error details for errors sent back
// to clients consuming services.
type ClientError struct {
	Error string `json:"error"`
}
