package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// Config holds shared configuration used by both the application and infrastructure.
type Config struct {
	StackName string
}

// ConfigFromContext reads configuration from CDK context values,
// which are passed via -c flags or cdk.json context.
func ConfigFromContext(app awscdk.App) Config {
	return Config{
		StackName: contextString(app, "stack_name"),
	}
}

func contextString(app awscdk.App, key string) string {
	val := app.Node().TryGetContext(jsii.String(key))
	if val == nil {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}
