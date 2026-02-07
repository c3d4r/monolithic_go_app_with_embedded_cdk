package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

// Synthesize creates the CDK app, reads context, defines stacks,
// and writes the cloud assembly to cdk.out/.
func Synthesize() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	cfg := ConfigFromContext(app)

	NewPicGalleryStack(app, "PicGalleryStack", cfg, &awscdk.StackProps{
		Env: env(),
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil // Use CDK_DEFAULT_ACCOUNT and CDK_DEFAULT_REGION
}
