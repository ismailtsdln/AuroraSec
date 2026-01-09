package report

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(result *audit.Result) {
	fmt.Printf("\n%s\n", color.New(color.FgCyan, color.Bold).Sprint("─── AURORASEC SCAN REPORT ───"))
	fmt.Printf("%s %s\n", color.CyanString("📅 Start:"), result.StartTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("%s %s\n", color.CyanString("🏁 End:  "), result.EndTime.Format("2006-01-02 15:04:05"))
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("SEVERITY", "STATUS", "MODULE", "TITLE", "RESOURCE")

	for _, f := range result.Findings {
		sevColor := getSeverityColor(f.Severity)
		statusColor := getStatusColor(f.Status)

		severityIcon := getSeverityIcon(f.Severity)
		statusIcon := getStatusIcon(f.Status)

		table.Append(
			sevColor(fmt.Sprintf("%s %s", severityIcon, f.Severity)),
			statusColor(fmt.Sprintf("%s %s", statusIcon, f.Status)),
			f.Module,
			f.Title,
			f.Resource,
		)
	}

	table.Render()

	fmt.Printf("\n%s\n", color.New(color.FgCyan, color.Bold).Sprint("─── SUMMARY ───"))
	fmt.Printf("%s %d Total Findings\n", color.WhiteString("📊"), result.Summary.Total)

	summaryLine := ""
	if result.Summary.Critical > 0 {
		summaryLine += color.New(color.FgRed, color.Bold).Sprintf("🚨 %d Critical  ", result.Summary.Critical)
	}
	if result.Summary.High > 0 {
		summaryLine += color.RedString("🔴 %d High  ", result.Summary.High)
	}
	if result.Summary.Medium > 0 {
		summaryLine += color.YellowString("🟡 %d Medium  ", result.Summary.Medium)
	}
	if result.Summary.Low > 0 {
		summaryLine += color.CyanString("🔵 %d Low  ", result.Summary.Low)
	}

	if summaryLine != "" {
		fmt.Println(summaryLine)
	}

	fmt.Printf("%s %d Passed  ", color.GreenString("✅"), result.Summary.Passed)
	fmt.Printf("%s %d Failed\n", color.RedString("❌"), result.Summary.Failed)
	fmt.Println()
}

func getSeverityIcon(s audit.Severity) string {
	switch s {
	case audit.SeverityCritical:
		return "🚨"
	case audit.SeverityHigh:
		return "🔴"
	case audit.SeverityMedium:
		return "🟡"
	case audit.SeverityLow:
		return "🔵"
	default:
		return "⚪"
	}
}

func getStatusIcon(s string) string {
	switch s {
	case "PASS":
		return "✅"
	case "FAIL":
		return "❌"
	default:
		return "❓"
	}
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
