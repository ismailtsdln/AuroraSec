# AuroraSec

**AuroraSec** is a next-generation AWS auditing and hardening tool built with Go. It provides a modular approach to security auditing, integration with AWS SDK v2, and high-performance reporting.

## Features

- 🔒 **IAM Hardening**: Checks for MFA, unused keys, and root account security.
- 📡 **Network Security**: Audit Security Groups, NACLs, and VPC configurations.
- 📦 **S3 & Storage**: Identify public buckets and unencrypted data.
- 📝 **Logging & Monitoring**: Verify CloudTrail, GuardDuty, and Config setup.
- 📊 **Rich Reporting**: Output findings in JSON, CSV, and interactive HTML.
- 🚀 **High Performance**: Built with Go for speed and concurrency.

## Installation

```bash
go install github.com/ismailtsdln/AuroraSec/cmd/aurorasec@latest
```

## Quick Start

```bash
# Basic audit using default profile
aurorasec audit

# Audit specific modules
aurorasec audit --modules iam,s3

# Generate HTML report
aurorasec audit --format html --output report.html
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
