package report

import (
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/ismailtsdln/AuroraSec/internal/pkg/audit"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>AuroraSec Audit Report</title>
    <style>
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 20px; }
        .container { max-width: 1200px; margin: auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px; }
        .summary { display: flex; gap: 20px; margin-bottom: 30px; }
        .summary-card { flex: 1; padding: 20px; border-radius: 8px; text-align: center; color: white; font-weight: bold; }
        .total { background-color: #34495e; }
        .critical { background-color: #e74c3c; }
        .high { background-color: #e67e22; }
        .medium { background-color: #f1c40f; color: #333; }
        .pass { background-color: #2ecc71; }
        table { width: 100%; border-collapse: collapse; margin-top: 20px; }
        th, td { padding: 12px; text-align: left; border-bottom: 1px solid #ddd; }
        th { background-color: #f8f9fa; color: #2c3e50; }
        tr:hover { background-color: #f1f1f1; }
        .severity-CRITICAL { color: #e74c3c; font-weight: bold; }
        .severity-HIGH { color: #e67e22; font-weight: bold; }
        .severity-MEDIUM { color: #f39c12; font-weight: bold; }
        .status-FAIL { color: #e74c3c; }
        .status-PASS { color: #2ecc71; }
        .remediation { font-size: 0.9em; color: #7f8c8d; font-style: italic; }
    </style>
</head>
<body>
    <div class="container">
        <h1>AuroraSec Audit Report</h1>
        <p><strong>Scan Date:</strong> {{ .EndTime.Format "2006-01-02 15:04:05" }}</p>
        
        <div class="summary">
            <div class="summary-card total">Total: {{ .Summary.Total }}</div>
            <div class="summary-card critical">Critical: {{ .Summary.Critical }}</div>
            <div class="summary-card high">High: {{ .Summary.High }}</div>
            <div class="summary-card medium">Medium: {{ .Summary.Medium }}</div>
            <div class="summary-card pass">Passed: {{ .Summary.Passed }}</div>
        </div>

        <table>
            <thead>
                <tr>
                    <th>Severity</th>
                    <th>Status</th>
                    <th>Module</th>
                    <th>Title</th>
                    <th>Resource</th>
                    <th>Remediation</th>
                </tr>
            </thead>
            <tbody>
                {{ range .Findings }}
                <tr>
                    <td><span class="severity-{{ .Severity }}">{{ .Severity }}</span></td>
                    <td><span class="status-{{ .Status }}">{{ .Status }}</span></td>
                    <td>{{ .Module }}</td>
                    <td>{{ .Title }}</td>
                    <td>{{ .Resource }}</td>
                    <td class="remediation">{{ .Remediation }}</td>
                </tr>
                {{ end }}
            </tbody>
        </table>
    </div>
</body>
</html>
`

func SaveHTML(result *audit.Result, filepath string) error {
	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %v", err)
	}

	if filepath == "" {
		filepath = fmt.Sprintf("aurorasec-report-%d.html", time.Now().Unix())
	}

	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create HTML file: %v", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, result); err != nil {
		return fmt.Errorf("failed to execute HTML template: %v", err)
	}

	fmt.Printf("✅ HTML report saved to %s\n", filepath)
	return nil
}
