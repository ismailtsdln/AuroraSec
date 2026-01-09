package networking

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

type NetworkingModule struct {
	client *ec2.Client
}

func NewNetworkingModule(cfg aws.Config) *NetworkingModule {
	return &NetworkingModule{
		client: ec2.NewFromConfig(cfg),
	}
}

func (m *NetworkingModule) Name() string {
	return "Networking"
}

func (m *NetworkingModule) Description() string {
	return "Networking security checks (Security Groups, VPC, NACLs)"
}

func (m *NetworkingModule) Audit(ctx context.Context) ([]audit.Finding, error) {
	var findings []audit.Finding

	// Check 1: Security Groups with 0.0.0.0/0 on sensitive ports
	sgFindings, err := m.checkSecurityGroups(ctx)
	if err == nil {
		findings = append(findings, sgFindings...)
	}

	return findings, nil
}

func (m *NetworkingModule) checkSecurityGroups(ctx context.Context) ([]audit.Finding, error) {
	output, err := m.client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{})
	if err != nil {
		return nil, err
	}

	var findings []audit.Finding
	for _, sg := range output.SecurityGroups {
		sgID := aws.ToString(sg.GroupId)
		sgName := aws.ToString(sg.GroupName)

		for _, perm := range sg.IpPermissions {
			for _, rangeInfo := range perm.IpRanges {
				if aws.ToString(rangeInfo.CidrIp) == "0.0.0.0/0" {
					finding := audit.Finding{
						Module:      "Networking",
						ID:          "NET_WIDE_OPEN_SG",
						Title:       "Security Group Open to the World",
						Description: "Security group allows traffic from 0.0.0.0/0 on potentially sensitive ports.",
						Severity:    audit.SeverityHigh,
						Resource:    fmt.Sprintf("%s (%s)", sgName, sgID),
						Status:      "FAIL",
						Remediation: "Restrict security group rules to specific IP ranges or other security groups.",
						Timestamp:   time.Now(),
					}
					findings = append(findings, finding)
				}
			}
		}
	}

	return findings, nil
}
