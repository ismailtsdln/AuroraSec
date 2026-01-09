package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

type S3Module struct {
	client *s3.Client
}

func NewS3Module(cfg aws.Config) *S3Module {
	return &S3Module{
		client: s3.NewFromConfig(cfg),
	}
}

func (m *S3Module) Name() string {
	return "S3"
}

func (m *S3Module) Description() string {
	return "S3 bucket security and encryption checks"
}

func (m *S3Module) Audit(ctx context.Context) ([]audit.Finding, error) {
	var findings []audit.Finding

	buckets, err := m.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}

	for _, b := range buckets.Buckets {
		bucketName := aws.ToString(b.Name)

		// Check 1: Bucket Encryption
		findings = append(findings, m.checkBucketEncryption(ctx, bucketName))

		// Check 2: Public Access Blocks
		findings = append(findings, m.checkPublicAccessBlock(ctx, bucketName))
	}

	return findings, nil
}

func (m *S3Module) checkBucketEncryption(ctx context.Context, bucketName string) audit.Finding {
	_, err := m.client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucketName),
	})

	finding := audit.Finding{
		Module:      "S3",
		ID:          "S3_BUCKET_ENCRYPTION",
		Title:       "S3 Bucket Default Encryption Enabled",
		Description: "Verify if the S3 bucket has default encryption enabled.",
		Severity:    audit.SeverityHigh,
		Resource:    bucketName,
		Timestamp:   time.Now(),
	}

	if err != nil {
		finding.Status = "FAIL"
		finding.Remediation = "Enable default encryption for the S3 bucket using AES-256 or AWS-KMS."
	} else {
		finding.Status = "PASS"
	}

	return finding
}

func (m *S3Module) checkPublicAccessBlock(ctx context.Context, bucketName string) audit.Finding {
	output, err := m.client.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{
		Bucket: aws.String(bucketName),
	})

	finding := audit.Finding{
		Module:      "S3",
		ID:          "S3_PUBLIC_ACCESS_BLOCK",
		Title:       "S3 Public Access Block Enabled",
		Description: "Verify if the S3 bucket has public access blocks enabled.",
		Severity:    audit.SeverityCritical,
		Resource:    bucketName,
		Timestamp:   time.Now(),
	}

	isBlocked := false
	if output.PublicAccessBlockConfiguration != nil && output.PublicAccessBlockConfiguration.BlockPublicAcls != nil {
		isBlocked = *output.PublicAccessBlockConfiguration.BlockPublicAcls
	}

	if err != nil || !isBlocked {
		finding.Status = "FAIL"
		finding.Remediation = "Enable 'Block all public access' for the S3 bucket."
	} else {
		finding.Status = "PASS"
	}

	return finding
}
