package report

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(result *audit.Result) {
	fmt.Printf("\n--- AuroraSec Scan Report ---\n")
	fmt.Printf("Started: %s\n", result.StartTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Ended:   %s\n", result.EndTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("\n")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("Severity", "Status", "Module", "Title", "Resource")

	for _, f := range result.Findings {
		sevColor := getSeverityColor(f.Severity)
		statusColor := getStatusColor(f.Status)

		table.Append(
			sevColor(string(f.Severity)),
			statusColor(f.Status),
			f.Module,
			f.Title,
			f.Resource,
		)
	}

	table.Render()

	fmt.Printf("\nSummary: %d Total Findings | ", result.Summary.Total)
	color.Red("%d Critical", result.Summary.Critical)
	fmt.Printf(" | ")
	color.Yellow("%d High", result.Summary.High)
	fmt.Printf(" | %d Medium | %d Low | ", result.Summary.Medium, result.Summary.Low)
	color.Green("%d Passed", result.Summary.Passed)
	fmt.Printf("\n")
}

func getSeverityColor(s audit.Severity) func(a ...interface{}) string {
	switch s {
	case audit.SeverityCritical:
		return color.New(color.FgRed, color.Bold).SprintFunc()
	case audit.SeverityHigh:
		return color.New(color.FgRed).SprintFunc()
	case audit.SeverityMedium:
		return color.New(color.FgYellow).SprintFunc()
	case audit.SeverityLow:
		return color.New(color.FgCyan).SprintFunc()
	default:
		return color.New(color.FgWhite).SprintFunc()
	}
}

func getStatusColor(s string) func(a ...interface{}) string {
	switch s {
	case "PASS":
		return color.New(color.FgGreen).SprintFunc()
	case "FAIL":
		return color.New(color.FgRed).SprintFunc()
	default:
		return color.New(color.FgYellow).SprintFunc()
	}
}
