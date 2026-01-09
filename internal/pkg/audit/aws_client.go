package audit

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// AWSClient wraps the AWS SDK configuration
type AWSClient struct {
	Config aws.Config
}

// NewAWSClient initializes the AWS SDK with default credentials and region
func NewAWSClient(ctx context.Context, profile string, region string) (*AWSClient, error) {
	var opts []func(*config.LoadOptions) error

	if profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}

	if region != "" {
		opts = append(opts, config.WithRegion(region))
	}

	// Add custom retryer
	opts = append(opts, config.WithRetryer(func() aws.Retryer {
		return NewCustomRetryer()
	}))

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	return &AWSClient{
		Config: cfg,
	}, nil
}
