package cli

import (
	"context"
	"fmt"

	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/modules/iam"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/modules/networking"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/modules/s3"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/report"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/ui"
	"github.com/ismailtsdln/AuroraSec/pkg/utils"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Run security audit checks on AWS environment",
	Long:  `AuroraSec audit scans your AWS infrastructure for security vulnerabilities, misconfigurations, and non-compliance with best practices.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		logger := utils.NewLogger("INFO")

		ui.PrintBanner()
		logger.Info("Initializing AWS Security Audit...")

		// 1. Initialize AWS Client
		awsClient, err := audit.NewAWSClient(ctx, "", "")
		if err != nil {
			return fmt.Errorf("failed to initialize AWS client: %v", err)
		}
		logger.Success("AWS SDK connection established")

		// 2. Initialize Audit Engine
		engine := audit.NewEngine()

		// 3. Register Modules (based on flags)
		modules, _ := cmd.Flags().GetStringSlice("modules")
		for _, m := range modules {
			logger.Info("Registering module: %s", m)
			switch m {
			case "iam":
				engine.RegisterModule(iam.NewIAMModule(awsClient.Config))
			case "s3":
				engine.RegisterModule(s3.NewS3Module(awsClient.Config))
			case "networking":
				engine.RegisterModule(networking.NewNetworkingModule(awsClient.Config))
			}
		}

		// 4. Run Audit
		logger.Info("Running security checks... Hold on!")
		result, err := engine.Run(ctx)
		if err != nil {
			return fmt.Errorf("audit failed: %v", err)
		}
		logger.Success("Audit completed successfully!")

		// 5. Report Results
		format, _ := cmd.Flags().GetString("format")
		outputFile, _ := cmd.Flags().GetString("output")

		switch format {
		case "table":
			report.PrintTable(result)
		case "json":
			if err := report.SaveJSON(result, outputFile); err != nil {
				return err
			}
		case "csv":
			if err := report.SaveCSV(result, outputFile); err != nil {
				return err
			}
		case "html":
			if err := report.SaveHTML(result, outputFile); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported output format: %s", format)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(auditCmd)

	auditCmd.Flags().StringSliceP("modules", "m", []string{"iam", "s3", "ec2"}, "Specific modules to audit (comma-separated)")
	auditCmd.Flags().StringP("format", "f", "table", "Output format (table, json, html)")
	auditCmd.Flags().StringP("output", "o", "", "Output file path")
}
