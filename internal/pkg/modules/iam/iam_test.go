package iam

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type mockIAMClient struct {
	IAMAPI
	GetAccountSummaryFunc func(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error)
}

func (m *mockIAMClient) GetAccountSummary(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error) {
	return m.GetAccountSummaryFunc(ctx, params, optFns...)
}

func (m *mockIAMClient) GetAccountPasswordPolicy(ctx context.Context, params *iam.GetAccountPasswordPolicyInput, optFns ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error) {
	return &iam.GetAccountPasswordPolicyOutput{}, nil
}

func TestCheckRootMFA(t *testing.T) {
	tests := []struct {
		name       string
		mfaEnabled int32
		wantStatus string
	}{
		{
			name:       "MFA Enabled",
			mfaEnabled: 1,
			wantStatus: "PASS",
		},
		{
			name:       "MFA Disabled",
			mfaEnabled: 0,
			wantStatus: "FAIL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockIAMClient{
				GetAccountSummaryFunc: func(ctx context.Context, params *iam.GetAccountSummaryInput, optFns ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error) {
					return &iam.GetAccountSummaryOutput{
						SummaryMap: map[string]int32{
							"AccountMFAEnabled": tt.mfaEnabled,
						},
					}, nil
				},
			}

			m := &IAMModule{client: mockClient}
			finding, err := m.checkRootMFA(context.Background())
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if finding.Status != tt.wantStatus {
				t.Errorf("got status %s, want %s", finding.Status, tt.wantStatus)
			}
		})
	}
}
