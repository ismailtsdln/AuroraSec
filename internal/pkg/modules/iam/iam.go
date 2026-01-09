package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

type IAMModule struct {
	client *iam.Client
}

func NewIAMModule(cfg aws.Config) *IAMModule {
	return &IAMModule{
		client: iam.NewFromConfig(cfg),
	}
}

func (m *IAMModule) Name() string {
	return "IAM"
}

func (m *IAMModule) Description() string {
	return "Identity and Access Management security checks"
}

func (m *IAMModule) Audit(ctx context.Context) ([]audit.Finding, error) {
	var findings []audit.Finding

	// Check 1: Root user MFA
	rootMFAFinding, err := m.checkRootMFA(ctx)
	if err == nil {
		findings = append(findings, rootMFAFinding)
	}

	// Check 2: Account Password Policy
	passwordPolicyFinding, err := m.checkPasswordPolicy(ctx)
	if err == nil {
		findings = append(findings, passwordPolicyFinding)
	}

	return findings, nil
}

func (m *IAMModule) checkRootMFA(ctx context.Context) (audit.Finding, error) {
	summary, err := m.client.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
	if err != nil {
		return audit.Finding{}, err
	}

	mfaEnabled := summary.SummaryMap["AccountMFAEnabled"]

	finding := audit.Finding{
		Module:      "IAM",
		ID:          "IAM_ROOT_MFA",
		Title:       "Root Account MFA Enabled",
		Description: "Check if the root account has Multi-Factor Authentication enabled.",
		Severity:    audit.SeverityCritical,
		Resource:    "Root Account",
		Timestamp:   time.Now(),
	}

	if mfaEnabled > 0 {
		finding.Status = "PASS"
	} else {
		finding.Status = "FAIL"
		finding.Remediation = "Enable MFA for the root account in the IAM console."
	}

	return finding, nil
}

func (m *IAMModule) checkPasswordPolicy(ctx context.Context) (audit.Finding, error) {
	_, err := m.client.GetAccountPasswordPolicy(ctx, &iam.GetAccountPasswordPolicyInput{})

	finding := audit.Finding{
		Module:      "IAM",
		ID:          "IAM_PASSWORD_POLICY",
		Title:       "Account Password Policy Defined",
		Description: "Check if a custom password policy is defined for IAM users.",
		Severity:    audit.SeverityMedium,
		Resource:    "Account Password Policy",
		Timestamp:   time.Now(),
	}

	if err != nil {
		finding.Status = "FAIL"
		finding.Remediation = "Define a strong password policy for the AWS account."
	} else {
		finding.Status = "PASS"
	}

	return finding, nil
}
