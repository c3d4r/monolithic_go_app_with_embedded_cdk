package infra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewPicGalleryStack(scope constructs.Construct, id string, cfg Config, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	// S3 bucket for gallery assets
	bucket := awss3.NewBucket(stack, jsii.String("GalleryBucket"), &awss3.BucketProps{
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
	})

	// Look up the hosted zone
	zone := awsroute53.HostedZone_FromHostedZoneAttributes(stack, jsii.String("Zone"), &awsroute53.HostedZoneAttributes{
		HostedZoneId: jsii.String(cfg.HostedZoneID),
		ZoneName:     jsii.String(cfg.ZoneName),
	})

	// TLS certificate for the domain
	cert := awscertificatemanager.NewCertificate(stack, jsii.String("Cert"), &awscertificatemanager.CertificateProps{
		DomainName: jsii.String(cfg.DomainName),
		Validation: awscertificatemanager.CertificateValidation_FromDns(zone),
	})

	// CloudFront distribution
	distribution := awscloudfront.NewDistribution(stack, jsii.String("Distribution"), &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			Origin: awscloudfrontorigins.S3BucketOrigin_WithOriginAccessControl(bucket, &awscloudfrontorigins.S3BucketOriginWithOACProps{}),
			ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
		DomainNames:        &[]*string{jsii.String(cfg.DomainName)},
		Certificate:        cert,
		DefaultRootObject:  jsii.String("index.html"),
	})

	// DNS alias record
	awsroute53.NewARecord(stack, jsii.String("AliasRecord"), &awsroute53.ARecordProps{
		Zone:       zone,
		RecordName: jsii.String(cfg.DomainName),
		Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewCloudFrontTarget(distribution)),
	})

	// Outputs
	awscdk.NewCfnOutput(stack, jsii.String("BucketName"), &awscdk.CfnOutputProps{
		Value:       bucket.BucketName(),
		Description: jsii.String("S3 bucket for gallery assets"),
	})

	awscdk.NewCfnOutput(stack, jsii.String("DistributionDomainName"), &awscdk.CfnOutputProps{
		Value:       distribution.DistributionDomainName(),
		Description: jsii.String("CloudFront distribution domain"),
	})

	awscdk.NewCfnOutput(stack, jsii.String("URL"), &awscdk.CfnOutputProps{
		Value:       jsii.String("https://" + cfg.DomainName),
		Description: jsii.String("Gallery URL"),
	})

	return stack
}
