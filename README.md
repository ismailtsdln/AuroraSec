# AuroraSec

[![AuroraSec CI](https://github.com/ismailtsdln/AuroraSec/actions/workflows/main.yml/badge.svg)](https://github.com/ismailtsdln/AuroraSec/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ismailtsdln/AuroraSec)](https://goreportcard.com/report/github.com/ismailtsdln/AuroraSec)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**AuroraSec** is a high-performance, modular AWS auditing and hardening tool built in Go. It allows security professionals to quickly identify misconfigurations and security risks across their AWS accounts.

## Core Features

- 🔒 **IAM Hardening**: Audits MFA status, password policies, and access key rotation.
- 📡 **Networking Security**: Identifies Security Groups with wide-open ports (0.0.0.0/0).
- 📦 **S3 Audit**: Detects public buckets and ensures default encryption is enabled.
- 📊 **Multi-Format Reporting**: Supports Terminal Tables, JSON, CSV, and interactive HTML reports.
- 🚀 **Modular Architecture**: Easy to extend with new security checks.
- 🛡️ **Built-in Resilience**: Custom retry logic and error wrapping for stable AWS SDK operations.

## Installation

### From Binary

Download the latest binary from the [Releases](https://github.com/ismailtsdln/AuroraSec/releases) page.

### From Source

```bash
go install github.com/ismailtsdln/AuroraSec/cmd/aurorasec@latest
```

## Usage

### Basic Audit

Run a full audit with default modules (IAM, S3, Networking):

```bash
aurorasec audit
```

### Specific Modules

Only run specific modules:

```bash
aurorasec audit --modules iam,s3
```

### Output Formats

Generate an HTML report:

```bash
aurorasec audit --format html --output report.html
```

Generate a JSON report for automation:

```bash
aurorasec audit --format json --output results.json
```

## Architecture

AuroraSec is designed for scalability and maintainability:

- **`cmd/`**: CLI entry point and command definitions.
- **`internal/pkg/audit/`**: Core engine that orchestrates module execution.
- **`internal/pkg/modules/`**: Plugin-style modules for different AWS services.
- **`internal/pkg/report/`**: Result formatters for various output styles.

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## Security

If you discover a security vulnerability, please see [SECURITY.md](SECURITY.md) for reporting instructions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
